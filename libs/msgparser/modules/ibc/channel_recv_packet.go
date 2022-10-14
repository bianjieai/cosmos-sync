package ibc

import (
	"fmt"
	codec "github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules/ibc/types"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/utils"
	icoreclient "github.com/okex/exchain/libs/ibc-go/modules/core/02-client/types"
	icorechannel "github.com/okex/exchain/libs/ibc-go/modules/core/04-channel/types"
	"strconv"
)

type DocMsgRecvPacket struct {
	PacketId        string `bson:"packet_id"`
	Packet          Packet `bson:"packet"`
	ProofCommitment string `bson:"proof_commitment"`
	ProofHeight     Height `bson:"proof_height"`
	Signer          string `bson:"signer"`
}

func (m *DocMsgRecvPacket) GetType() string {
	return MsgTypeRecvPacket
}

func (m *DocMsgRecvPacket) BuildMsg(v interface{}) {
	msg := v.(*MsgRecvPacket)
	m.Signer = msg.Signer
	m.ProofHeight = loadHeight(msg.ProofHeight)
	m.ProofCommitment = utils.MarshalJsonIgnoreErr(msg.ProofCommitment)
	m.Packet = loadPacket(msg.Packet)
	m.PacketId = fmt.Sprintf("%v%v%v%v%v", msg.Packet.SourcePort, msg.Packet.SourceChannel,
		msg.Packet.DestinationPort, msg.Packet.DestinationChannel, msg.Packet.Sequence)

}

func (m *DocMsgRecvPacket) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgRecvPacket
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	packetData := UnmarshalPacketData(msg.Packet.GetData())
	addrs = append(addrs, msg.Signer, packetData.Receiver, packetData.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}

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
	dataStr := fmt.Sprint(data)
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
