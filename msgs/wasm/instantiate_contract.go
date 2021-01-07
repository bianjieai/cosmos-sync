package wasm

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
)

type DocMsgInstantiateContract struct {
	// Sender is the that actor that signed the messages
	Sender string `bson:"sender"`
	// Admin is an optional address that can execute migrations
	Admin string `bson:"admin"`
	// CodeID is the reference to the stored WASM code
	CodeID uint64 `bson:"code_id"`
	// Label is optional metadata to be stored with a contract instance.
	Label string `bson:"label"`
	// InitMsg json encoded message to be passed to the contract on instantiation
	InitMsg string `bson:"init_msg"`
	// InitFunds coins that are transferred to the contract on instantiation
	InitFunds []models.Coin `bson:"init_funds"`
}

func (m *DocMsgInstantiateContract) GetType() string {
	return MsgTypeInstantiateContract
}

func (m *DocMsgInstantiateContract) BuildMsg(v interface{}) {
	msg := v.(*MsgInstantiateContract)
	m.Sender = msg.Sender
	m.Admin = msg.Admin
	m.CodeID = msg.CodeID
	m.Label = msg.Label
	m.InitMsg = string(msg.InitMsg)
	m.InitFunds = models.BuildDocCoins(msg.InitFunds)

}

func (m *DocMsgInstantiateContract) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgInstantiateContract
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Sender, msg.Admin)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
