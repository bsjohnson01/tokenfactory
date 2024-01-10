package module

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bsjohnson01/tokenfactory/keeper"
	"github.com/bsjohnson01/tokenfactory/types"
)

// InitGenesis initialized the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, elem := range genState.DenomList {
		k.SetDenom(ctx, elem)
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DenomList = k.GetAllDenom(ctx)

	return genesis
}
