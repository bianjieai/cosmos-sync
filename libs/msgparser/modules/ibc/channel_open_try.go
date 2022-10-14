package ibc

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/utils"
)

type DocMsgChannelOpenTry struct {
	PortId              string  `bson:"port_id"`
	PreviousChannelId   string  `bson:"previous_channel_id"`
	Channel             Channel `bson:"channel"`
	CounterpartyVersion string  `bson:"counterparty_version"`
	ProofInit           string  `bson:"proof_init"`
	ProofHeight         Height  `bson:"proof_height"`
	Signer              string  `bson:"signer"`
}

func (m *DocMsgChannelOpenTry) GetType() string {
	return MsgTypeChannelOpenTry
}

func (m *DocMsgChannelOpenTry) BuildMsg(v interface{}) {
	msg := v.(*MsgChannelOpenTry)
	m.Signer = msg.Signer
	m.PortId = msg.PortId
	m.PreviousChannelId = msg.PreviousChannelId
	m.Channel = loadChannel(msg.Channel)
	m.CounterpartyVersion = msg.CounterpartyVersion
	m.ProofInit = utils.MarshalJsonIgnoreErr(msg.ProofInit)
	m.ProofHeight = loadHeight(msg.ProofHeight)
}

func (m *DocMsgChannelOpenTry) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgChannelOpenTry
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
