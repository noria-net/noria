#!/bin/bash

DIR=".multinode"

./dev/scripts/multinodes.sh init 5
for ((i = 1; i <= 5; i++)); do
  sed -i.bak 's/"signed_blocks_window": "[^"]*"/"signed_blocks_window": "10"/g' $DIR/val$i/config/genesis.json
  sed -i.bak 's/"min_signed_per_window": "[^"]*"/"min_signed_per_window": "0.3"/g' $DIR/val$i/config/genesis.json
  sed -i.bak 's/"unbonding_time": "[^"]*"/"unbonding_time": "10s"/g' $DIR/val$i/config/genesis.json
  sed -i.bak 's/"slash_receiver": "[^"]*"/"slash_receiver": "noria1h9szngpyzvhng3mjgq2n8u943ulv3cfsper7dl"/g' $DIR/val$i/config/genesis.json
done
./dev/scripts/multinodes.sh start 5
# sleep 5
# docker stop val4
