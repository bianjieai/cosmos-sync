package service

import (
	"encoding/hex"
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
	"strings"
)

type (
	DocMsgUpdateRequestContext struct {
		RequestContextID  string       `bson:"request_context_id" yaml:"request_context_id"`
		Providers         []string     `bson:"providers" yaml:"providers"`
		Consumer          string       `bson:"consumer" yaml:"consumer"`
		ServiceFeeCap     models.Coins `bson:"service_fee_cap" yaml:"service_fee_cap"`
		Timeout           int64        `bson:"timeout" yaml:"timeout"`
		RepeatedFrequency uint64       `bson:"repeated_frequency" yaml:"repeated_frequency"`
		RepeatedTotal     int64        `bson:"repeated_total" yaml:"repeated_total"`
	}
)

func (m *DocMsgUpdateRequestContext) GetType() string {
	return MsgTypeUpdateRequestContext
}

func (m *DocMsgUpdateRequestContext) BuildMsg(v interface{}) {
	msg := v.(MsgUpdateRequestContext)

	loadProviders := func() (ret []string) {
		for _, one := range msg.Providers {
			ret = append(ret, one.String())
		}
		return
	}

	var coins models.Coins
	for _, one := range msg.ServiceFeeCap {
		coins = append(coins, &models.Coin{Denom: one.Denom, Amount: one.Amount.String()})
	}

	m.RequestContextID = strings.ToUpper(hex.EncodeToString(msg.RequestContextID))
	m.Providers = loadProviders()
	m.Consumer = msg.Consumer.String()
	m.ServiceFeeCap = coins
	m.Timeout = msg.Timeout
	m.RepeatedFrequency = msg.RepeatedFrequency
	m.RepeatedTotal = msg.RepeatedTotal
}

func (m *DocMsgUpdateRequestContext) HandleTxMsg(msg MsgUpdateRequestContext) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, m.Providers...)
	addrs = append(addrs, m.Consumer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
