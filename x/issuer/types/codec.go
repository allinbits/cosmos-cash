package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

	// this line is used by starport scaffolding # 1
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

//nolint
func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateIssuer{},
		&MsgBurnToken{},
		&MsgMintToken{},
		&MsgPauseToken{},
		&MsgIssueUserCredential{},
		&MsgRevokeCredential{},
	)
}

var (
	// ModuleCdc codec used by the module (protobuf)
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
