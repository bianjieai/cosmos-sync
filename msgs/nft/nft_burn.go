package nft

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"strings"
)

type (
	DocMsgNFTBurn struct {
		Sender string `bson:"sender"`
		ID     string `bson:"id"`
		Denom  string `bson:"denom"`
	}
)

func (m *DocMsgNFTBurn) GetType() string {
	return MsgTypeNFTBurn
}

func (m *DocMsgNFTBurn) BuildMsg(v interface{}) {
	msg := v.(MsgNFTBurn)

	m.Sender = msg.Sender.String()
	m.ID = strings.ToLower(msg.ID)
	m.Denom = strings.ToLower(msg.Denom)
}

func (m *DocMsgNFTBurn) HandleTxMsg(msg MsgNFTBurn) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, msg.Sender.String())
	handler := func() (Msg,  []string) {
		return m,  addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
