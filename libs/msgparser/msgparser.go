package msgparser

import (
	"github.com/bianjieai/irita-sync/libs/logger"
	msg_parser "github.com/kaifei-bianjie/msg-parser"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/kaifei-bianjie/msg-parser/types"
	"gopkg.in/mgo.v2/txn"
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

func (parser *msgParser) HandleTxMsg(v types.SdkMsg) (MsgDocInfo, []txn.Op) {
	handleFunc, err := parser.rh.GetRoute(v.Route())
	if err != nil {
		logger.Error(err.Error(),
			logger.String("route", v.Route()),
			logger.String("type", v.Type()))
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

func handleNft(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Nft.HandleTxMsg(v)
	return docInfo
}

func handleService(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Service.HandleTxMsg(v)
	return docInfo
}

func handleToken(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Token.HandleTxMsg(v)
	return docInfo
}

func handleOracle(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Oracle.HandleTxMsg(v)
	return docInfo
}

func handleRecord(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Record.HandleTxMsg(v)
	return docInfo
}

func handleRandom(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Random.HandleTxMsg(v)
	return docInfo
}

func handleHtlc(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Htlc.HandleTxMsg(v)
	return docInfo
}

func handleCoinswap(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Coinswap.HandleTxMsg(v)
	return docInfo
}

func handleIdentity(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Identity.HandleTxMsg(v)
	return docInfo
}
