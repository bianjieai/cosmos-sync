package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
)

type DocMsgConnectionOpenConfirm struct {
	ConnectionId string `bson:"connection_id"`
	ProofAck     string `bson:"proof_ack"`
	ProofHeight  Height `bson:"proof_height"`
	Signer       string `bson:"signer"`
}

func (m *DocMsgConnectionOpenConfirm) GetType() string {
	return MsgTypeConnectionOpenConfirm
}

func (m *DocMsgConnectionOpenConfirm) BuildMsg(v interface{}) {
	msg := v.(*MsgConnectionOpenConfirm)
	m.Signer = msg.Signer
	m.ConnectionId = msg.ConnectionId
	m.ProofAck = utils.MarshalJsonIgnoreErr(msg.ProofAck)
	m.ProofHeight = loadHeight(msg.ProofHeight)
}

func (m *DocMsgConnectionOpenConfirm) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgConnectionOpenConfirm
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
