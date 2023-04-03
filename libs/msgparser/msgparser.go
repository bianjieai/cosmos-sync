package msgparser

import (
	"github.com/bianjieai/cosmos-sync/libs/logger"
	_ "github.com/bianjieai/cosmos-sync/libs/logger"
	commonparser "github.com/kaifei-bianjie/common-parser"
	. "github.com/kaifei-bianjie/common-parser/modules"
	"github.com/kaifei-bianjie/common-parser/types"
	cosmosmod "github.com/kaifei-bianjie/cosmosmod-parser"
	irismod "github.com/kaifei-bianjie/irismod-parser"
	tibcmod "github.com/kaifei-bianjie/tibc-mod-parser"
	"strings"
)

type MsgParser interface {
	HandleTxMsg(v types.SdkMsg) MsgDocInfo
	GetModule(data string) string
}

var (
	cosmosModClient cosmosmod.MsgClient
	irisModClient   irismod.MsgClient
	tibcModClient   tibcmod.MsgClient
	RouteClientMap  map[string]commonparser.Client
)

func NewMsgParser() MsgParser {
	return &msgParser{}
}

type msgParser struct {
}

// Handler returns the MsgServiceHandler for a given msg or nil if not found.
func (parser msgParser) GetModule(data string) string {
	var (
		route string
	)

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
	} else if strings.HasPrefix(data, "/cosmos.feegrant") {
		route = FeegrantRouteKey
	} else if strings.HasPrefix(data, "/cosmos.authz.") {
		route = AuthzRouteKey
	} else if strings.HasPrefix(data, "/cosmos.group.") {
		route = GroupRouteKey
	} else if strings.HasPrefix(data, "/tibc.core.") {
		route = TIbcRouteKey
	} else if strings.HasPrefix(data, "/tibc.apps.") {
		route = TIbcTransferRouteKey
	} else if strings.HasPrefix(data, "/irismod.nft.") {
		route = NftRouteKey
	} else if strings.HasPrefix(data, "/irismod.mt.") {
		route = MtRouteKey
	} else if strings.HasPrefix(data, "/irismod.farm.") {
		route = FarmRouteKey
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
	} else if strings.HasPrefix(data, "/ethermint.evm.") {
		route = EvmRouteKey
	} else {
		route = data
	}
	return route
}

func (parser msgParser) HandleTxMsg(v types.SdkMsg) MsgDocInfo {
	msgName := types.MsgTypeURL(v)
	module := parser.GetModule(msgName)
	client, ok := RouteClientMap[module]
	if !ok {
		logger.Warn("not support module",
			logger.String("route", module),
			logger.String("type", module))
		return MsgDocInfo{}
	}
	msg, _ := client.HandleTxMsg(v)
	return msg

}

func init() {
	cosmosModClient = cosmosmod.NewMsgClient()
	irisModClient = irismod.NewMsgClient()
	tibcModClient = tibcmod.NewMsgClient()
	AdaptEthermintEncodingConfig()
	RouteClientMap = map[string]commonparser.Client{
		BankRouteKey:         cosmosModClient.Bank,
		ServiceRouteKey:      irisModClient.Service,
		NftRouteKey:          irisModClient.Nft,
		MtRouteKey:           irisModClient.Mt,
		RecordRouteKey:       irisModClient.Record,
		TokenRouteKey:        irisModClient.Token,
		CoinswapRouteKey:     irisModClient.Coinswap,
		CrisisRouteKey:       cosmosModClient.Crisis,
		DistributionRouteKey: cosmosModClient.Distribution,
		SlashingRouteKey:     cosmosModClient.Slashing,
		EvidenceRouteKey:     cosmosModClient.Evidence,
		HtlcRouteKey:         irisModClient.Htlc,
		StakingRouteKey:      cosmosModClient.Staking,
		GovRouteKey:          cosmosModClient.Gov,
		FeegrantRouteKey:     cosmosModClient.Feegrant,
		RandomRouteKey:       irisModClient.Random,
		OracleRouteKey:       irisModClient.Oracle,
		IbcRouteKey:          cosmosModClient.Ibc,
		IbcTransferRouteKey:  cosmosModClient.Ibc,
		FarmRouteKey:         irisModClient.Farm,
		TIbcTransferRouteKey: tibcModClient.Tibc,
		TIbcRouteKey:         tibcModClient.Tibc,
		AuthzRouteKey:        cosmosModClient.Authz,
		GroupRouteKey:        cosmosModClient.Group,
		EvmRouteKey:          irisModClient.Evm,
	}
}
