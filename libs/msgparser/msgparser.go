package msgparser

import (
	"github.com/bianjieai/irita-sync/libs/logger"
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
func (parser msgParser) getModule(v types.SdkMsg) (string, string) {
	var (
		route   string
		msgType string
	)
	if legacyMsg, ok := v.(types.LegacyMsg); ok {
		route = legacyMsg.Route()
		msgType = legacyMsg.Type()
	} else {
		data := types.MsgTypeURL(v)
		if strings.HasPrefix(data, "/ibc.core.") {
			route = IbcRouteKey
		} else if strings.HasPrefix(data, "/ibc.applications.") {
			route = IbcTransferRouteKey
		} else if strings.HasPrefix(data, "/tibc.core.") {
			route = TIbcRouteKey
		} else if strings.HasPrefix(data, "/tibc.apps.") {
			route = TIbcTransferRouteKey
		} else {
			route = data
		}
	}
	return route, msgType
}

func (parser msgParser) HandleTxMsg(v types.SdkMsg) (MsgDocInfo, []txn.Op) {
	module, msgType := parser.getModule(v)
	handleFunc, err := parser.rh.GetRoute(module)
	if err != nil {
		logger.Error(err.Error(),
			logger.String("route", module),
			logger.String("type", msgType))
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
func handleService(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Service.HandleTxMsg(v)
	return docInfo
}
func handleNft(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Nft.HandleTxMsg(v)
	return docInfo
}
func handleRecord(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Record.HandleTxMsg(v)
	return docInfo
}
func handleToken(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Token.HandleTxMsg(v)
	return docInfo
}
func handleCoinswap(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Coinswap.HandleTxMsg(v)
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
func handleEvidence(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Evidence.HandleTxMsg(v)
	return docInfo
}
func handleHtlc(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Htlc.HandleTxMsg(v)
	return docInfo
}
func handleStaking(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Staking.HandleTxMsg(v)
	return docInfo
}
func handleGov(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Gov.HandleTxMsg(v)
	return docInfo
}
func handleRandom(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Random.HandleTxMsg(v)
	return docInfo
}
func handleOracle(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Oracle.HandleTxMsg(v)
	return docInfo
}
func handleIbc(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Ibc.HandleTxMsg(v)
	return docInfo
}

func handleTIbc(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Tibc.HandleTxMsg(v)
	return docInfo
}

func handleFarm(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Farm.HandleTxMsg(v)
	return docInfo
}
