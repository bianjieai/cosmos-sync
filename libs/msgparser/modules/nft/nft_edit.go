package nft

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"strings"
)

type (
	DocMsgNFTEdit struct {
		Sender  string `bson:"sender"`
		Id      string `bson:"id"`
		Denom   string `bson:"denom"`
		URI     string `bson:"uri"`
		Data    string `bson:"data"`
		Name    string `bson:"name"`
		UriHash string `bson:"uri_hash"`
	}
)

func (m *DocMsgNFTEdit) GetType() string {
	return MsgTypeNFTEdit
}

func (m *DocMsgNFTEdit) BuildMsg(v interface{}) {
	msg := v.(*MsgNFTEdit)

	m.Sender = msg.Sender
	m.Id = strings.ToLower(msg.Id)
	m.Denom = strings.ToLower(msg.DenomId)
	m.URI = msg.URI
	m.Data = msg.Data
	m.Name = msg.Name
	m.UriHash = msg.UriHash
}

func (m *DocMsgNFTEdit) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgNFTEdit)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return modules.CreateMsgDocInfo(v, handler)
}
