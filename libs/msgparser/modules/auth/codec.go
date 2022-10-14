package auth

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

func init() {
	codec.RegisterAppModules(auth.AppModuleBasic{})
}
