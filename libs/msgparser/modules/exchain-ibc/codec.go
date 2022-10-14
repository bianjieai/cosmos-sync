package exchain_ibc

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/okchain-codec"
	ibctransfer "github.com/okex/exchain/libs/ibc-go/modules/apps/transfer"
	ibc "github.com/okex/exchain/libs/ibc-go/modules/core"
)

func init() {
	codec.RegisterAppOkModules(
		ibc.AppModuleBasic{},
		ibctransfer.AppModuleBasic{},
	)
}
