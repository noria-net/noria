package coinmaster

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/noria-net/noria/testutil/sample"
	coinmastersimulation "github.com/noria-net/noria/x/coinmaster/simulation"
	"github.com/noria-net/noria/x/coinmaster/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = coinmastersimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCoinmasterMint = "op_weight_msg_mint"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCoinmasterMint int = 100

	opWeightMsgCoinmasterBurn = "op_weight_msg_burn"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCoinmasterBurn int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	coinmasterGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&coinmasterGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	coinmasterParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMinters), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(coinmasterParams.Minters))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCoinmasterMint int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCoinmasterMint, &weightMsgCoinmasterMint, nil,
		func(_ *rand.Rand) {
			weightMsgCoinmasterMint = defaultWeightMsgCoinmasterMint
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCoinmasterMint,
		coinmastersimulation.SimulateMsgCoinmasterMint(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCoinmasterBurn int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCoinmasterBurn, &weightMsgCoinmasterBurn, nil,
		func(_ *rand.Rand) {
			weightMsgCoinmasterBurn = defaultWeightMsgCoinmasterBurn
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCoinmasterBurn,
		coinmastersimulation.SimulateMsgCoinmasterBurn(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
