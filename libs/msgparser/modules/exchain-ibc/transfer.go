package exchain_ibc

import (
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	models "github.com/bianjieai/cosmos-sync/libs/msgparser/types"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/utils"
)

type DocMsgTransfer struct {
	PacketId         string      `bson:"packet_id"`
	SourcePort       string      `bson:"source_port"`
	SourceChannel    string      `bson:"source_channel"`
	Token            models.Coin `bson:"token"`
	Sender           string      `bson:"sender"`
	Receiver         string      `bson:"receiver"`
	TimeoutHeight    Height      `bson:"timeout_height"`
	TimeoutTimestamp int64       `bson:"timeout_timestamp"`
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
	m.TimeoutTimestamp = int64(msg.TimeoutTimestamp)
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

type Height struct {
	RevisionNumber int64 `bson:"revision_number"`
	RevisionHeight int64 `bson:"revision_height"`
}

// Packet defines a type that carries data across different chains through IBC
type Packet struct {
	Sequence           int64      `bson:"sequence"`
	SourcePort         string     `bson:"source_port"`
	SourceChannel      string     `bson:"source_channel"`
	DestinationPort    string     `bson:"destination_port"`
	DestinationChannel string     `bson:"destination_channel"`
	Data               PacketData `bson:"data"`
	TimeoutHeight      Height     `bson:"timeout_height"`
	TimeoutTimestamp   int64      `bson:"timeout_timestamp"`
}

//FungibleTokenPacketData
type PacketData struct {
	Denom    string `bson:"denom" json:"denom"`
	Amount   int64  `bson:"amount" json:"amount"`
	Sender   string `bson:"sender" json:"sender"`
	Receiver string `bson:"receiver" json:"receiver"`
}
