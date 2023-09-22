package identity

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	"github.com/bianjieai/iritamod/modules/identity"
)

func init() {
	codec.RegisterAppModules(identity.AppModuleBasic{})
}
