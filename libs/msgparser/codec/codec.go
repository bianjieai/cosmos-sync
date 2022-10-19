package codec

import (
	okexchaincodec "github.com/okex/exchain/app/codec"
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	_ "github.com/okex/exchain/libs/cosmos-sdk/crypto"
	ibctx "github.com/okex/exchain/libs/cosmos-sdk/types/ibc-adapter"
	"github.com/okex/exchain/libs/cosmos-sdk/types/module"
	"github.com/okex/exchain/libs/cosmos-sdk/x/auth"
	cosmoscryptocodec "github.com/okex/exchain/libs/cosmos-sdk/x/auth/ibc-tx"
	stdtx "github.com/okex/exchain/libs/cosmos-sdk/x/auth/types"
)

var (
	appModules = []module.AppModuleBasic{auth.AppModuleBasic{}}
	codecProxy *codec.CodecProxy
	txDecoder  ibctx.IbcTxDecoder
)

func InitTxDecoder() {
	moduleBasics := module.NewBasicManager(appModules...)
	cdc := okexchaincodec.MakeCodec(moduleBasics)
	interfaceReg := okexchaincodec.MakeIBC(moduleBasics)
	protoCodec := codec.NewProtoCodec(interfaceReg)
	codecProxy = codec.NewCodecProxy(protoCodec, cdc)
	txDecoder = cosmoscryptocodec.IbcTxDecoder(codecProxy.GetProtocMarshal())
	return
}

func GetCodec() *codec.ProtoCodec {
	return codecProxy.GetProtocMarshal()
}

func GetSigningTx(txBytes []byte) (*stdtx.IbcTx, error) {
	tx, err := txDecoder(txBytes)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func RegisterAppModules(module ...module.AppModuleBasic) {
	appModules = append(appModules, module...)
}
