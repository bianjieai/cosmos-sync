package service

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
)

type (
	DocMsgEnableServiceBinding struct {
		ServiceName string       `bson:"service_name" yaml:"service_name"`
		Provider    string       `bson:"provider" yaml:"provider"`
		Deposit     models.Coins `bson:"deposit" yaml:"deposit"`
		Owner       string       `bson:"owner" yaml:"owner"`
	}
)

func (m *DocMsgEnableServiceBinding) GetType() string {
	return MsgTypeEnableServiceBinding
}

func (m *DocMsgEnableServiceBinding) BuildMsg(v interface{}) {
	msg := v.(*MsgEnableServiceBinding)

	var coins models.Coins
	for _, one := range msg.Deposit {
		coins = append(coins, &models.Coin{Denom: one.Denom, Amount: one.Amount.String()})
	}

	m.ServiceName = msg.ServiceName
	m.Provider = msg.Provider
	m.Deposit = coins
	m.Owner = msg.Owner
}

func (m *DocMsgEnableServiceBinding) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg MsgEnableServiceBinding
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Owner, msg.Provider)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
