#!/bin/bash

HEIGHT=$1
KEY_NAME=$2
UPGRADE_NAME="v0.1.2"
BINARY_DIR=".noria"
CHAIN_ID="oasis-3"
DENOM="unoria"
GAS_PRICE_DENOM="ucrd"
GAS_PRICE="0.0025"
NODE="TBA"
export DAEMON_NAME="noriad"
export DAEMON_HOME="$HOME/$BINARY_DIR"

BINARY_DOWNLOAD_URL="https://github.com/noria-net/noria/releases/download/v1.2.0/noria_linux_amd64.tar.gz?checksum=sha256:28420d2d6b20136a97f256cc21d09dd02c61a46b3ef00914f1d0a5c54db6e15c"

# Block height and sender key name must be set
if [ -z "$1" ] && [ -z "$2" ]; then
  echo "Block height and/or sender key name is missing"
  exit 1
fi

# submit upgrade proposal
$DAEMON_NAME tx gov submit-proposal software-upgrade $UPGRADE_NAME \
  --title "Upgrade to $UPGRADE_NAME" \
  --description "Upgrade to $UPGRADE_NAME" \
  --upgrade-info="{\"binaries\":{\"linux/amd64\":\"$BINARY_DOWNLOAD_URL\"}}" \
  --deposit 10000000$DENOM \
  --upgrade-height $HEIGHT \
  --from $KEY_NAME \
  --chain-id $CHAIN_ID \
  --home $DAEMON_HOME \
  --node $NODE \
  --yes \
  --gas-prices $GAS_PRICE$GAS_PRICE_DENOM \
  --gas auto \
  --gas-adjustment 1.5 \
  --broadcast-mode block

PROPOSAL_ID=$($DAEMON_NAME q gov proposals limit 1 --reverse --output json --home $DAEMON_HOME --node $NODE | jq '.proposals[0].proposal_id | tonumber')

# vote on the proposal
$DAEMON_NAME tx gov vote $PROPOSAL_ID yes --from $KEY_NAME --chain-id $CHAIN_ID --home $DAEMON_HOME --node $NODE --yes --gas-prices $GAS_PRICE$GAS_PRICE_DENOM --gas auto --gas-adjustment 1.5 --broadcast-mode block

