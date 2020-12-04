package cdc

import (
	"github.com/bianjieai/iritamod/modules/admin"
	"github.com/bianjieai/iritamod/modules/identity"
	iritaparams "github.com/bianjieai/iritamod/modules/params"
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
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/irisnet/irismod/modules/coinswap"
	"github.com/irisnet/irismod/modules/htlc"
	"github.com/irisnet/irismod/modules/nft"
	"github.com/irisnet/irismod/modules/oracle"
	"github.com/irisnet/irismod/modules/random"
	"github.com/irisnet/irismod/modules/record"
	"github.com/irisnet/irismod/modules/service"
	"github.com/irisnet/irismod/modules/token"
	ibcrecord "gitlab.bianjie.ai/cschain/cschain/modules/ibc/applications/record"
	recordtypes "gitlab.bianjie.ai/cschain/cschain/modules/ibc/applications/record/types"
	ibc "gitlab.bianjie.ai/cschain/cschain/modules/ibc/core"
	bcos "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/bcos/types"
	brochain "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/brochain/types"
	fabric "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/fabric/types"
	tendermint "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/tendermint/types"
	wutong "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/wutong/types"
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
		oracle.AppModuleBasic{},
		random.AppModuleBasic{},
		ibc.AppModuleBasic{},
		ibcrecord.AppModuleBasic{},
		admin.AppModuleBasic{},
		iritaparams.AppModuleBasic{},
	)
)

// 初始化账户地址前缀
func init() {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := ctypes.NewInterfaceRegistry()
	brochain.RegisterInterfaces(interfaceRegistry)
	fabric.RegisterInterfaces(interfaceRegistry)
	tendermint.RegisterInterfaces(interfaceRegistry)
	wutong.RegisterInterfaces(interfaceRegistry)
	bcos.RegisterInterfaces(interfaceRegistry)
	recordtypes.RegisterInterfaces(interfaceRegistry)
	moduleBasics.RegisterInterfaces(interfaceRegistry)
	moduleBasics.RegisterLegacyAminoCodec(amino)
	std.RegisterInterfaces(interfaceRegistry)
	std.RegisterLegacyAminoCodec(amino)
	sdk.RegisterInterfaces(interfaceRegistry)
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)
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

func GetMarshaler() codec.Marshaler {
	return encodecfg.Marshaler
}
