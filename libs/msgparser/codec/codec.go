package codec

import (
	okexchaincodec "github.com/okex/exchain/app/codec"
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	_ "github.com/okex/exchain/libs/cosmos-sdk/crypto"
	"github.com/okex/exchain/libs/cosmos-sdk/types/module"
	okmodule "github.com/okex/exchain/libs/cosmos-sdk/types/module"
	"github.com/okex/exchain/libs/cosmos-sdk/x/auth"
	cosmoscryptocodec "github.com/okex/exchain/libs/cosmos-sdk/x/auth/ibc-tx"
	stdtx "github.com/okex/exchain/libs/cosmos-sdk/x/auth/types"
)

var (
	appOkModules = []okmodule.AppModuleBasic{auth.AppModuleBasic{}}
	codecProxy   *codec.CodecProxy
)

func InitTxDecoder() {
	moduleBasics := module.NewBasicManager(appOkModules...)
	cdc := okexchaincodec.MakeCodec(moduleBasics)
	interfaceReg := okexchaincodec.MakeIBC(moduleBasics)
	protoCodec := codec.NewProtoCodec(interfaceReg)
	codecProxy = codec.NewCodecProxy(protoCodec, cdc)
	return
}

func GetCodec() *codec.ProtoCodec {
	return codecProxy.GetProtocMarshal()
}

func GetSigningTx(txBytes []byte) (*stdtx.IbcTx, error) {
	tx, err := cosmoscryptocodec.IbcTxDecoder(codecProxy.GetProtocMarshal())(txBytes)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func RegisterAppOkModules(module ...okmodule.AppModuleBasic) {
	appOkModules = append(appOkModules, module...)
}
