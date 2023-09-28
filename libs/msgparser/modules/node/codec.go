package node

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	"github.com/bianjieai/iritamod/modules/node"
)

func init() {
	codec.RegisterAppModules(
		node.AppModuleBasic{},
	)
}
