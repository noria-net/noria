#!/bin/bash
#
# Test an upgrade on your local environment

HEIGHT=$1
UPGRADE_NAME="v0.7.4"
BINARY_DIR=".noria"
CHAIN_ID="oasis-3"
DENOM="unoria"
GAS_PRICE="0.0025"
GAS_PRICE_DENOM="ucrd"
BINARY_DOWNLOAD_URL="xxx"
FROM_KEY_NAME="me"
export DAEMON_NAME="noriad"
export DAEMON_HOME="$HOME/$BINARY_DIR"

if ! command -v cosmovisor &>/dev/null; then
  echo "\n\ncosmovisor could not be found"
  exit
fi

# submit upgrade proposal
echo Submitting upgrade proposal
$DAEMON_NAME tx gov submit-proposal software-upgrade $UPGRADE_NAME \
    --title "Upgrade to $UPGRADE_NAME" \
    --description "Upgrade to $UPGRADE_NAME" \
    --upgrade-info='{"binaries":{"linux/amd64":$BINARY_DOWNLOAD_URL}}' \
    --deposit 10000000$DENOM \
    --upgrade-height $HEIGHT \
    --from $FROM_KEY_NAME \
    --chain-id $CHAIN_ID \
    --keyring-backend test \
    --home $DAEMON_HOME \
    --node tcp://localhost:26657 \
    --yes \
    --gas-prices $GAS_PRICE$GAS_PRICE_DENOM \
    --gas auto \
    --gas-adjustment 1.5 \
    --broadcast-mode block

# vote on the proposal
echo Voting YES on the upgrade proposal
$DAEMON_NAME tx gov vote 1 yes \
    --from $FROM_KEY_NAME \
    --chain-id $CHAIN_ID \
    --keyring-backend test \
    --home $DAEMON_HOME \
    --node tcp://localhost:26657 \
    --yes \
    --gas-prices $GAS_PRICE$GAS_PRICE_DENOM \
    --gas auto \
    --gas-adjustment 1.5 \
    --broadcast-mode block

# make install
# mkdir -p $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME/bin/
# cp `which $DAEMON_NAME` $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME/bin/