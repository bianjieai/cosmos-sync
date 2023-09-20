package bank

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

func init() {
	codec.RegisterAppModules(bank.AppModuleBasic{})
}
