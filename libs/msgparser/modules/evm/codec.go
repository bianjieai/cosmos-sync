package evm

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	"github.com/tharsis/ethermint/x/evm"
)

func init() {
	codec.RegisterAppModules(
		evm.AppModuleBasic{},
	)
}
