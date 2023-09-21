package service

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	"gitlab.cschain.tech/csmod/modules/service"
)

func init() {
	codec.RegisterAppModules(service.AppModuleBasic{})
}
