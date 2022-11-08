package msgparser

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	commoncodec "github.com/kaifei-bianjie/common-parser/codec"
	"github.com/tharsis/ethermint/encoding"
)

func AdaptEthermintEncodingConfig() {
	moduleBasics := module.NewBasicManager(commoncodec.AppModules...)
	commoncodec.Encodecfg = encoding.MakeConfig(moduleBasics)
}
