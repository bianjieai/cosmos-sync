package msgparser

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/evmos/ethermint/encoding"
	commoncodec "github.com/kaifei-bianjie/common-parser/codec"
)

func AdaptEthermintEncodingConfig() {
	moduleBasics := module.NewBasicManager(commoncodec.AppModules...)
	commoncodec.Encodecfg = encoding.MakeConfig(moduleBasics)
}
