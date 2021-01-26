package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
)

type DocMsgAcknowledgement struct {
	Packet          Packet `bson:"packet"`
	Acknowledgement string `bson:"acknowledgement"`
	ProofAcked      string `bson:"proof_acked"`
	ProofHeight     Height `bson:"proof_height"`
	Signer          string `bson:"signer"`
}

func (m *DocMsgAcknowledgement) GetType() string {
	return MsgTypeAcknowledgement
}

func (m *DocMsgAcknowledgement) BuildMsg(v interface{}) {

	msg := v.(*MsgAcknowledgement)
	m.Signer = msg.Signer
	m.ProofHeight = loadHeight(msg.ProofHeight)
	m.Acknowledgement = utils.MarshalJsonIgnoreErr(msg.Acknowledgement)
	m.ProofAcked = utils.MarshalJsonIgnoreErr(msg.ProofAcked)
	m.Packet = loadPacket(msg.Packet)

}

func (m *DocMsgAcknowledgement) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgAcknowledgement
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
