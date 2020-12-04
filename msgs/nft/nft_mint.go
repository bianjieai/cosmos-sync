package nft

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
	"strings"
)

type DocMsgNFTMint struct {
	Sender    string `bson:"sender"`
	Recipient string `bson:"recipient"`
	Denom     string `bson:"denom"`
	Id        string `bson:"id"`
	URI       string `bson:"uri"`
	Data      string `bson:"data"`
	Name      string `bson:"name"`
}

func (m *DocMsgNFTMint) GetType() string {
	return MsgTypeNFTMint
}

func (m *DocMsgNFTMint) BuildMsg(v interface{}) {
	msg := v.(*MsgNFTMint)

	m.Sender = msg.Sender
	m.Recipient = msg.Recipient
	m.Id = strings.ToLower(msg.Id)
	m.Denom = strings.ToLower(msg.DenomId)
	m.URI = msg.URI
	m.Data = msg.Data
	m.Name = msg.Name
}

func (m *DocMsgNFTMint) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg MsgNFTMint
	)
	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Sender, msg.Recipient)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
