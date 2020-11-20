package cdc

import (
	"github.com/cosmos/cosmos-sdk/codec"
	ctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer"
	ibc "github.com/cosmos/cosmos-sdk/x/ibc/core"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

var (
	encodecfg    params.EncodingConfig
	moduleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		//service.AppModuleBasic{},
		//nft.AppModuleBasic{},
		//htlc.AppModuleBasic{},
		//coinswap.AppModuleBasic{},
		//record.AppModuleBasic{},
		//token.AppModuleBasic{},
		gov.AppModuleBasic{},
		staking.AppModuleBasic{},
		distribution.AppModuleBasic{},
		slashing.AppModuleBasic{},
		evidence.AppModuleBasic{},
		crisis.AppModuleBasic{},
		//identity.AppModuleBasic{},
		//htlc.AppModuleBasic{},
		//coinswap.AppModuleBasic{},
		//oracle.AppModuleBasic{},
		//random.AppModuleBasic{},
		//ibc.AppModuleBasic{},
		//ibctransfer.AppModuleBasic{},
		//admin.AppModuleBasic{},
		//validator.AppModuleBasic{},
		//iritaslash.AppModuleBasic{},
		ibc.AppModuleBasic{},
		transfer.AppModuleBasic{},
	)
)

// 初始化账户地址前缀
func init() {
	var cdc = codec.NewLegacyAmino()

	interfaceRegistry := ctypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)

	encodecfg = params.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             cdc,
	}
	std.RegisterLegacyAminoCodec(encodecfg.Amino)
	std.RegisterInterfaces(encodecfg.InterfaceRegistry)
	moduleBasics.RegisterLegacyAminoCodec(encodecfg.Amino)
	moduleBasics.RegisterInterfaces(encodecfg.InterfaceRegistry)
}

func GetTxDecoder() sdk.TxDecoder {
	return encodecfg.TxConfig.TxDecoder()
}

func GetMarshaler() codec.Marshaler {
	return encodecfg.Marshaler
}