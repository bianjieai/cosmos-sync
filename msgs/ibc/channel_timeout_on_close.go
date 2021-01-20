package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
)

type DocMsgTimeoutOnClose struct {
	Packet           Packet `bson:"packet"`
	ProofUnreceived  string `bson:"proof_unreceived"`
	ProofClose       string `bson:"proof_close"`
	ProofHeight      Height `bson:"proof_height"`
	NextSequenceRecv uint64 `bson:"next_sequence_recv"`
	Signer           string `bson:"signer"`
}

func (m *DocMsgTimeoutOnClose) GetType() string {
	return MsgTypeTimeoutOnClose
}

func (m *DocMsgTimeoutOnClose) BuildMsg(v interface{}) {
	msg := v.(*MsgTimeoutOnClose)
	m.Signer = msg.Signer
	m.NextSequenceRecv = msg.NextSequenceRecv
	m.ProofUnreceived = utils.MarshalJsonIgnoreErr(m.ProofUnreceived)
	m.ProofClose = utils.MarshalJsonIgnoreErr(m.ProofClose)
	m.Packet = loadPacket(msg.Packet)
	m.ProofHeight = loadHeight(msg.ProofHeight)
}

func (m *DocMsgTimeoutOnClose) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgTimeoutOnClose
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
