package msgparser

import (
	"github.com/bianjieai/cosmos-sync/libs/logger"
	msg_parser "github.com/kaifei-bianjie/msg-parser"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/kaifei-bianjie/msg-parser/types"
	"gopkg.in/mgo.v2/txn"
	"strings"
)

type MsgParser interface {
	HandleTxMsg(v types.SdkMsg) (MsgDocInfo, []txn.Op)
}

var (
	_client msg_parser.MsgClient
)

func NewMsgParser(router Router) MsgParser {
	return &msgParser{
		rh: router,
	}
}

type msgParser struct {
	rh Router
}

// Handler returns the MsgServiceHandler for a given msg or nil if not found.
func (parser msgParser) getModule(v types.SdkMsg) string {
	var (
		route string
	)

	data := types.MsgTypeURL(v)
	if strings.HasPrefix(data, "/ibc.core.") {
		route = IbcRouteKey
	} else if strings.HasPrefix(data, "/ibc.applications.") {
		route = IbcTransferRouteKey
	} else if strings.HasPrefix(data, "/cosmos.bank.") {
		route = BankRouteKey
	} else if strings.HasPrefix(data, "/cosmos.crisis.") {
		route = CrisisRouteKey
	} else if strings.HasPrefix(data, "/cosmos.distribution.") {
		route = DistributionRouteKey
	} else if strings.HasPrefix(data, "/cosmos.slashing.") {
		route = SlashingRouteKey
	} else if strings.HasPrefix(data, "/cosmos.evidence.") {
		route = EvidenceRouteKey
	} else if strings.HasPrefix(data, "/cosmos.staking.") {
		route = StakingRouteKey
	} else if strings.HasPrefix(data, "/cosmos.gov.") {
		route = GovRouteKey
	} else {
		route = data
	}
	return route
}

func (parser msgParser) HandleTxMsg(v types.SdkMsg) (MsgDocInfo, []txn.Op) {
	module := parser.getModule(v)
	handleFunc, err := parser.rh.GetRoute(module)
	if err != nil {
		logger.Warn(err.Error(),
			logger.String("route", module),
			logger.String("type", module))
		return MsgDocInfo{}, nil
	}
	return handleFunc(v), nil
}

func init() {
	_client = msg_parser.NewMsgClient()
}

func handleBank(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Bank.HandleTxMsg(v)
	return docInfo
}
func handleCrisis(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Crisis.HandleTxMsg(v)
	return docInfo
}
func handleDistribution(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Distribution.HandleTxMsg(v)
	return docInfo
}
func handleSlashing(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Slashing.HandleTxMsg(v)
	return docInfo
}
func handleStaking(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Staking.HandleTxMsg(v)
	return docInfo
}
func handleEvidence(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Evidence.HandleTxMsg(v)
	return docInfo
}
func handleGov(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Gov.HandleTxMsg(v)
	return docInfo
}
func handleIbc(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Ibc.HandleTxMsg(v)
	return docInfo
}
