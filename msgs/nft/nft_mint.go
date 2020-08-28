package nft

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"strings"
)

type DocMsgNFTMint struct {
	Sender    string `bson:"sender"`
	Recipient string `bson:"recipient"`
	Denom     string `bson:"denom"`
	ID        string `bson:"id"`
	URI       string `bson:"uri"`
	Data      string `bson:"data"`
	Name      string `bson:"name"`
}

func (m *DocMsgNFTMint) GetType() string {
	return MsgTypeNFTMint
}

func (m *DocMsgNFTMint) BuildMsg(v interface{}) {
	msg := v.(MsgNFTMint)

	m.Sender = msg.Sender.String()
	m.Recipient = msg.Recipient.String()
	m.ID = strings.ToLower(msg.ID)
	m.Denom = strings.ToLower(msg.Denom)
	m.URI = msg.URI
	m.Data = msg.Data
	m.Name = msg.Name
}

func (m *DocMsgNFTMint) HandleTxMsg(msg MsgNFTMint) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, m.Sender, m.Recipient)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
