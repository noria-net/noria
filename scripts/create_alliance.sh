#!/bin/bash

NEW_DENOM=$1
KEY_NAME="me"
BINARY_DIR=".noria"
CHAIN_ID="oasis-3"
DENOM="unoria"
GAS_PRICE_DENOM="ucrd"
GAS_PRICE="0.0025"
NODE="http://127.0.0.1:26657/"
export DAEMON_NAME="noriad"
export DAEMON_HOME="$HOME/$BINARY_DIR"

# Block height and sender key name must be set
if [ -z "$1" ]; then
  echo "Parameter new denom is missing"
  exit 1
fi

exe() { echo "EXE\$ $@" ; "$@" ; }

# submit parameter change proposal
exe $DAEMON_NAME tx gov submit-legacy-proposal create-alliance $NEW_DENOM 1 0 1 0 1 1s \
  --deposit 10000000unoria \
  --from "$KEY_NAME" \
  --chain-id "$CHAIN_ID" \
  --home "$DAEMON_HOME" \
  --node "$NODE" \
  --yes \
  --gas-prices $GAS_PRICE$GAS_PRICE_DENOM \
  --gas auto \
  --gas-adjustment 1.75

sleep 2

PROPOSAL_ID=$($DAEMON_NAME q gov proposals limit 1 --reverse --output json --home $DAEMON_HOME --node $NODE | jq '.proposals[0].id | tonumber')

# vote on the proposal
exe $DAEMON_NAME tx gov vote $PROPOSAL_ID yes --from $KEY_NAME --chain-id $CHAIN_ID --home $DAEMON_HOME --node $NODE --yes --gas-prices $GAS_PRICE$GAS_PRICE_DENOM --gas auto --gas-adjustment 1.75
