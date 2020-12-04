package oracle

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/models"
)

type DocMsgCreateFeed struct {
	FeedName          string        `bson:"feed_name" yaml:"feed_name"`
	LatestHistory     uint64        `bson:"latest_history" yaml:"latest_history"`
	Description       string        `bson:"description"`
	Creator           string        `bson:"creator"`
	ServiceName       string        `bson:"service_name" yaml:"service_name"`
	Providers         []string      `bson:"providers"`
	Input             string        `bson:"input"`
	Timeout           int64         `bson:"timeout"`
	ServiceFeeCap     []models.Coin `bson:"service_fee_cap" yaml:"service_fee_cap"`
	RepeatedFrequency uint64        `bson:"repeated_frequency" yaml:"repeated_frequency"`
	AggregateFunc     string        `bson:"aggregate_func" yaml:"aggregate_func"`
	ValueJsonPath     string        `bson:"value_json_path" yaml:"value_json_path"`
	ResponseThreshold uint32        `bson:"response_threshold" yaml:"response_threshold"`
}

func (m *DocMsgCreateFeed) GetType() string {
	return TxTypeCreateFeed
}

func (m *DocMsgCreateFeed) BuildMsg(v interface{}) {

	msg := v.(*MsgCreateFeed)

	m.FeedName = msg.FeedName
	m.LatestHistory = msg.LatestHistory
	m.Description = msg.Description
	m.Creator = msg.Creator
	m.ServiceName = msg.ServiceName
	for _, val := range msg.GetProviders() {
		m.Providers = append(m.Providers, val)
	}
	m.Input = msg.Input
	m.Timeout = msg.Timeout
	m.ServiceFeeCap = models.BuildDocCoins(msg.ServiceFeeCap)
	m.RepeatedFrequency = msg.RepeatedFrequency
	m.AggregateFunc = msg.AggregateFunc
	m.ValueJsonPath = msg.ValueJsonPath
	m.ResponseThreshold = msg.ResponseThreshold

}

func (m *DocMsgCreateFeed) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgCreateFeed
	)

	ConvertMsg(v,&msg)
	addrs = append(addrs, msg.Creator)
	for _, val := range msg.GetProviders() {
		addrs = append(addrs, val)
	}
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
