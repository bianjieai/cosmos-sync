package random

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
)

type DocTxMsgRequestRand struct {
	Consumer      string        `bson:"consumer"`       // request address
	BlockInterval uint64        `bson:"block_interval"` // block interval after which the requested random number will be generated
	Oracle        bool          `bson:"oracle"`
	ServiceFeeCap []models.Coin `bson:"service_fee_cap"`
}

func (doctx *DocTxMsgRequestRand) GetType() string {
	return TxTypeRequestRand
}

func (doctx *DocTxMsgRequestRand) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgRequestRandom)
	doctx.Consumer = msg.Consumer
	doctx.BlockInterval = msg.BlockInterval
	doctx.Oracle = msg.Oracle
	doctx.ServiceFeeCap = models.BuildDocCoins(msg.ServiceFeeCap)
}

func (doctx *DocTxMsgRequestRand) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgRequestRandom
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Consumer)
	handler := func() (Msg, []string) {
		return doctx, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
