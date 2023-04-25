package types

type CoinmasterMsg struct {
	Coinmaster *CoinmasterSubMsg `json:"coinmaster,omitempty"`
}

type MsgCustomCoinmasterMint struct {
	Amount string `json:"amount,omitempty"`
}

type MsgCustomCoinmasterBurn struct {
	Amount string `json:"amount,omitempty"`
}
type CoinmasterSubMsg struct {
	Mint *MsgCustomCoinmasterMint `json:"mint,omitempty"`
	Burn *MsgCustomCoinmasterBurn `json:"burn,omitempty"`
}

type CoinmasterQuery struct {
	Coinmaster *CoinmasterSubQuery `json:"coinmaster"`
}

// See https://github.com/CosmWasm/token-bindings/blob/main/packages/bindings/src/query.rs
type CoinmasterSubQuery struct {
	Params *GetParams `json:"params"`
}

type GetParams struct{}

type ParamsProperties struct {
	Minters string `json:"minters"`
	Denoms  string `json:"denoms"`
}

type ParamsResponse struct {
	Params ParamsProperties `json:"params"`
}
