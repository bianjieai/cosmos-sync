package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
)

type DocMsgSubmitMisbehaviour struct {
	ClientId     string `bson:"client_id"`
	Misbehaviour string `bson:"misbehaviour"`
	Signer       string `bson:"signer"`
}

func (m *DocMsgSubmitMisbehaviour) GetType() string {
	return MsgTypeSubmitMisbehaviourClient
}

func (m *DocMsgSubmitMisbehaviour) BuildMsg(v interface{}) {
	msg := v.(*MsgSubmitMisbehaviour)
	m.ClientId = msg.ClientId
	m.Misbehaviour = ConvertAny(msg.Misbehaviour)
	m.Signer = msg.Signer
}

func (m *DocMsgSubmitMisbehaviour) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgSubmitMisbehaviour
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
