package ibc

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	tmjson "github.com/tendermint/tendermint/libs/json"
	pc "github.com/tendermint/tendermint/proto/tendermint/crypto"
	ibcrecord "gitlab.bianjie.ai/cschain/cschain/modules/ibc/applications/record"
	ibc "gitlab.bianjie.ai/cschain/cschain/modules/ibc/core"
)

func init() {
	tmjson.RegisterType((*pc.PublicKey_Sm2)(nil), "tendermint.crypto.PublicKey_Sm2")
	codec.RegisterAppModules(
		ibc.AppModuleBasic{},
		ibcrecord.AppModuleBasic{},
	)
}
