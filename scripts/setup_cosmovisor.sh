#!/bin/bash

BINARY_DIR=".noria"
export DAEMON_NAME="noriad"
export DAEMON_HOME="$HOME/$BINARY_DIR"
export DAEMON_ALLOW_DOWNLOAD_BINARIES=true

if ! command -v go &>/dev/null; then
  echo "\n\ngolang could not be found"
  exit
fi

if ! command -v cosmovisor &>/dev/null; then
  echo "\n\ncosmovisor could not be found, installing..."
  go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@latest
  echo "cosmovisor installed"
fi

echo -e "\nCreating cosmovisor folder structure in $DAEMON_HOME"
mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin
cp `which $DAEMON_NAME` $DAEMON_HOME/cosmovisor/genesis/bin/
echo -e "Cosmovisor ready"
