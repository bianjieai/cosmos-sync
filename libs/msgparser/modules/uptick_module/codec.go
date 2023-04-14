package uptick_module

import (
	collection "github.com/UptickNetwork/uptick/x/collection/module"
	"github.com/UptickNetwork/uptick/x/erc20"
	"github.com/UptickNetwork/uptick/x/erc721"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
)

func init() {
	codec.RegisterAppModules(
		erc20.AppModuleBasic{},
		erc721.AppModuleBasic{},
		collection.AppModuleBasic{},
	)
}
