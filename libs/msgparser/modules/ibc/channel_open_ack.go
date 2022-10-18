package ibc

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/utils"
)

type DocMsgChannelOpenAck struct {
	PortId                string `bson:"port_id"`
	ChannelId             string `bson:"channel_id"`
	CounterpartyChannelId string `bson:"counterparty_channel_id"`
	CounterpartyVersion   string `bson:"counterparty_version"`
	ProofTry              string `bson:"proof_try"`
	ProofHeight           Height `bson:"proof_height"`
	Signer                string `bson:"signer"`
}

func (m *DocMsgChannelOpenAck) GetType() string {
	return MsgTypeChannelOpenAck
}

func (m *DocMsgChannelOpenAck) BuildMsg(v interface{}) {
	msg := v.(*MsgChannelOpenAck)
	m.Signer = msg.Signer
	m.PortId = msg.PortId
	m.ChannelId = msg.ChannelId
	m.CounterpartyChannelId = msg.CounterpartyChannelId
	m.CounterpartyVersion = msg.CounterpartyVersion
	m.ProofTry = utils.MarshalJsonIgnoreErr(msg.ProofTry)
	m.ProofHeight = loadHeight(msg.ProofHeight)

}

func (m *DocMsgChannelOpenAck) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgChannelOpenAck
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
