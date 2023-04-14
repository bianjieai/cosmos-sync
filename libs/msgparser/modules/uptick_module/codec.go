package uptick_module

import (
	"github.com/UptickNetwork/uptick/x/erc20"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
)

func init() {
	codec.RegisterAppModules(
		erc20.AppModuleBasic{},
	)
}
