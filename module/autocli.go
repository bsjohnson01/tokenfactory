package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/bsjohnson01/tokenfactory/api"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (m AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "DenomAll",
					Use:       "list-denom",
					Short:     "List all Denom",
				},
				{
					RpcMethod:      "Denom",
					Use:            "show-denom [id]",
					Short:          "Shows a Denom",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateDenom",
					Use:            "create-denom",
					Short:          "Create a new Denom",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "description"}, {ProtoField: "ticker"}, {ProtoField: "precision"}, {ProtoField: "url"}, {ProtoField: "maxSupply"}, {ProtoField: "canChangeMaxSupply"}, {ProtoField: "denom"}},
				},
				{
					RpcMethod:      "UpdateDenom",
					Use:            "update-denom",
					Short:          "Update Denom",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "description"}, {ProtoField: "url"}, {ProtoField: "maxSupply"}, {ProtoField: "canChangeMaxSupply"}, {ProtoField: "denom"}},
				},
				{
					RpcMethod:      "MintAndSendTokens",
					Use:            "mint-and-send-tokens [denom] [amount] [recipient]",
					Short:          "Send a MintAndSendTokens tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom"}, {ProtoField: "amount"}, {ProtoField: "recipient"}},
				},
				{
					RpcMethod:      "UpdateOwner",
					Use:            "update-owner [denom] [new-owner]",
					Short:          "Send a UpdateOwner tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "denom"}, {ProtoField: "newOwner"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
