package service

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	"github.com/irisnet/irismod/modules/service"
)

func init() {
	codec.RegisterAppModules(service.AppModuleBasic{})
}
