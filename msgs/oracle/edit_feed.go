package oracle

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
)

type DocMsgEditFeed struct {
	FeedName          string        `bson:"feed_name" yaml:"feed_name"`
	LatestHistory     uint64        `bson:"latest_history" yaml:"latest_history"`
	Description       string        `bson:"description"`
	Creator           string        `bson:"creator"`
	Providers         []string      `bson:"providers"`
	Timeout           int64         `bson:"timeout"`
	ServiceFeeCap     []models.Coin `bson:"service_fee_cap" yaml:"service_fee_cap"`
	RepeatedFrequency uint64        `bson:"repeated_frequency" yaml:"repeated_frequency"`
	ResponseThreshold uint32        `bson:"response_threshold" yaml:"response_threshold"`
}

func (m *DocMsgEditFeed) GetType() string {
	return TxTypeEditFeed
}

func (m *DocMsgEditFeed) BuildMsg(v interface{}) {
	msg := v.(*MsgEditFeed)

	m.FeedName = msg.FeedName
	m.LatestHistory = msg.LatestHistory
	m.Description = msg.Description
	m.Creator = msg.Creator
	for _, val := range msg.GetProviders() {
		m.Providers = append(m.Providers, val)
	}
	m.Timeout = msg.Timeout
	m.ServiceFeeCap = models.BuildDocCoins(msg.ServiceFeeCap)
	m.RepeatedFrequency = msg.RepeatedFrequency
	m.ResponseThreshold = msg.ResponseThreshold
}

func (m *DocMsgEditFeed) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgEditFeed
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Creator)
	for _, val := range msg.GetProviders() {
		addrs = append(addrs, val)
	}
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
