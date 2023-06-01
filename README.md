# Welcome to Noria

The Noria Network is the only stable-denominated crypto network. It’s envisioned to be the most accessible decentralized network for stable money.

## Architectural Overview

The Noria Network is a Cosmos SDK blockchain powered by CometBFT for the networking and consensus layers. The network utilizes modified Cosmos modules to manage the blockchain, including:

- Minting - Responsible for minting tokens and managing inflation
- CosmWasm - A runtime for WebAssembly Smart Contracts
- Staking - Manages changes to validators
- Governance - Modified governance module to split technical consensus from social consensus, and manages the governance system for token holders
- IBC - Inter-Blockchain Communication protocol that enables asset transfers across Cosmos blockchains
- Token Factory - Allows permissionless creation of new native tokens [[link]](https://github.com/noria-net/token-factory)
- Coinmaster - Allows permissioned creation of new native tokens
- IBC Hooks - Allows receiving GMP (General Message Passing) messages from other blockchains to execute smart contracts [[link]](https://github.com/noria-net/ibc-hooks)

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
