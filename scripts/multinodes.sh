#!/bin/bash

# Requires noria image: make build-image

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

### VARS

IMG="noria/noriad"
NETWORK="noria_multinode"
DIR=".multinode"
NUM_VALIDATORS=$2

### ARGS VALIDATION

if [[ $# -lt 2 ]]; then
  echo "Usage: $0 <init|start|stop|clean> <number of validators>"
  exit 1
fi

if [[ $1 != "init" && $1 != "start" && $1 != "stop" && $1 != "clean" ]]; then
  echo "Usage: $0 <init|start|stop|clean> <number of validators>"
  exit 1
fi

if [[ $NUM_VALIDATORS -lt 1 && $NUM_VALIDATORS -gt 10 ]]; then
  echo "Keep the number of validators between 1 and 10"
  exit 1
fi

### CLEANUP

if [[ $1 == "clean" ]]; then
  rm -rf $DIR
  for ((i = 1; i <= $NUM_VALIDATORS; i++)); do
    docker stop val$i
    docker container rm val$i
  done
  echo -e "\nCleaned up\n"
  exit 0
fi

### STOPPING NODES

if [[ $1 == "stop" ]]; then
  for ((i = 1; i <= $NUM_VALIDATORS; i++)); do
    docker stop val$i
  done
  echo -e "\nNodes stopped\n"
  exit 0
fi

### STARTING NODES

if [[ $1 == "start" ]]; then

  if [[ $(docker network ls | grep -c "$NETWORK") -eq 0 ]]; then
    echo -e "\nCreating docker network..."
    docker network create --subnet=172.172.0.0/16 $NETWORK
  fi

  docker run -d --name val1 -v $(pwd)/$DIR/val1:/root/\.noria -v $(pwd):/app --net $NETWORK -p 1317:1317 -p 26657:26657 $IMG start
  for ((i = 2; i <= $NUM_VALIDATORS; i++)); do
    KEY=val$i
    docker run --name $KEY -d -v $(pwd)/$DIR/$KEY:/root/\.noria -v $(pwd):/app --net $NETWORK $IMG start
  done
  echo -e "\nNetwork started\n"
  echo -e "\nTo stop: $0 stop $NUM_VALIDATORS\n"
  exit 0
fi

### INITIALIZING NODES

if [[ $1 == "init" ]]; then
  rm -rf $DIR
  for ((i = 1; i <= $NUM_VALIDATORS; i++)); do
    docker stop val$i
    docker container rm val$i
  done

  docker run -v $(pwd)/$DIR/val1:/root/\.noria -v $(pwd):/app $IMG init val1 --chain-id oasis-3
  docker run -v $(pwd)/$DIR/val1:/root/\.noria -v $(pwd):/app $IMG config keyring-backend test
  docker run -v $(pwd)/$DIR/val1:/root/\.noria -v $(pwd):/app $IMG keys add val1 --output json >$DIR/val1/key.json
  docker run -v $(pwd)/$DIR/val1:/root/\.noria -v $(pwd):/app $IMG genesis add-genesis-account val1 1000000000ucrd,1000000000unoria
  docker run -v $(pwd)/$DIR/val1:/root/\.noria -v $(pwd):/app $IMG genesis gentx val1 100000000unoria --chain-id oasis-3 --commission-rate 0.1 --commission-max-rate 0.2 --commission-max-change-rate 0.01
  ADDR=$(docker run -v $(pwd)/$DIR/val1:/root/\.noria -v $(pwd):/app $IMG keys show val1 -a)
  PEER1=$(docker run -v $(pwd)/$DIR/val1:/root/\.noria -v $(pwd):/app $IMG tendermint show-node-id)

  update_genesis $DIR/val1/config
  update_configs $DIR/val1/config

  for ((i = 2; i <= $NUM_VALIDATORS; i++)); do

    KEY=val$i

    docker run -v $(pwd)/$DIR/$KEY:/root/\.noria -v $(pwd):/app $IMG init $KEY --chain-id oasis-3
    docker run -v $(pwd)/$DIR/$KEY:/root/\.noria -v $(pwd):/app $IMG config keyring-backend test
    docker run -v $(pwd)/$DIR/$KEY:/root/\.noria -v $(pwd):/app $IMG keys add $KEY --output json >$DIR/$KEY/key.json
    update_configs $DIR/$KEY/config
    sed -i.bak 's/^persistent_peers\ =\ .*/persistent_peers\ =\ \"'$PEER1'@val1:26656\"/g' $DIR/$KEY/config/config.toml

    cp $DIR/val1/config/genesis.json $DIR/$KEY/config/genesis.json
    docker run -v $(pwd)/$DIR/$KEY:/root/\.noria -v $(pwd):/app $IMG genesis add-genesis-account $KEY 1000000000ucrd,1000000000unoria
    docker run -v $(pwd)/$DIR/$KEY:/root/\.noria -v $(pwd):/app $IMG genesis gentx $KEY 100000000unoria --chain-id oasis-3 --commission-rate 0.1 --commission-max-rate 0.2 --commission-max-change-rate 0.01
    cp $DIR/$KEY/config/genesis.json $DIR/val1/config/genesis.json
    cp $DIR/$KEY/config/gentx/* $DIR/val1/config/gentx/

  done

  docker run -v $(pwd)/$DIR/val1:/root/\.noria -v $(pwd):/app $IMG genesis collect-gentxs

  for ((i = 2; i <= $NUM_VALIDATORS; i++)); do
    KEY=val$i
    cp $DIR/val1/config/genesis.json $DIR/$KEY/config/genesis.json
  done

  chmod -R 777 $DIR

  echo -e "\nNodes initialized\n"
  echo -e "\nTo start: $0 start $NUM_VALIDATORS\n"
  exit 0
fi
