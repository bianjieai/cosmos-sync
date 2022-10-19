package ibc

import (
	"fmt"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules/types"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/utils"
	icoreclient "github.com/okex/exchain/libs/ibc-go/modules/core/02-client/types"
	icorechannel "github.com/okex/exchain/libs/ibc-go/modules/core/04-channel/types"
	"strconv"
)

func loadPacket(packet icorechannel.Packet) Packet {
	return Packet{
		Sequence:           int64(packet.Sequence),
		SourcePort:         packet.SourcePort,
		SourceChannel:      packet.SourceChannel,
		DestinationPort:    packet.DestinationPort,
		DestinationChannel: packet.DestinationChannel,
		Data:               UnmarshalPacketData(packet.GetData()),
		TimeoutTimestamp:   ConvertUint64ToInt64(packet.TimeoutTimestamp),
		TimeoutHeight:      loadHeight(packet.TimeoutHeight)}
}

func UnmarshalPacketData(bytesdata []byte) PacketData {
	var (
		packetData FungibleTokenPacketData
		data       PacketData
	)
	err := codec.GetCodec().UnmarshalJSON(bytesdata, &packetData)
	if err != nil {
		fmt.Println(err.Error())
	}
	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(packetData), &data)
	return data
}

func loadHeight(height icoreclient.Height) Height {
	return Height{
		RevisionNumber: ConvertUint64ToInt64(height.RevisionNumber),
		RevisionHeight: ConvertUint64ToInt64(height.RevisionHeight)}
}

func ConvertUint64ToInt64(data uint64) int64 {
	dataStr := strconv.FormatUint(data, 10)
	if len(dataStr) <= 19 {
		return int64(data)
	}

	dataStr = dataStr[:19]
	value, err := strconv.ParseInt(dataStr, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	return value
}
