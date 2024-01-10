package keeper

import (
	"context"

	"github.com/bsjohnson01/tokenfactory/types"
)

func (k *Keeper) InitGenesis(ctx context.Context, data *types.GenesisState) error {
	for _, elem := range data.DenomList {
		k.SetDenom(ctx, elem)
	}
	if err := k.SetParams(ctx, data.Params); err != nil {
		return err
	}

	return nil
}

func (k *Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DenomList = k.GetAllDenom(ctx)

	return genesis, nil
}
