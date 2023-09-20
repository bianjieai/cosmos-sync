package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	ctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/tendermint/tendermint/types"

	recordtypes "gitlab.bianjie.ai/cschain/cschain/modules/ibc/applications/record/types"
	bcos "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/bcos/types"
	brochain "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/brochain/types"
	fabric "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/fabric/types"
	tendermint "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/tendermint/types"
	wutong "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/wutong/types"

	enccodec "github.com/tharsis/ethermint/encoding/codec"
)

var (
	AppModules []module.AppModuleBasic
	Encodecfg  params.EncodingConfig
)

func GetTxDecoder() sdk.TxDecoder {
	return Encodecfg.TxConfig.TxDecoder()
}

func GetMarshaler() codec.Codec {
	return Encodecfg.Marshaler
}

func GetSigningTx(txBytes types.Tx) (signing.Tx, error) {
	Tx, err := GetTxDecoder()(txBytes)
	if err != nil {
		return nil, err
	}
	signingTx := Tx.(signing.Tx)
	return signingTx, nil
}

func RegisterAppModules(module ...module.AppModuleBasic) {
	AppModules = append(AppModules, module...)
}

// MakeEncodingConfig 初始化账户地址前缀
func MakeEncodingConfig() {
	var cdc = codec.NewLegacyAmino()

	interfaceRegistry := ctypes.NewInterfaceRegistry()
	brochain.RegisterInterfaces(interfaceRegistry)
	fabric.RegisterInterfaces(interfaceRegistry)
	tendermint.RegisterInterfaces(interfaceRegistry)
	wutong.RegisterInterfaces(interfaceRegistry)
	bcos.RegisterInterfaces(interfaceRegistry)
	recordtypes.RegisterInterfaces(interfaceRegistry)

	moduleBasics := module.NewBasicManager(AppModules...)
	std.RegisterInterfaces(interfaceRegistry)
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)

	enccodec.RegisterLegacyAminoCodec(cdc)
	moduleBasics.RegisterLegacyAminoCodec(cdc)
	enccodec.RegisterInterfaces(interfaceRegistry)
	moduleBasics.RegisterInterfaces(interfaceRegistry)

	Encodecfg = params.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             cdc,
	}
}
