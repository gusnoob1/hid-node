package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDID{}, "ssi/CreateDID", nil)
	cdc.RegisterConcrete(&MsgUpdateDID{}, "ssi/UpdateDID", nil)
	cdc.RegisterConcrete(&MsgCreateSchema{}, "ssi/CreateSchema", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDID{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateDID{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSchema{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
