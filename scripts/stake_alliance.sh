#!/bin/bash

KEY_NAME="me"
BINARY_DIR=".noria"
CHAIN_ID="oasis-3"
DENOM=$1
GAS_PRICE_DENOM="ucrd"
GAS_PRICE="0.0025"
NODE="http://127.0.0.1:26657/"
export DAEMON_NAME="noriad"
export DAEMON_HOME="$HOME/$BINARY_DIR"

exe() { echo "EXECUTING: $@" ; ./scripts/tx.sh "$@" ; }

# delegate to the validator through alliance 
VAL=$($DAEMON_NAME q staking validators --output json | jq '.validators[0].operator_address' | sed 's/\"//g')
exe noriad tx alliance delegate $VAL 100000000$DENOM --from me --gas auto --gas-adjustment 1.75 --gas-prices $GAS_PRICE$GAS_PRICE_DENOM