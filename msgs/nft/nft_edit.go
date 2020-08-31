package nft

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"strings"
)

type (
	DocMsgNFTEdit struct {
		Sender string `bson:"sender"`
		ID     string `bson:"id"`
		Denom  string `bson:"denom"`
		URI    string `bson:"uri"`
		Data   string `bson:"data"`
		Name   string `bson:"name"`
	}
)

func (m *DocMsgNFTEdit) GetType() string {
	return MsgTypeNFTEdit
}

func (m *DocMsgNFTEdit) BuildMsg(v interface{}) {
	msg := v.(MsgNFTEdit)

	m.Sender = msg.Sender.String()
	m.ID = strings.ToLower(msg.ID)
	m.Denom = strings.ToLower(msg.Denom)
	m.URI = msg.URI
	m.Data = msg.Data
	m.Name = msg.Name
}

func (m *DocMsgNFTEdit) HandleTxMsg(msg MsgNFTEdit) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, msg.Sender.String())
	handler := func() (Msg,  []string) {
		return m,  addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
