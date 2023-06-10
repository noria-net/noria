# Chain Upgrades

It is good practice to test chain upgrades locally before sending them for devnet or testnet. This document describes how to do that.

## Usage

Execute the following command:  
> `make FROM="<FROM>" TO="<TO>" NAME="<NAME>" run-noria-upgrade`

__FROM__: The version of the chain you want to upgrade from.  
__TO__: The version of the chain you want to upgrade to. `current` can be used to upgrade to version in your current working directory.

_Both of these parameters can be tag names or commit hashes._

__NAME__: The name of the upgrade used by the proposal.
