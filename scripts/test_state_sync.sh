#!/bin/bash

echo -e "\nStopping any running noriad instance..."
sudo killall noriad

BIN="noriad"
HOME1="/tmp/noria1"
CONFIG1="/tmp/noria1/config"
RPC1="26657"
FLAG1="--home $HOME1"

HOME2="/tmp/noria2"
CONFIG2="/tmp/noria2/config"
RPC2="27657"
GRPC2="10090"
GRPCWEB2="10091"
ADDR2="27658"
P2P2="27656"
PPROF2="7060"
FLAG2="--home $HOME2"

CHAIN_ID="oasis-3"
DENOM="unoria"
GAS_PRICE_DENOM="ucrd"

echo -e "\nRemoving previous config folders ($HOME1, $HOME2)"
rm -rf $HOME1 $HOME2

echo "Setting keyring to \"test\""
$BIN $FLAG1 config keyring-backend test
$BIN $FLAG2 config keyring-backend test

echo "Setting chain-id to \"$CHAIN_ID\""
$BIN $FLAG1 config chain-id $CHAIN_ID
$BIN $FLAG2 config chain-id $CHAIN_ID

$BIN $FLAG1 keys add me
$BIN $FLAG2 keys add me

$BIN $FLAG1 init me --overwrite --chain-id $CHAIN_ID
$BIN $FLAG2 init me --overwrite --chain-id $CHAIN_ID

$BIN $FLAG1 add-genesis-account me 1000000000$GAS_PRICE_DENOM,1000000000$DENOM

$BIN $FLAG1 gentx me 100000000$DENOM --chain-id oasis-3 --commission-rate 0.1 --commission-max-rate 0.2 --commission-max-change-rate 0.01

$BIN $FLAG1 collect-gentxs

sed -i.bak "s/stake/$DENOM/g" $CONFIG1/genesis.json

rm $CONFIG1/genesis.json.bak

sed -i.bak 's/^timeout_commit\ =\ .*/timeout_commit\ =\ \"0.5s\"/g' $CONFIG1/config.toml
rm $CONFIG1/config.toml.bak

sed -i.bak "s/^minimum-gas-prices\ =\ .*/minimum-gas-prices\ =\ \"0.0025$GAS_PRICE_DENOM\"/g" $CONFIG1/app.toml

sed -i.bak 's/^enable\ =\ false/enable\ =\ true/g' $CONFIG1/app.toml
sed -i.bak 's/^swagger\ =\ false/swagger\ =\ true/g' $CONFIG1/app.toml
sed -i.bak '/Rosetta API/{n; s/true/false/}' $CONFIG1/app.toml
rm $CONFIG1/app.toml.bak

cp $CONFIG1/genesis.json $CONFIG1/app.toml $CONFIG1/config.toml $CONFIG2/

sed -i.bak 's/^snapshot-interval\ =\ 0/snapshot-interval\ =\ 10/g' $CONFIG1/app.toml
sed -i.bak 's/^snapshot-keep-recent\ =\ 2/snapshot-keep-recent\ =\ 10/g' $CONFIG1/app.toml

$BIN $FLAG1 --home $HOME1 start --pruning=nothing >/dev/null 2>&1 &

sleep 3

$BIN $FLAG1 --home $HOME1 tx bank send $($BIN $FLAG1 keys show me -a) $($BIN $FLAG2 keys show me -a) 1ucrd --from me --fees 500ucrd -y

sleep 3

SNAP_RPC="http://localhost:$RPC1"

BLOCK_HEIGHT=5
TRUST_HASH=$(curl -s "$SNAP_RPC/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)

sed -i.bak -E "s|^(enable[[:space:]]+=[[:space:]]+).*$|\1true| ; \
s|^(rpc_servers[[:space:]]+=[[:space:]]+).*$|\1\"$SNAP_RPC,$SNAP_RPC\"| ; \
s|^(trust_height[[:space:]]+=[[:space:]]+).*$|\1$BLOCK_HEIGHT| ; \
s|^(trust_hash[[:space:]]+=[[:space:]]+).*$|\1\"$TRUST_HASH\"| ; \
s|^(seeds[[:space:]]+=[[:space:]]+).*$|\1\"\"|" $CONFIG2/config.toml

sed -i.bak 's/^address\ =\ \"tcp:\/\/0.0.0.0:1317\"/address\ =\ \"tcp:\/\/0.0.0.0:2317\"/g' $CONFIG2/app.toml

NODE_ID=$($BIN $FLAG1 --home $HOME1 tendermint show-node-id)
PEER1="$NODE_ID@127.0.0.1:26656"

$BIN $FLAG2 --home $HOME2 start --address $ADDR2 \
  --grpc.address "localhost:$GRPC2" \
  --grpc-web.address "localhost:$GRPCWEB2" \
  --rpc.laddr "tcp://localhost:$RPC2" \
  --p2p.laddr "tcp://localhost:$P2P2" \
  --p2p.persistent_peers "$PEER1" \
  --rpc.pprof_laddr "localhost:$PPROF2" \
  --pruning=nothing
