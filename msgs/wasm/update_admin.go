package wasm

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

type DocMsgUpdateAdmin struct {
	// Sender is the that actor that signed the messages
	Sender string `bson:"sender"`
	// NewAdmin address to be set
	NewAdmin string `bson:"new_admin"`
	// Contract is the address of the smart contract
	Contract string `bson:"contract"`
}

func (m *DocMsgUpdateAdmin) GetType() string {
	return MsgTypeUpdateAdmin
}

func (m *DocMsgUpdateAdmin) BuildMsg(v interface{}) {
	msg := v.(*MsgUpdateAdmin)
	m.Sender = msg.Sender
	m.Contract = msg.Contract
	m.NewAdmin = msg.NewAdmin

}

func (m *DocMsgUpdateAdmin) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgUpdateAdmin
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Sender, msg.NewAdmin, msg.Contract)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
