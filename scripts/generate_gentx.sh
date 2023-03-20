#!/bin/bash

# This script wipes your config folder (~/.noria),
# creates a new wallet named <moniker>
# and prepares everything to be able to generate a gentx
# to be part of the genesis validator set.

# Enter your validator moniker here
MONIKER="me"

# Customize your BINARY HOME
BINARY_HOME="$HOME/.noria"

# Customise your keyring
KEYRING="os"

# Customise your validator commission rates
COMMISSION_RATE=0.05
COMMISSION_MAX_RATE=0.2
COMMISSION_MAX_CHANGE_RATE=0.01

# DO NOT CHANGE THE FOLLOWING
CONFIG_HOME="$BINARY_HOME/config"
CHAIN_ID="oasis-3"
DENOM="unoria"
GAS_PRICE="0.0025"
GAS_PRICE_DENOM="ucrd"

echo -e "\nRemoving previous config folder ($BINARY_HOME)"
rm -rf $BINARY_HOME

echo "Setting keyring to $KEYRING"
noriad config keyring-backend $KEYRING

# Set the default chain to use
echo "Setting chain-id to \"$CHAIN_ID\""
noriad config chain-id $CHAIN_ID

# Create a new wallet named $MONIKER
noriad keys add $MONIKER

# Initialize a new genesis.json file
noriad init $MONIKER --chain-id $CHAIN_ID >/dev/null 2>&1

# Copy the repo genesis over to your $BINARY_HOME/config/genesis.json
cp ./genesis.json $CONFIG_HOME/genesis.json

# Add your freshly created account to the new chain genesis
noriad add-genesis-account $MONIKER 5000000$GAS_PRICE_DENOM,100000000000$DENOM >/dev/null 2>&1

echo -e "\n"

# Generate the genesis transaction to create a new validator
noriad gentx $MONIKER 100000000000$DENOM --chain-id oasis-3 --commission-rate $COMMISSION_RATE --commission-max-rate $COMMISSION_MAX_RATE --commission-max-change-rate $COMMISSION_MAX_CHANGE_RATE

# Copy local files to create pull request
cp $CONFIG_HOME/gentx/*.json ./gentx/
cp $CONFIG_HOME/genesis.json ./genesis.json

echo -e "\n\nGentx successfully created.\n"
echo -e "\nPlease submit your gentx via pull request to the Github repository and make sure to add your moniker in the title.\n"
