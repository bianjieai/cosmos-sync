package node

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
)

type DocMsgRemoveValidator struct {
	Id       string `bson:"id"`
	Operator string `bson:"operator"`
}

func (m *DocMsgRemoveValidator) GetType() string {
	return MsgTypeRemoveValidator
}

func (m *DocMsgRemoveValidator) BuildMsg(v interface{}) {
	msg := v.(*MsgNodeRemove)

	m.Id = msg.Id
	m.Operator = msg.Operator
}

func (m *DocMsgRemoveValidator) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgNodeRemove)
	addrs = append(addrs, msg.Operator)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
