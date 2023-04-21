# Coinmaster Module

Coinmaster is a custom Cosmos SDK module that allows for the creation, minting, and burning of custom tokens. It can be used in conjunction with the CosmWasm smart contract platform.

## Permission

The **Coinmaster** module has two guards in place:

- Whitelisted denominations, a comma-separated string of denominations. You can only mint/burn whitelisted denoms. The initial whitelist is set to an empty string, which allows minting and burning any denomination. This can be changed through governance with a `param-change` proposal such as:

```bash
{
  "title": "Update whitelisted denominations",
  "description": "Update whitelisted denominations",
  "changes": [
    {
      "subspace": "coinmaster",
      "key": "Denoms",
      "value": "ucrd,unoria,nwbtc"
    }
  ],
  "deposit": "1000000unoria"
}
```

- Whitelisted minters, a comma-sepratated string of account addresses (smart contract supported). Only Minters are allowed to mint/burn coins. This can be changed through governance with a `param-change` proposal such as:

```bash
{
  "title": "Update whitelisted minters",
  "description": "Update whitelisted minters",
  "changes": [
    {
      "subspace": "coinmaster",
      "key": "Minters",
      "value": "noria1qrh465lh5ygkaqu0nc2wdxfv5nkmwl3xlqf7jl"
    }
  ],
  "deposit": "1000000unoria"
}
```

The default value of the minter is an empty string, which allows anyone to mint/burn.

## Minting

```bash
noriad tx coinmaster mint \
1000000ucrd \
--from="myWalletName" \
--gas="auto" \
--gas-prices="0.0025ucrd" \
--gas-adjustment="1.5" \
--chain-id="oasis-3"
```

## Burning

```bash
noriad tx coinmaster burn \
1000000ucrd \
--from="myWalletName" \
--gas="auto" \
--gas-prices="0.0025ucrd" \
--gas-adjustment="1.5" \
--chain-id="oasis-3"
```

## Wasm Bindings

The Coinmaster module can be used in conjunction with CosmWasm smart contracts. The following is an example of how to use the module in a smart contract:

```rust
// Types

#[cw_serde]
pub enum ExecuteMsg {
    Mint { amount: String },
    Burn { amount: String },
}

#[cw_serde]
/// A number of Custom messages that can call into the Noria custom modules bindings
pub enum NoriaMsg {
    Coinmaster(CoinmasterMsg),
}

#[cw_serde]
/// A number of Custom messages that can call into the Coinmaster bindings
pub enum CoinmasterMsg {
    Mint { amount: String },
    Burn { amount: String },
}

impl From<NoriaMsg> for CosmosMsg<NoriaMsg> {
    fn from(msg: NoriaMsg) -> CosmosMsg<NoriaMsg> {
        CosmosMsg::Custom(msg)
    }
}

impl CustomMsg for NoriaMsg {}
```

```rust
// Execute mint
let mint_msg = NoriaMsg::Coinmaster(CoinmasterMsg::Mint {
    amount: String::from("1000000ucrd"),
});

let res = Response::new()
    .add_message(mint_msg);
```

```json
// expected payload
{
  "coinmaster": {
    "mint": {
      "amount": "1000000ucrd"
    }
  }
}
```

```rust
// Execute burn
let burn_msg = NoriaMsg::Coinmaster(CoinmasterMsg::Burn {
    amount: String::from("1000000ucrd"),
});

let res = Response::new()
    .add_message(burn_msg);
```

```json
// expected payload
{
  "coinmaster": {
    "burn": {
      "amount": "1000000ucrd"
    }
  }
}
```

## Protobuf

```protobuf

package norianet.noria.coinmaster;

// Protobuf value for MsgCoinmasterMint: /norianet.noria.coinmaster.MsgCoinmasterMint
// Protobuf value for MsgCoinmasterBurn: /norianet.noria.coinmaster.MsgCoinmasterBurn
```
