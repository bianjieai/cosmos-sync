package codec

import (
	"github.com/bianjieai/cosmos-sync/libs/logger"
	cryptocodec "github.com/okex/exchain/app/crypto/ethsecp256k1"
	ethermint "github.com/okex/exchain/app/types"
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	"github.com/okex/exchain/libs/cosmos-sdk/codec/types"
	_ "github.com/okex/exchain/libs/cosmos-sdk/crypto"
	"github.com/okex/exchain/libs/cosmos-sdk/crypto/keys"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	"github.com/okex/exchain/libs/cosmos-sdk/types/module"
	okmodule "github.com/okex/exchain/libs/cosmos-sdk/types/module"
	"github.com/okex/exchain/libs/cosmos-sdk/x/auth"
	cosmoscryptocodec "github.com/okex/exchain/libs/cosmos-sdk/x/auth/ibc-tx"
	authTypes "github.com/okex/exchain/libs/cosmos-sdk/x/auth/types"
	"github.com/okex/exchain/libs/cosmos-sdk/x/auth/vesting"
)

var (
	appOkModules = []okmodule.AppModuleBasic{auth.AppModuleBasic{}}
	codecProxy   *codec.CodecProxy
)

//// 初始化账户地址前缀
//func MakeEncodingConfig() {
//	var cdc = codec.New()
//	cryptocodec.RegisterCrypto(cdc)
//
//	interfaceRegistry := ctypes.NewInterfaceRegistry()
//	moduleBasics := module.NewBasicManager(appOkModules...)
//	moduleBasics.RegisterInterfaces(interfaceRegistry)
//	std.RegisterInterfaces(interfaceRegistry)
//	marshaler := codec.NewProtoCodec(interfaceRegistry)
//	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)
//
//	encodecfg = params.EncodingConfig{
//		InterfaceRegistry: interfaceRegistry,
//		Marshaler:         marshaler,
//		TxConfig:          txCfg,
//		Amino:             cdc,
//	}
//}
func InitTxDecoder() {
	moduleBasics := module.NewBasicManager(appOkModules...)
	codecProxy, _ = MakeCodecSuit(moduleBasics)
}
func MakeCodec(bm module.BasicManager) *codec.Codec {
	cdc := codec.New()
	bm.RegisterCodec(cdc)
	vesting.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	cryptocodec.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	ethermint.RegisterCodec(cdc)
	keys.RegisterCodec(cdc) // temporary. Used to register keyring.Info

	return cdc
}

func MakeIBC(bm module.BasicManager) types.InterfaceRegistry {
	interfaceReg := types.NewInterfaceRegistry()
	bm.RegisterInterfaces(interfaceReg)
	cosmoscryptocodec.PubKeyRegisterInterfaces(interfaceReg)
	return interfaceReg
}

func MakeCodecSuit(bm module.BasicManager) (*codec.CodecProxy, types.InterfaceRegistry) {
	aminoCodec := MakeCodec(bm)
	interfaceReg := MakeIBC(bm)
	protoCdc := codec.NewProtoCodec(interfaceReg)
	return codec.NewCodecProxy(protoCdc, aminoCodec), interfaceReg
}

func GetSigningTx(txBytes []byte) (sdk.Tx, error) {
	if codecProxy == nil {
		logger.Fatal("InitTxDecoder not work")
	}
	var tx authTypes.StdTx
	err := codecProxy.UnmarshalBinaryLengthPrefixed(txBytes, &tx)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func RegisterAppOkModules(module ...okmodule.AppModuleBasic) {
	appOkModules = append(appOkModules, module...)
}
