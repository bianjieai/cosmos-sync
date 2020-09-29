package cdc

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/irisnet/irismod/modules/nft"
	"github.com/irisnet/irismod/modules/record"
	"github.com/irisnet/irismod/modules/service"
	"github.com/irisnet/irismod/modules/token"
	"github.com/irisnet/irismod/modules/htlc"
	"github.com/irisnet/irismod/modules/coinswap"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"gitlab.bianjie.ai/irita-pro/iritamod/modules/identity"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	ctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/std"
)

var (
	encodecfg    params.EncodingConfig
	moduleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		service.AppModuleBasic{},
		nft.AppModuleBasic{},
		record.AppModuleBasic{},
		token.AppModuleBasic{},
		gov.AppModuleBasic{},
		staking.AppModuleBasic{},
		distribution.AppModuleBasic{},
		slashing.AppModuleBasic{},
		evidence.AppModuleBasic{},
		crisis.AppModuleBasic{},
		identity.AppModuleBasic{},
		htlc.AppModuleBasic{},
		coinswap.AppModuleBasic{},
	)
)

// 初始化账户地址前缀
func init() {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := ctypes.NewInterfaceRegistry()
	moduleBasics.RegisterInterfaces(interfaceRegistry)
	sdk.RegisterInterfaces(interfaceRegistry)
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, std.DefaultPublicKeyCodec{}, tx.DefaultSignModes)

	encodecfg = params.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}

func GetTxDecoder() sdk.TxDecoder {
	return encodecfg.TxConfig.TxDecoder()
}

func GetAmino() *codec.LegacyAmino {
	return encodecfg.Amino
}
