package wasm

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

type DocMsgMigrateContract struct {
	// Sender is the that actor that signed the messages
	Sender string `bson:"sender"`
	// Contract is the address of the smart contract
	Contract string `bson:"contract"`
	// CodeID references the new WASM code
	CodeID uint64 `bson:"code_id"`
	// MigrateMsg json encoded message to be passed to the contract on migration
	MigrateMsg string `bson:"migrate_msg"`
}

func (m *DocMsgMigrateContract) GetType() string {
	return MsgTypeMigrateContract
}

func (m *DocMsgMigrateContract) BuildMsg(v interface{}) {
	msg := v.(*MsgMigrateContract)
	m.Sender = msg.Sender
	m.Contract = msg.Contract
	m.CodeID = msg.CodeID
	m.MigrateMsg = string(msg.MigrateMsg)

}

func (m *DocMsgMigrateContract) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgMigrateContract
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Sender, msg.Contract)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
