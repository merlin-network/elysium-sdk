package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

// RegisterLegacyAminoCodec registers the necessary x/staking interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// legacy.RegisterAminoMsg(cdc, &MsgCreateValidator{}, "cosmos-sdk/MsgCreateValidator")
	// legacy.RegisterAminoMsg(cdc, &MsgEditValidator{}, "cosmos-sdk/MsgEditValidator")
	// legacy.RegisterAminoMsg(cdc, &MsgDelegate{}, "cosmos-sdk/MsgDelegate")
	// legacy.RegisterAminoMsg(cdc, &MsgUndelegate{}, "cosmos-sdk/MsgUndelegate")
	// legacy.RegisterAminoMsg(cdc, &MsgBeginRedelegate{}, "cosmos-sdk/MsgBeginRedelegate")
	// legacy.RegisterAminoMsg(cdc, &MsgCancelUnbondingDelegation{}, "cosmos-sdk/MsgCancelUnbondingDelegation")

	legacy.RegisterAminoMsg(cdc, &MsgTokenizeShares{}, "lsm/MsgTokenizeShares")
	legacy.RegisterAminoMsg(cdc, &MsgRedeemTokensforShares{}, "lsm/MsgRedeemTokensforShares")
	legacy.RegisterAminoMsg(cdc, &MsgTransferTokenizeShareRecord{}, "lsm/MsgTransferTokenizeShareRecord")

	// cdc.RegisterInterface((*isStakeAuthorization_Validators)(nil), nil)
	// legacy.RegisterAminoMsg(cdc, &StakeAuthorization_AllowList{}, "cosmos-sdk/StakeAuthorization/AllowList")
	// legacy.RegisterAminoMsg(cdc, &StakeAuthorization_DenyList{}, "cosmos-sdk/StakeAuthorization/DenyList")
	// legacy.RegisterAminoMsg(cdc, &StakeAuthorization{}, "cosmos-sdk/StakeAuthorization")
}

// RegisterInterfaces registers the x/staking interfaces types with the interface registry
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateValidator{},
		&MsgEditValidator{},
		&MsgDelegate{},
		&MsgUndelegate{},
		&MsgBeginRedelegate{},
		&MsgCancelUnbondingDelegation{},
		&MsgTokenizeShares{},
		&MsgRedeemTokensforShares{},
		&MsgTransferTokenizeShareRecord{},
	)
	registry.RegisterImplementations(
		(*authz.Authorization)(nil),
		&StakeAuthorization{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
