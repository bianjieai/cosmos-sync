package crisis

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/bianjieai/irita-sync/msgs"

)

type DocMsgVerifyInvariant struct {
	Sender              string `bson:"sender"`
	InvariantModuleName string `bson:"invariant_module_name" yaml:"invariant_module_name"`
	InvariantRoute      string `bson:"invariant_route" yaml:"invariant_route"`
}

func (m *DocMsgVerifyInvariant) GetType() string {
	return MsgTypeVerifyInvariant
}

func (m *DocMsgVerifyInvariant) BuildMsg(v interface{}) {
	var msg MsgVerifyInvariant
	data, _ := json.Marshal(v)
	json.Unmarshal(data, &msg)

}

func (m *DocMsgVerifyInvariant) HandleTxMsg(msg sdk.Msg) MsgDocInfo {

	var (
		addrs []string
	)

	addrs = append(addrs, m.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
