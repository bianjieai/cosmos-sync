package ibc

import (
	"fmt"
	cdc "github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	icoreclient "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	icorechannel "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
)

func loadPacket(packet icorechannel.Packet) Packet {
	return Packet{
		Sequence:           int64(packet.Sequence),
		SourcePort:         packet.SourcePort,
		SourceChannel:      packet.SourceChannel,
		DestinationPort:    packet.DestinationPort,
		DestinationChannel: packet.DestinationChannel,
		Data:               UnmarshalPacketData(packet.GetData()),
		TimeoutTimestamp:   int64(packet.TimeoutTimestamp),
		TimeoutHeight:      loadHeight(packet.TimeoutHeight)}
}

func UnmarshalPacketData(bytesdata []byte) PacketData {
	var (
		packetData FungibleTokenPacketData
	)
	err := cdc.GetMarshaler().UnmarshalJSON(bytesdata, &packetData)
	if err != nil {
		fmt.Println(err.Error())
	}

	return PacketData{
		Denom:    packetData.Denom,
		Amount:   packetData.Amount,
		Receiver: packetData.Receiver,
		Sender:   packetData.Sender,
	}
}

func loadHeight(height icoreclient.Height) Height {
	return Height{
		RevisionNumber: int64(height.RevisionNumber),
		RevisionHeight: int64(height.RevisionHeight)}
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
	Amount   string `bson:"amount" json:"amount"`
	Sender   string `bson:"sender" json:"sender"`
	Receiver string `bson:"receiver" json:"receiver"`
}
