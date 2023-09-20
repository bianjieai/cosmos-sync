package record

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	"github.com/irisnet/irismod/modules/record"
)

func init() {
	codec.RegisterAppModules(record.AppModuleBasic{})
}
