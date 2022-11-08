package msgparser

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	commoncodec "github.com/kaifei-bianjie/common-parser/codec"
	"github.com/tharsis/ethermint/encoding"
)

// MakeEncodingConfig 初始化账户地址前缀
func MakeEncodingConfig() {
	moduleBasics := module.NewBasicManager(commoncodec.AppModules...)
	commoncodec.Encodecfg = encoding.MakeConfig(moduleBasics)
}
