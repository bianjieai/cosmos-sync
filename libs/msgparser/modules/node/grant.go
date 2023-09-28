package node

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
)

type DocMsgGrantNode struct {
	Name        string `bson:"name"`
	Certificate string `bson:"certificate"`
	Operator    string `bson:"operator"`
}

func (m *DocMsgGrantNode) GetType() string {
	return MsgTypeGrantNode
}

func (m *DocMsgGrantNode) BuildMsg(v interface{}) {
	msg := v.(*MsgNodeGrant)

	m.Name = msg.Name
	m.Certificate = msg.Certificate
	m.Operator = msg.Operator
}

func (m *DocMsgGrantNode) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgNodeGrant)
	addrs = append(addrs, msg.Operator)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
