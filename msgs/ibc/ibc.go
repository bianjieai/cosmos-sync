package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
	"github.com/cosmos/cosmos-sdk/types"
	icoreclient "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	icorechannel "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
)

func HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch v.Type() {
	case new(MsgRecvPacket).Type():
		docMsg := DocMsgRecvPacket{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgTransfer).Type():
		docMsg := DocMsgTransfer{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgCreateClient).Type():
		docMsg := DocMsgCreateClient{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgUpdateClient).Type():
		docMsg := DocMsgUpdateClient{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgUpgradeClient).Type():
		docMsg := DocMsgUpgradeClient{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgSubmitMisbehaviour).Type():
		docMsg := DocMsgSubmitMisbehaviour{}
		msgDocInfo = docMsg.HandleTxMsg(v)

	case new(MsgConnectionOpenInit).Type():
		docMsg := DocMsgConnectionOpenInit{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgConnectionOpenTry).Type():
		docMsg := DocMsgConnectionOpenTry{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgConnectionOpenAck).Type():
		docMsg := DocMsgConnectionOpenAck{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgConnectionOpenConfirm).Type():
		docMsg := DocMsgConnectionOpenConfirm{}
		msgDocInfo = docMsg.HandleTxMsg(v)

	case new(MsgChannelOpenInit).Type():
		docMsg := DocMsgChannelOpenInit{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgChannelOpenTry).Type():
		docMsg := DocMsgChannelOpenTry{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgChannelOpenAck).Type():
		docMsg := DocMsgChannelOpenAck{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgChannelOpenConfirm).Type():
		docMsg := DocMsgChannelOpenConfirm{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgChannelCloseInit).Type():
		docMsg := DocMsgChannelCloseInit{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgChannelCloseConfirm).Type():
		docMsg := DocMsgChannelCloseConfirm{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgTimeout).Type():
		docMsg := DocMsgTimeout{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgTimeoutOnClose).Type():
		docMsg := DocMsgTimeoutOnClose{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgAcknowledgement).Type():
		docMsg := DocMsgAcknowledgement{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	default:
		ok = false
	}
	return msgDocInfo, ok
}

func loadPacket(packet icorechannel.Packet) Packet {
	return Packet{
		Sequence:           packet.Sequence,
		SourcePort:         packet.SourcePort,
		SourceChannel:      packet.SourceChannel,
		DestinationPort:    packet.DestinationPort,
		DestinationChannel: packet.DestinationChannel,
		Data:               utils.MarshalJsonIgnoreErr(packet.Data),
		TimeoutTimestamp:   packet.TimeoutTimestamp,
		TimeoutHeight:      loadHeight(packet.TimeoutHeight)}
}

func loadHeight(height icoreclient.Height) Height {
	return Height{
		RevisionNumber: height.RevisionNumber,
		RevisionHeight: height.RevisionHeight}
}

func loadChannel(channel icorechannel.Channel) Channel {
	return Channel{
		State:          int32(channel.State),
		Ordering:       int32(channel.State),
		Version:        channel.Version,
		ConnectionHops: channel.ConnectionHops,
		Counterparty: ChannelCounterparty{
			ChannelId: channel.Counterparty.ChannelId,
			PortId:    channel.Counterparty.PortId,
		},
	}
}

type Channel struct {
	State          int32               `bson:"state"`
	Ordering       int32               `bson:"ordering"`
	Counterparty   ChannelCounterparty `bson:"counterparty"`
	ConnectionHops []string            `bson:"connection_hops"`
	Version        string              `bson:"version"`
}
type ChannelCounterparty struct {
	PortId    string `bson:"port_id"`
	ChannelId string `bson:"channel_id"`
}

type Height struct {
	RevisionNumber uint64 `bson:"revision_number"`
	RevisionHeight uint64 `bson:"revision_height"`
}

// Packet defines a type that carries data across different chains through IBC
type Packet struct {
	Sequence           uint64 `bson:"sequence"`
	SourcePort         string `bson:"source_port"`
	SourceChannel      string `bson:"source_channel"`
	DestinationPort    string `bson:"destination_port"`
	DestinationChannel string `bson:"destination_channel"`
	Data               string `bson:"data"`
	TimeoutHeight      Height `bson:"timeout_height"`
	TimeoutTimestamp   uint64 `bson:"timeout_timestamp"`
}