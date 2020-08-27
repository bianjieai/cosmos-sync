package bank

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
)

type (
	DocMsgSend struct {
		FromAddress string        `bson:"from_address"`
		ToAddress   string        `bson:"to_address"`
		Amount      []models.Coin `bson:"amount"`
	}
)

func (m *DocMsgSend) GetType() string {
	return MsgTypeSend
}

func (m *DocMsgSend) BuildMsg(v interface{}) {
	msg := v.(MsgSend)
	m.FromAddress = msg.FromAddress.String()
	m.ToAddress = msg.ToAddress.String()
	m.Amount = models.BuildDocCoins(msg.Amount)
}

func (m *DocMsgSend) HandleTxMsg(msg MsgSend) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, m.FromAddress, m.ToAddress)

	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
