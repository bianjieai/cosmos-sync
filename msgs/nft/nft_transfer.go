package nft

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"strings"
)

type (
	DocMsgNFTTransfer struct {
		Sender    string `bson:"sender"`
		Recipient string `bson:"recipient"`
		URI       string `bson:"uri"`
		Name      string `bson:"name"`
		Denom     string `bson:"denom"`
		ID        string `bson:"id"`
		Data      string `bson:"data"`
	}
)

func (m *DocMsgNFTTransfer) GetType() string {
	return MsgTypeNFTTransfer
}

func (m *DocMsgNFTTransfer) BuildMsg(v interface{}) {
	msg := v.(*MsgNFTTransfer)

	m.Sender = msg.Sender.String()
	m.Recipient = msg.Recipient.String()
	m.ID = strings.ToLower(msg.ID)
	m.Denom = strings.ToLower(msg.Denom)
	m.URI = msg.URI
	m.Data = msg.Data
	m.Name = msg.Name
}

func (m *DocMsgNFTTransfer) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgNFTTransfer
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Sender.String(), msg.Recipient.String())
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
