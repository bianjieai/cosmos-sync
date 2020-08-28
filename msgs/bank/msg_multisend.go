package bank

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

type (
	DocMsgMultiSend struct {
		Inputs   []Item   `bson:"inputs"`
		Outputs  []Item   `bson:"outputs"`
		TempData []string `bson:"-"`
	}
	Item struct {
		Address string        `bson:"address"`
		Coins   []models.Coin `bson:"coins"`
	}
)

func (m *DocMsgMultiSend) GetType() string {
	return MsgTypeMultiSend
}

func (m *DocMsgMultiSend) BuildMsg(v interface{}) {
	msg := v.(MsgMultiSend)
	for _, one := range msg.Inputs {
		m.Inputs = append(m.Inputs, Item{Address: one.Address.String(), Coins: models.BuildDocCoins(one.Coins)})
		m.TempData = append(m.TempData, one.Address.String())
	}
	for _, one := range msg.Outputs {
		m.Outputs = append(m.Outputs, Item{Address: one.Address.String(), Coins: models.BuildDocCoins(one.Coins)})
		m.TempData = append(m.TempData, one.Address.String())
	}

}

func (m *DocMsgMultiSend) HandleTxMsg(msg bank.MsgMultiSend) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, m.TempData...)

	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
