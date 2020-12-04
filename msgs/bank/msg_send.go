package bank

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
)

type (
	DocMsgSend struct {
		FromAddress string        `bson:"fromaddress"`
		ToAddress   string        `bson:"toaddress"`
		Amount      []models.Coin `bson:"amount"`
	}
)

func (m *DocMsgSend) GetType() string {
	return MsgTypeSend
}

func (m *DocMsgSend) BuildMsg(v interface{}) {
	msg := v.(*MsgSend)
	m.FromAddress = msg.FromAddress
	m.ToAddress = msg.ToAddress
	m.Amount = models.BuildDocCoins(msg.Amount)
}

func (m *DocMsgSend) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg MsgSend
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)

	addrs = append(addrs, msg.FromAddress, msg.ToAddress)

	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
