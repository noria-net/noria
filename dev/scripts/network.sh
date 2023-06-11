#!/bin/bash

# Requires noria and hermes images: make build-image && make build-hermes-image

### FUNCTIONS

update_genesis() {
  sed -i.bak "s/stake/unoria/g" $1/genesis.json
  sed -i.bak 's/"inflation": "[^"]*"/"inflation": "0\.0"/g' $1/genesis.json
  sed -i.bak 's/"inflation_rate_change": "[^"]*"/"inflation_rate_change": "0\.0"/g' $1/genesis.json
  sed -i.bak 's/"inflation_min": "[^"]*"/"inflation_min": "0\.0"/g' $1/genesis.json
  sed -i.bak 's/"voting_period": "[^"]*"/"voting_period": "5s"/g' $1/genesis.json
  sed -i.bak 's/"quorum": "[^"]*"/"quorum": "0.000001"/g' $1/genesis.json
  sed -i.bak 's/"reward_delay_time": "[^"]*"/"reward_delay_time": "1s"/g' $1/genesis.json
  sed -i.bak 's/"admin": "[^"]*"/"admin": "'$ADDR'"/g' $1/genesis.json

  tmp=$(mktemp)
  jq '.app_state.tokenfactory.params.denom_creation_fee[0].denom = "ucrd"' $1/genesis.json >"$tmp" && mv "$tmp" $1/genesis.json
}

update_configs() {
  sed -i.bak 's/^timeout_commit\ =\ .*/timeout_commit\ =\ \"1s\"/g' $1/config.toml
  sed -i.bak "s/^minimum-gas-prices\ =\ .*/minimum-gas-prices\ =\ \"0.0025ucrd\"/g" $1/app.toml
  sed -i.bak 's/^enable\ =\ false/enable\ =\ true/g' $1/app.toml
  sed -i.bak 's/^swagger\ =\ false/swagger\ =\ true/g' $1/app.toml
  sed -i.bak '/Rosetta API/{n; s/true/false/}' $1/app.toml
  sed -i.bak "s/^enabled-unsafe-cors\ =\ .*/enabled-unsafe-cors\ =\ true/g" $1/app.toml
  sed -i.bak "s/localhost/0.0.0.0/g" $1/app.toml
  sed -i.bak "s/127\.0\.0\.1/0.0.0.0/g" $1/config.toml
}

stop_nodes() {
  docker stop hermes > /dev/null 2>&1
  docker stop val1 > /dev/null 2>&1
  docker stop val2 > /dev/null 2>&1
}

rm_nodes() {
  docker container rm hermes > /dev/null 2>&1
  docker container rm val1 > /dev/null 2>&1
  docker container rm val2 > /dev/null 2>&1
}

### VARS

HERMES_CONFIG=$(
  cat <<-END
[global]\n
log_level = "info"\n
\n
[mode.clients]\n
enabled = true\n
refresh = true\n
misbehaviour = true\n
\n
[mode.connections]\n
enabled = false\n
\n
[mode.channels]\n
enabled = false\n
\n
[mode.packets]\n
enabled = true\n
clear_interval = 100\n
clear_on_start = true\n
tx_confirmation = false\n
auto_register_counterparty_payee = false\n
\n
[rest]\n
enabled = false\n
host = "127.0.0.1"\n
port = 3000\n
\n
[telemetry]\n
enabled = false\n
host = "127.0.0.1"\n
port = 3001\n
\n
[[chains]]\n
id = "oasis-3"\n
type = "CosmosSdk"\n
rpc_addr = "http://val1:26657"\n
websocket_addr = "ws://val1:26657/websocket"\n
grpc_addr = "http://val1:9090"\n
rpc_timeout = "10s"\n
batch_delay = "500ms"\n
trusted_node = false\n
account_prefix = "noria"\n
key_name = "testkey"\n
key_store_type = "Test"\n
store_prefix = "ibc"\n
default_gas = 100000\n
max_gas = 400000\n
gas_multiplier = 1.5\n
max_msg_num = 30\n
max_tx_size = 180000\n
max_grpc_decoding_size = 33554432\n
clock_drift = "5s"\n
max_block_time = "30s"\n
ccv_consumer_chain = false\n
memo_prefix = ""\n
sequential_batch_tx = false\n
\n
[chains.trust_threshold]\n
numerator = "1"\n
denominator = "3"\n
\n
[chains.gas_price]\n
price = 0.0025\n
denom = "ucrd"\n
\n
[chains.address_type]\n
derivation = "cosmos"\n
\n
[[chains]]\n
id = "oasis-4"\n
type = "CosmosSdk"\n
rpc_addr = "http://val2:26657"\n
websocket_addr = "ws://val2:26657/websocket"\n
grpc_addr = "http://val2:9090"\n
rpc_timeout = "10s"\n
batch_delay = "500ms"\n
trusted_node = false\n
account_prefix = "noria"\n
key_name = "testkey"\n
key_store_type = "Test"\n
store_prefix = "ibc"\n
default_gas = 100000\n
max_gas = 400000\n
gas_multiplier = 1.5\n
max_msg_num = 30\n
max_tx_size = 180000\n
max_grpc_decoding_size = 33554432\n
clock_drift = "5s"\n
max_block_time = "30s"\n
ccv_consumer_chain = false\n
memo_prefix = ""\n
sequential_batch_tx = false\n
\n
[chains.trust_threshold]\n
numerator = "1"\n
denominator = "3"\n
\n
[chains.gas_price]\n
price = 0.0025\n
denom = "ucrd"\n
\n
[chains.address_type]\n
derivation = "cosmos"\n
\n
END
)

IMG="noria/noriad"
NETWORK="noria_network"
DIR=".network"
NUM_VALIDATORS=2
RUN_AS_USER=" -u $(id -u):$(id -g) "
CONTAINER_HOME=" --home /tmp/noria "

### ARGS VALIDATION

if [[ $# -lt 1 ]]; then
  echo "Usage: $0 <init|start|stop|clean>"
  exit 1
fi

if [[ $1 != "init" && $1 != "start" && $1 != "stop" && $1 != "clean" ]]; then
  echo "Usage: $0 <init|start|stop|clean>"
  exit 1
fi

### CLEANUP

if [[ $1 == "clean" ]]; then
  stop_nodes
  rm_nodes
  rm -rf $DIR
  echo -e "\nCleaned up\n"
  exit 0
fi

### STOPPING NODES

if [[ $1 == "stop" ]]; then
  stop_nodes
  echo -e "\nNodes stopped\n"
  exit 0
fi

### STARTING NODES

if [[ $1 == "start" ]]; then

  if [[ $(docker network ls | grep -c "$NETWORK") -eq 0 ]]; then
    echo -e "\nCreating docker network..."
    docker network create --subnet=172.173.0.0/16 $NETWORK > /dev/null 2>&1
  fi

  # Start nodes
  docker run $RUN_AS_USER -d --name val1 -v $(pwd)/$DIR/val1:/tmp/noria -v $(pwd):/app --net $NETWORK -p 1317:1317 -p 26657:26657 $IMG $CONTAINER_HOME start > /dev/null 2>&1
  docker run $RUN_AS_USER -d --name val2 -v $(pwd)/$DIR/val2:/tmp/noria -v $(pwd):/app --net $NETWORK -p 1318:1317 -p 26658:26657 $IMG $CONTAINER_HOME start > /dev/null 2>&1

  echo -e "\nStarting relayer and creating channel...\n"

  # Give some time for the nodes to start
  sleep 8

  # Create relayer channel
  docker run --rm -v $(pwd)/$DIR/val1/relayer:/root/\.hermes --net $NETWORK noria/hermes create channel --order unordered --a-chain oasis-3 --b-chain oasis-4 --a-port transfer --b-port transfer --new-client-connection --yes

  # Start relayer
  docker run -d --name hermes -v $(pwd)/$DIR/val1/relayer:/root/\.hermes --net $NETWORK noria/hermes start > /dev/null 2>&1

  echo -e "\nRelayer started\n"
  echo -e "\nFollow logs with: docker logs hermes -f\n"

  exit 0
fi

### INITIALIZING NODES

if [[ $1 == "init" ]]; then
  docker stop hermes > /dev/null 2>&1
  docker container rm hermes > /dev/null 2>&1
  rm -rf $DIR
  for ((i = 1; i <= $NUM_VALIDATORS; i++)); do
    docker stop val$i > /dev/null 2>&1
    docker container rm val$i > /dev/null 2>&1
    mkdir -p $DIR/val$i
  done
  mkdir -p $DIR/val1/relayer

  # chain oasis-3
  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val1:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME init val1 --chain-id oasis-3 > /dev/null 2>&1
  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val1:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME config keyring-backend test
  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val1:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME config chain-id oasis-3
  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val1:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME keys add val1 --output json >$DIR/val1/key.json
  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val1:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME keys add relayer --output json >$DIR/val1/relayer/relayer_key.json
  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val1:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME genesis add-genesis-account val1 1000000000ucrd,1000000000unoria
  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val1:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME genesis add-genesis-account relayer 1000000000ucrd,1000000000unoria
  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val1:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME genesis gentx val1 100000000unoria --chain-id oasis-3 --commission-rate 0.1 --commission-max-rate 0.2 --commission-max-change-rate 0.01 > /dev/null 2>&1

  update_genesis $DIR/val1/config
  update_configs $DIR/val1/config

  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val1:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME genesis collect-gentxs > /dev/null 2>&1

  # chain oasis-4
  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val2:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME init val2 --chain-id oasis-4 > /dev/null 2>&1
  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val2:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME config keyring-backend test
  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val2:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME config chain-id oasis-4
  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val2:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME keys add val2 --output json >$DIR/val2/key.json
  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val1:/tmp/val1 -v $(pwd)/$DIR/val2:/tmp/noria -v $(pwd):/app --entrypoint /bin/bash $IMG -c "cat /tmp/val1/relayer/relayer_key.json  | jq -r '.mnemonic' | noriad $CONTAINER_HOME keys add relayer --recover"
  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val2:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME genesis add-genesis-account val2 1000000000ucrd,1000000000unoria
  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val2:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME genesis add-genesis-account relayer 1000000000ucrd,1000000000unoria
  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val2:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME genesis gentx val2 100000000unoria --chain-id oasis-4 --commission-rate 0.1 --commission-max-rate 0.2 --commission-max-change-rate 0.01 > /dev/null 2>&1

  update_genesis $DIR/val2/config
  update_configs $DIR/val2/config

  docker run $RUN_AS_USER --rm -v $(pwd)/$DIR/val2:/tmp/noria -v $(pwd):/app $IMG $CONTAINER_HOME genesis collect-gentxs > /dev/null 2>&1

  # Hermes Relayer
  echo -e $HERMES_CONFIG >$DIR/val1/relayer/config.toml
  docker run --rm -v $(pwd)/$DIR/val1/relayer:/root/\.hermes noria/hermes keys add --chain oasis-3 --key-file /root/.hermes/relayer_key.json > /dev/null 2>&1
  docker run --rm -v $(pwd)/$DIR/val1/relayer:/root/\.hermes noria/hermes keys add --chain oasis-4 --key-file /root/.hermes/relayer_key.json > /dev/null 2>&1
  docker run --rm -v $(pwd)/$DIR/val1/relayer:/root/\.hermes --entrypoint /bin/bash noria/hermes -c "chown -R $(id -u):$(id -g) /root/.hermes"

  echo -e "\nNodes initialized\n"
  echo -e "\nTo start: $0 start\n"
  exit 0
fi
