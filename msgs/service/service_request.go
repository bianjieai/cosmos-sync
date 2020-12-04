package service

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
)

type (
	DocMsgCallService struct {
		ServiceName       string       `bson:"service_name"`
		Providers         []string     `bson:"providers"`
		Consumer          string       `bson:"consumer"`
		Input             string       `bson:"input"`
		ServiceFeeCap     models.Coins `bson:"service_fee_cap"`
		Timeout           int64        `bson:"timeout"`
		SuperMode         bool         `bson:"super_mode"`
		Repeated          bool         `bson:"repeated"`
		RepeatedFrequency uint64       `bson:"repeated_frequency"`
		RepeatedTotal     int64        `bson:"repeated_total"`
	}
)

func (m *DocMsgCallService) GetType() string {
	return MsgTypeCallService
}

func (m *DocMsgCallService) BuildMsg(msg interface{}) {
	v := msg.(*MsgCallService)

	var coins models.Coins
	for _, one := range v.ServiceFeeCap {
		coins = append(coins, &models.Coin{Denom: one.Denom, Amount: one.Amount.String()})
	}
	m.ServiceName = v.ServiceName
	m.Providers = m.loadProviders(v)
	m.Consumer = v.Consumer.String()
	m.Input = v.Input
	m.ServiceFeeCap = coins
	m.Timeout = v.Timeout
	//m.Input = hex.EncodeToString(v.Input)
	m.SuperMode = v.SuperMode
	m.Repeated = v.Repeated
	m.RepeatedFrequency = v.RepeatedFrequency
	m.RepeatedTotal = v.RepeatedTotal
}

func (m *DocMsgCallService)loadProviders(v *MsgCallService) (ret []string) {
	for _, one := range v.Providers {
		ret = append(ret, one.String())
	}
	return
}

func (m *DocMsgCallService) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg MsgCallService
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)

	addrs = append(addrs, m.loadProviders(&msg)...)
	addrs = append(addrs, msg.Consumer.String())
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
