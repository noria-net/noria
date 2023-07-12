#!/bin/bash

AMOUNT=$1
NEW_DENOM=$2
KEY_NAME="val1"
BINARY_DIR=".noria"
CHAIN_ID="oasis-3"
DENOM="unoria"
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

# submit parameter change proposal
# create-alliance denom reward-weight reward-weight-min reward-weight-max consensus-weight consensus-cap take-rate reward-change-rate reward-change-interval
exe $DAEMON_NAME tx gov submit-legacy-proposal create-alliance $NEW_DENOM 1 0 1 1 0.1 0 1 1s \
  --deposit 10000000unoria \
  --from "$KEY_NAME" \
  --chain-id "$CHAIN_ID" \
  --home "$DAEMON_HOME" \
  --node "$NODE" \
  --yes \
  --gas-prices $GAS_PRICE$GAS_PRICE_DENOM \
  --gas auto \
  --gas-adjustment 1.75

PROPOSAL_ID=$($DAEMON_NAME q gov proposals limit 1 --reverse --output json --home $DAEMON_HOME --node $NODE | jq '.proposals[0].id | tonumber')

# vote on the proposal
exe $DAEMON_NAME tx gov vote $PROPOSAL_ID yes --from $KEY_NAME --chain-id $CHAIN_ID --home $DAEMON_HOME --node $NODE --yes --gas-prices $GAS_PRICE$GAS_PRICE_DENOM --gas auto --gas-adjustment 1.75

sleep 5

# delegate to the validator through alliance
VAL=$($DAEMON_NAME q staking validators --output json | jq '.validators[] | select(.description.moniker | contains("val2")).operator_address' | sed 's/\"//g')
exe noriad tx alliance delegate $VAL $AMOUNT$NEW_DENOM --from $KEY_NAME --chain-id $CHAIN_ID --home $DAEMON_HOME --node $NODE --yes --gas-prices $GAS_PRICE$GAS_PRICE_DENOM --gas auto --gas-adjustment 1.75


# VAL=$($DAEMON_NAME q staking validators --output json | jq '.validators[] | select(.description.moniker | contains("val3")).operator_address' | sed 's/\"//g')
# exe noriad tx alliance delegate $VAL $AMOUNT$NEW_DENOM --from $KEY_NAME --chain-id $CHAIN_ID --home $DAEMON_HOME --node $NODE --yes --gas-prices $GAS_PRICE$GAS_PRICE_DENOM --gas auto --gas-adjustment 1.75
