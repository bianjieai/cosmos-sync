package nft

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"strings"
)

type DocMsgIssueDenom struct {
	Id               string `bson:"id"`
	Name             string `bson:"name"`
	Sender           string `bson:"sender"`
	Schema           string `bson:"schema"`
	Symbol           string `bson:"symbol"`
	MintRestricted   bool   `bson:"mint_restricted"`
	UpdateRestricted bool   `bson:"update_restricted"`
	Description      string `bson:"description"`
	Uri              string `bson:"uri"`
	UriHash          string `bson:"uri_hash"`
	Data             string `bson:"data"`
}

func (m *DocMsgIssueDenom) GetType() string {
	return MsgTypeIssueDenom
}

func (m *DocMsgIssueDenom) BuildMsg(v interface{}) {
	msg := v.(*MsgIssueDenom)

	m.Sender = msg.Sender
	m.Schema = msg.Schema
	m.Id = strings.ToLower(msg.Id)
	m.Name = msg.Name
	m.Symbol = msg.Symbol
	m.MintRestricted = msg.MintRestricted
	m.UpdateRestricted = msg.UpdateRestricted
	m.Description = msg.Description
	m.Uri = msg.Uri
	m.UriHash = msg.UriHash
	m.Data = msg.Data
}

func (m *DocMsgIssueDenom) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgIssueDenom)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return modules.CreateMsgDocInfo(v, handler)
}
