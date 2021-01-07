package wasm

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
)

type DocMsgExecuteContract struct {
	// Sender is the that actor that signed the messages
	Sender string `bson:"sender"`
	// Contract is the address of the smart contract
	Contract string `bson:"contract"`
	// Msg json encoded message to be passed to the contract
	Msg []byte `bson:"msg"`
	// SentFunds coins that are transferred to the contract on execution
	SentFunds []models.Coin `bson:"sent_funds"`
}

func (m *DocMsgExecuteContract) GetType() string {
	return MsgTypeExecuteContract
}

func (m *DocMsgExecuteContract) BuildMsg(v interface{}) {
	msg := v.(*MsgExecuteContract)
	m.Sender = msg.Sender
	m.Contract = msg.Contract
	m.Msg = msg.Msg
	m.SentFunds = models.BuildDocCoins(msg.SentFunds)
}

func (m *DocMsgExecuteContract) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgExecuteContract
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
