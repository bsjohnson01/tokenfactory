package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	//registry.RegisterImplementations((*sdk.Msg)(nil),
	//	&MsgCreateDenom{},
	//	&MsgUpdateDenom{},
	//)
	//registry.RegisterImplementations((*sdk.Msg)(nil),
	//	&MsgMintAndSendTokens{},
	//)
	//registry.RegisterImplementations((*sdk.Msg)(nil),
	//	&MsgUpdateOwner{},
	//)
	//// this line is used by starport scaffolding # 3
	//
	//registry.RegisterImplementations((*sdk.Msg)(nil),
	//	&MsgUpdateParams{},
	//)
	//msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
