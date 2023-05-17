#!/bin/bash
#
# Test an upgrade on your local environment

HEIGHT=$1
UPGRADE_NAME="$2"
BINARY_DIR=".noria"
CHAIN_ID="oasis-3"
DENOM="unoria"
GAS_PRICE="0.0025"
GAS_PRICE_DENOM="ucrd"
FROM_KEY_NAME="me"
DAEMON_NAME="noriad"
DAEMON_HOME="$HOME/$BINARY_DIR"

# submit upgrade proposal
echo Submitting upgrade proposal
$DAEMON_NAME tx gov submit-legacy-proposal software-upgrade $UPGRADE_NAME \
  --title "Upgrade to $UPGRADE_NAME" \
  --description "Upgrade to $UPGRADE_NAME" \
  --upgrade-info='{}' \
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
  --no-validate

sleep 2

# vote on the proposal
echo "\n\nVoting YES on the upgrade proposal\n\n"

PROPOSAL_ID=$($DAEMON_NAME q gov proposals limit 1 --reverse --output json | jq '.proposals[0].id | tonumber')
# vote on the proposal
$DAEMON_NAME tx gov vote $PROPOSAL_ID yes \
  --from $FROM_KEY_NAME \
  --chain-id $CHAIN_ID \
  --home $DAEMON_HOME \
  --yes \
  --gas-prices $GAS_PRICE$GAS_PRICE_DENOM \
  --gas auto \
  --node tcp://localhost:26657 \
  --gas-adjustment 1.5
