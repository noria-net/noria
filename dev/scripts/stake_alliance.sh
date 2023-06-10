#!/bin/bash

KEY_NAME="me"
BINARY_DIR=".noria"
CHAIN_ID="oasis-3"
AMOUNT=$1
DENOM=$2
GAS_PRICE_DENOM="ucrd"
GAS_PRICE="0.0025"
NODE="http://127.0.0.1:26657/"
export DAEMON_NAME="noriad"
export DAEMON_HOME="$HOME/$BINARY_DIR"

# Staking amount must be set
if [ -z "$1" ]; then
  echo "Parameter amount is missing"
  exit 1
fi

# Denom must be set
if [ -z "$2" ]; then
  echo "Parameter denom is missing"
  exit 1
fi

exe() {
  echo "EXECUTING: $@"
  ./dev/scripts/tx.sh "$@"
}

# delegate to the validator through alliance
VAL=$($DAEMON_NAME q staking validators --output json | jq '.validators[0].operator_address' | sed 's/\"//g')
exe noriad tx alliance delegate $VAL $AMOUNT$DENOM --from me --gas auto --gas-adjustment 1.75 --gas-prices $GAS_PRICE$GAS_PRICE_DENOM
