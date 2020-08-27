package service

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
)

type (
	DocMsgUpdateServiceBinding struct {
		ServiceName string       `bson:"service_name" yaml:"service_name"`
		Provider    string       `bson:"provider" yaml:"provider"`
		Deposit     models.Coins `bson:"deposit" yaml:"deposit"`
		Pricing     string       `bson:"pricing" yaml:"pricing"`
		QoS         uint64       `bson:"qos" yaml:"qos"`
		Owner       string       `bson:"owner" yaml:"owner"`
	}
)

func (m *DocMsgUpdateServiceBinding) GetType() string {
	return MsgTypeUpdateServiceBinding
}

func (m *DocMsgUpdateServiceBinding) BuildMsg(v interface{}) {
	msg := v.(MsgUpdateServiceBinding)

	var coins models.Coins
	for _, one := range msg.Deposit {
		coins = append(coins, &models.Coin{Denom: one.Denom, Amount: one.Amount.String()})
	}

	m.ServiceName = msg.ServiceName
	m.Provider = msg.Provider.String()
	m.Deposit = coins
	m.Pricing = msg.Pricing
	m.QoS = msg.QoS
	m.Owner = msg.Owner.String()
}

func (m *DocMsgUpdateServiceBinding) HandleTxMsg(msg MsgUpdateServiceBinding) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, m.Owner, m.Provider)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
