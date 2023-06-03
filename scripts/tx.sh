#!/bin/bash

###############################
# Usage: tx.sh <cosmos tx...>
###############################
# This script is used to submit a transaction to the blockchain
# and wait for it to be included in a block.
###############################

if [ $# -eq 0 ]; then
  echo "Usage: tx.sh <cosmos tx...>"
  exit 1
fi

HASH=$("$@" --output json -y | jq '.txhash' | sed -r 's/\"//g')
echo -e "\nWaiting for $HASH to be included in a block..."

CODE="-1"
TIMEOUT=10
MSG="..."

while [ $TIMEOUT -gt 0 ]; do
  $1 q tx "$HASH" >/dev/null 2>&1
  RESULT=$?
  if [ $RESULT -ne 0 ]; then
    sleep 1
    TIMEOUT=$((TIMEOUT - 1))
    if [ $TIMEOUT -eq 0 ]; then
      break
    fi
    continue
  fi

  MSG=$($1 q tx "$HASH" --output json)
  CODE=$(echo "$MSG" | jq '.code | tonumber')
  TIMEOUT=0
done

if [ "$CODE" -ne 0 ]; then
  RAWLOG=$(echo "$MSG" | jq '.raw_log')
  echo -e "\nError [code $CODE]: $RAWLOG"
  exit 1
else
  echo -e "\nSuccess: $HASH"
  exit 0
fi
