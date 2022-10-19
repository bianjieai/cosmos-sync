package ibc

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules/types"
)

// MsgUpdateClient defines a message to update an IBC client
type DocMsgUpdateClient struct {
	ClientId string `bson:"client_id" yaml:"client_id"`
	Header   string `bson:"header" yaml:"header"`
	Signer   string `bson:"signer" yaml:"signer"`
}

func (m *DocMsgUpdateClient) GetType() string {
	return MsgTypeUpdateClient
}

func (m *DocMsgUpdateClient) BuildMsg(v interface{}) {
	msg := v.(*MsgUpdateClient)

	m.ClientId = msg.ClientId
	m.Signer = msg.Signer
}

func (m *DocMsgUpdateClient) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
	)
	msg := v.(*MsgUpdateClient)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}