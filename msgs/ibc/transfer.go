package ibc

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
)

type DocMsgTransfer struct {
	SourcePort       string      `bson:"source_port"`
	SourceChannel    string      `bson:"source_channel"`
	Token            models.Coin `bson:"token"`
	Sender           string      `bson:"sender"`
	Receiver         string      `bson:"receiver"`
	TimeoutHeight    Height      `bson:"timeout_height"`
	TimeoutTimestamp uint64      `bson:"timeout_timestamp"`
}

func (m *DocMsgTransfer) GetType() string {
	return MsgTypeIBCTransfer
}

func (m *DocMsgTransfer) BuildMsg(v interface{}) {
	msg := v.(*MsgTransfer)
	m.SourcePort = msg.SourcePort
	m.SourceChannel = msg.SourceChannel
	m.Sender = msg.Sender
	m.Receiver = msg.Receiver
	m.TimeoutTimestamp = msg.TimeoutTimestamp
	m.TimeoutHeight = loadHeight(msg.TimeoutHeight)
	m.Token = models.BuildDocCoin(msg.Token)
}

func (m *DocMsgTransfer) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgTransfer
	)
	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Sender, msg.Receiver)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
