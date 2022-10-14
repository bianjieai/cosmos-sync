package ibc

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	ibctransfer "github.com/cosmos/ibc-go/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/modules/core"
)

func init() {
	codec.RegisterAppModules(
		ibc.AppModuleBasic{},
		ibctransfer.AppModuleBasic{},
	)
}
