#!/bin/sh

echo -e "\nUpgrading Noria from $NORIA_FROM_VERSION to $NORIA_TO_VERSION\n via $NORIA_UPGRADE_NAME\n"

cd /app
git pull --all

# Build TO noriad version
if [[ "$NORIA_TO_VERSION" == "current" ]]; then
	git config --global --add safe.directory /current
	cd ../current
	rm -rf build
	make build
	mv build/noriad /tmp/noriad
	cd ../app
else
	git checkout $NORIA_TO_VERSION
	rm -rf build
	make build
	mv build/noriad /tmp/noriad
fi

# Copy a working version of required scripts
git checkout a7975921a8c63d09f9330c3aec1f6d88593246ab
cp /app/scripts/init_local.sh /init_local.sh
cp /app/txTest/upgrade.sh /upgrade.sh

# Install FROM noriad version
git checkout $NORIA_FROM_VERSION
rm -rf build
make build && make install

# cp /app/scripts/init_local.sh /init_local.sh
# cp /app/txTest/upgrade.sh /upgrade.sh

# Initialize the noriad config
source /init_local.sh

# Prepare Cosmovisor configs
BINARY_DIR=".noria"
export DAEMON_NAME="noriad"
export DAEMON_HOME="$HOME/$BINARY_DIR"
export DAEMON_ALLOW_DOWNLOAD_BINARIES=false

# Make voting period shorter for rapid testing
sed -i.bak 's/"voting_period": "[^"]*"/"voting_period": "8s"/g' $DAEMON_HOME/config/genesis.json >/dev/null 2>&1

# Move both FROM noriad and TO noriad to the cosmovisor folder, ready for upgrade
mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin
cp $(which $DAEMON_NAME) $DAEMON_HOME/cosmovisor/genesis/bin/

mkdir -p $DAEMON_HOME/cosmovisor/upgrades/$NORIA_UPGRADE_NAME/bin
cp /tmp/noriad $DAEMON_HOME/cosmovisor/upgrades/$NORIA_UPGRADE_NAME/bin/noriad

# Define upgrade function that will create a proposal and vote for it
upgrade() {

	HEIGHT=$(curl -s http://localhost:26657/status | jq -r '.result.sync_info.latest_block_height | tonumber')
	UPGRADE_HEIGHT=$(($HEIGHT + 10))
	/upgrade.sh $UPGRADE_HEIGHT $NORIA_UPGRADE_NAME
}

# Start the node and the upgrade flow
cosmovisor run start &
sleep 3 &&
	upgrade &
wait
