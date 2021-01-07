package wasm

import . "github.com/bianjieai/irita-sync/msgs"

// MsgClearAdmin removes any admin stored for a smart contract
type DocMsgClearAdmin struct {
	// Sender is the that actor that signed the messages
	Sender string `bson:"sender"`
	// Contract is the address of the smart contract
	Contract string `bson:"contract"`
}

func (m *DocMsgClearAdmin) GetType() string {
	return MsgTypeClearAdmin
}

func (m *DocMsgClearAdmin) BuildMsg(v interface{}) {
	msg := v.(*MsgClearAdmin)
	m.Sender = msg.Sender
	m.Contract = msg.Contract
}

func (m *DocMsgClearAdmin) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgClearAdmin
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
