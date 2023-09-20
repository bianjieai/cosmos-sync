package service

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	models "github.com/bianjieai/cosmos-sync/libs/msgparser/types"
	"strings"
)

type (
	DocMsgUpdateRequestContext struct {
		RequestContextID  string        `bson:"request_context_id" yaml:"request_context_id"`
		Providers         []string      `bson:"providers" yaml:"providers"`
		Consumer          string        `bson:"consumer" yaml:"consumer"`
		ServiceFeeCap     []models.Coin `bson:"service_fee_cap" yaml:"service_fee_cap"`
		Timeout           int64         `bson:"timeout" yaml:"timeout"`
		RepeatedFrequency int64         `bson:"repeated_frequency" yaml:"repeated_frequency"`
		RepeatedTotal     int64         `bson:"repeated_total" yaml:"repeated_total"`
	}
)

func (m *DocMsgUpdateRequestContext) GetType() string {
	return MsgTypeUpdateRequestContext
}

func (m *DocMsgUpdateRequestContext) BuildMsg(v interface{}) {
	msg := v.(*MsgUpdateRequestContext)

	var coins []models.Coin
	for _, one := range msg.ServiceFeeCap {
		coins = append(coins, models.Coin{Denom: one.Denom, Amount: one.Amount.String()})
	}

	m.RequestContextID = strings.ToUpper(msg.RequestContextId)
	m.Providers = msg.Providers
	m.Consumer = msg.Consumer
	m.ServiceFeeCap = coins
	m.Timeout = msg.Timeout
	m.RepeatedFrequency = int64(msg.RepeatedFrequency)
	m.RepeatedTotal = msg.RepeatedTotal
}

func (m *DocMsgUpdateRequestContext) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgUpdateRequestContext)
	addrs = append(addrs, msg.Providers...)
	addrs = append(addrs, msg.Consumer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return modules.CreateMsgDocInfo(v, handler)
}
