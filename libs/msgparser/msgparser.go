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
	if strings.HasPrefix(data, "/cosmos.bank.") {
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
	} else if strings.HasPrefix(data, "/tibc.core.") {
		route = TIbcRouteKey
	} else if strings.HasPrefix(data, "/tibc.apps.") {
		route = TIbcTransferRouteKey
	} else if strings.HasPrefix(data, "/irismod.nft.") {
		route = NftRouteKey
	} else if strings.HasPrefix(data, "/irismod.coinswap.") {
		route = CoinswapRouteKey
	} else if strings.HasPrefix(data, "/irismod.token.") {
		route = TokenRouteKey
	} else if strings.HasPrefix(data, "/irismod.record.") {
		route = RecordRouteKey
	} else if strings.HasPrefix(data, "/irismod.service.") {
		route = ServiceRouteKey
	} else if strings.HasPrefix(data, "/irismod.htlc.") {
		route = HtlcRouteKey
	} else if strings.HasPrefix(data, "/irismod.random.") {
		route = RandomRouteKey
	} else if strings.HasPrefix(data, "/irismod.oracle.") {
		route = OracleRouteKey
	} else if strings.HasPrefix(data, "/iritamod.identity.") {
		route = IdentityRouteKey
	} else if strings.HasPrefix(data, "/ethermint.evm.") {
		route = EvmRouteKey
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

func handleTIbc(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Tibc.HandleTxMsg(v)
	return docInfo
}

func handleIdentity(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Identity.HandleTxMsg(v)
	return docInfo
}

func handleEvm(v types.SdkMsg) MsgDocInfo {
	docInfo, _ := _client.Evm.HandleTxMsg(v)
	return docInfo
}
