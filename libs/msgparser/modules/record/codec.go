package record

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	"gitlab.cschain.tech/csmod/modules/record"
)

func init() {
	codec.RegisterAppModules(record.AppModuleBasic{})
}
