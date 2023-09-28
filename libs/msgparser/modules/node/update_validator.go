package node

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
)

type DocMsgUpdateValidator struct {
	Id          string `bson:"id"`
	Name        string `bson:"name"`
	Certificate string `bson:"certificate"`
	Power       int64  `bson:"power"`
	Description string `bson:"description"`
	Operator    string `bson:"operator"`
}

func (m *DocMsgUpdateValidator) GetType() string {
	return MsgTypeUpdateValidator
}

func (m *DocMsgUpdateValidator) BuildMsg(v interface{}) {
	msg := v.(*MsgNodeUpdate)

	m.Id = msg.Id
	m.Name = msg.Name
	m.Certificate = msg.Certificate
	m.Power = msg.Power
	m.Description = msg.Description
	m.Operator = msg.Operator
}

func (m *DocMsgUpdateValidator) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgNodeUpdate)
	addrs = append(addrs, msg.Operator)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
