package ibc

import (
	"fmt"
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules/types"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/utils"
)

type DocMsgAcknowledgement struct {
	PacketId        string `bson:"packet_id"`
	Packet          Packet `bson:"packet"`
	Acknowledgement string `bson:"acknowledgement"`
	ProofAcked      string `bson:"proof_acked"`
	ProofHeight     Height `bson:"proof_height"`
	Signer          string `bson:"signer"`
}

func (m *DocMsgAcknowledgement) GetType() string {
	return MsgTypeAcknowledgement
}

func (m *DocMsgAcknowledgement) BuildMsg(v interface{}) {

	msg := v.(*MsgAcknowledgement)
	m.Signer = msg.Signer
	m.ProofHeight = loadHeight(msg.ProofHeight)
	m.Acknowledgement = UnmarshalAcknowledgement(msg.Acknowledgement)
	m.ProofAcked = utils.MarshalJsonIgnoreErr(msg.ProofAcked)
	m.Packet = loadPacket(msg.Packet)
	m.PacketId = fmt.Sprintf("%v%v%v%v%v", msg.Packet.SourcePort, msg.Packet.SourceChannel,
		msg.Packet.DestinationPort, msg.Packet.DestinationChannel, msg.Packet.Sequence)

}

func (m *DocMsgAcknowledgement) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
	)

	msg := v.(*MsgAcknowledgement)
	packetData := UnmarshalPacketData(msg.Packet.GetData())
	addrs = append(addrs, msg.Signer, packetData.Receiver, packetData.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
