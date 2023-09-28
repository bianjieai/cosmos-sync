package node

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
)

type DocMsgCreateValidator struct {
	Name        string `bson:"name"`
	Certificate string `bson:"certificate"`
	Power       int64  `bson:"power"`
	Description string `bson:"description"`
	Operator    string `bson:"operator"`
}

func (m *DocMsgCreateValidator) GetType() string {
	return MsgTypeCreateValidator
}

func (m *DocMsgCreateValidator) BuildMsg(v interface{}) {
	msg := v.(*MsgNodeCreate)

	m.Name = msg.Name
	m.Certificate = msg.Certificate
	m.Power = msg.Power
	m.Description = msg.Description
	m.Operator = msg.Operator
}

func (m *DocMsgCreateValidator) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgNodeCreate)
	addrs = append(addrs, msg.Operator)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
