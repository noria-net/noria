# Welcome to Noria

The Noria Network is the only stable-denominated crypto network. It’s envisioned to be the most accessible decentralized network for stable money.

## Architectural Overview
The Noria Network is a Cosmos SDK blockchain powered by Tendermint for the networking and consensus layers. The network utilizes modified Cosmos modules to manage the blockchain, including:
- Minting - Responsible for minting tokens and managing inflation
- CosmWasm - A runtime for WebAssembly Smart Contracts
- Staking - Manages changes to validators
- Governance - Modified governance module to split technical consensus from social consensus, and manages the governance system for token holders
- IBC - Inter-Blockchain Communication protocol that enables asset transfers across Cosmos blockchains

The combination of different innovations in the way that Noria Network is designed will allow the network to offer something unique to developers and communities:
- low, stable, and predictable gas fees
- focus on practicality, ease of use, and real world applications
- delegators will elect governance focused representatives or representative groups for operational decision making

## Information for Validators, Developers, and more

Please visit Noria's documentation site for information on
- Running a Noria validator
- Building on Noria
- Developing, testing, and deploying CosmWasm smart contracts

Noria Documentation: https://noria.notaku.site/

# Releasing Updates
## Noria Release Script
This script creates a release for the Noria project on GitHub, including a zipped binary of the noriad executable and a SHA256 hash of the file. The release notes include the version number and the SHA256 hash.

### Usage

1. Ensure that you have set up a [personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) on GitHub with permission to create releases and upload assets.
1. Clone the [Noria repository](https://github.com/noria-net/noria/)
1. Make the script executable by running `chmod +x ./scripts/release_noria.sh`
1. Run the script with `./scripts/release_noria.sh`
1. Follow the prompts to enter your GitHub username and repository name.

The script will create a new release on GitHub with the tag name, name, and release notes specified in the script. The zipped binary will be uploaded as an asset to the release.

### Troubleshooting
* If you receive a 401 Unauthorized error when running the script, double-check that your personal access token has permission to create releases and upload assets.

* If you receive a 404 Not Found error when running the script, double-check that you have entered the correct GitHub username and repository name.

* If you receive any other errors, try running the script with the --verbose flag to see more detailed output.

That's it! If you have any questions or issues with the script, feel free to open an issue on the Noria repository.
