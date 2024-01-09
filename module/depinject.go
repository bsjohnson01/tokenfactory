package module

import (
	"context"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	modulev1 "github.com/bsjohnson01/tokenfactory/api/module"
	"github.com/bsjohnson01/tokenfactory/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

var (
	_ appmodule.AppModule       = (*AppModule)(nil)
	_ appmodule.HasBeginBlocker = (*AppModule)(nil)
	_ appmodule.HasEndBlocker   = (*AppModule)(nil)
)

func (m AppModule) IsOnePerModuleType() {

}

func (m AppModule) IsAppModule() {

}

func (m AppModule) BeginBlock(_ context.Context) error {
	return nil
}

func (m AppModule) EndBlock(_ context.Context) error {
	return nil
}

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	StoreService store.KVStoreService
	Cdc          codec.Codec
	Config       *modulev1.Module
	Logger       log.Logger

	//AccountKeeper types.AccountKeeper
	//BankKeeper    types.BankKeeper
}

type ModuleOutputs struct {
	depinject.Out

	TokenfactoryKeeper keeper.Keeper
	Module             appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}
	k := keeper.NewKeeper(
		in.Cdc,
		in.StoreService,
		in.Logger,
		authority.String(),
		//in.AccountKeeper,
		//in.BankKeeper,
	)

	m := NewAppModule(
		in.Cdc,
		k,
		//in.AccountKeeper,
		//in.BankKeeper,
	)

	return ModuleOutputs{TokenfactoryKeeper: k, Module: m}
}
