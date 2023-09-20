package nft

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	"github.com/irisnet/irismod/modules/nft"
)

func init() {
	codec.RegisterAppModules(nft.AppModuleBasic{})
}
