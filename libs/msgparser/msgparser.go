package msgparser

import (
	"github.com/bianjieai/cosmos-sync/libs/logger"
	common_parser "github.com/kaifei-bianjie/common-parser"
	. "github.com/kaifei-bianjie/common-parser/modules"
	"github.com/kaifei-bianjie/common-parser/types"
	cosmosmod_parser "github.com/kaifei-bianjie/cosmosmod-parser"
	irismod_parser "github.com/kaifei-bianjie/irismod-parser"
	iritachain_mod_parser "github.com/kaifei-bianjie/iritachain-mod-parser"
	iritamod_parser "github.com/kaifei-bianjie/iritamod-parser"
	spartanchain_mod_parser "github.com/kaifei-bianjie/spartanchain-mod-parser"
	tibc_mod_parser "github.com/kaifei-bianjie/tibc-mod-parser"
	"strings"
)

type MsgParser interface {
	HandleTxMsg(v types.SdkMsg) MsgDocInfo
}

var (
	irisModClient         irismod_parser.MsgClient
	cosmosModClient       cosmosmod_parser.MsgClient
	iritaChainModClient   iritachain_mod_parser.MsgClient
	iritaModClient        iritamod_parser.MsgClient
	tibcModClient         tibc_mod_parser.MsgClient
	spartanChainModClient spartanchain_mod_parser.MsgClient

	RouteClientMap map[string]common_parser.Client
)

func NewMsgParser() MsgParser {
	return &msgParser{}
}

type msgParser struct{}

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
		route = CosmosSlashingRouteKey
	} else if strings.HasPrefix(data, "/cosmos.evidence.") {
		route = EvidenceRouteKey
	} else if strings.HasPrefix(data, "/cosmos.staking.") {
		route = StakingRouteKey
	} else if strings.HasPrefix(data, "/cosmos.gov.") {
		route = GovRouteKey
	} else if strings.HasPrefix(data, "/cosmos.feegrant") {
		route = FeegrantRouteKey
	} else if strings.HasPrefix(data, "/tibc.core.") {
		route = TIbcRouteKey
	} else if strings.HasPrefix(data, "/tibc.apps.") {
		route = TIbcTransferRouteKey
	} else if strings.HasPrefix(data, "/irismod.nft.") {
		route = NftRouteKey
	} else if strings.HasPrefix(data, "/irismod.mt.") {
		route = MtRouteKey
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
	} else if strings.HasPrefix(data, "/iritamod.slashing") {
		route = IritaSlashingRouteKey
	} else if strings.HasPrefix(data, "/iritamod.perm") {
		route = PermRouteKey
	} else {
		route = data
	}
	return route
}

func (parser msgParser) HandleTxMsg(v types.SdkMsg) MsgDocInfo {
	module := parser.getModule(v)
	client := RouteClientMap[module]
	if client == nil {
		logger.Warn("no support msg parse",
			logger.String("route", module),
			logger.String("type", module))
		return MsgDocInfo{}
	}
	msgDocInfo, b := client.HandleTxMsg(v)
	if !b {
		logger.Warn("HandleTxMsg error",
			logger.String("route", module),
			logger.String("type", module))
		return MsgDocInfo{}
	}
	return msgDocInfo
}

func init() {
	irisModClient = irismod_parser.NewMsgClient()
	cosmosModClient = cosmosmod_parser.NewMsgClient()
	iritaChainModClient = iritachain_mod_parser.NewMsgClient()
	iritaModClient = iritamod_parser.NewMsgClient()
	tibcModClient = tibc_mod_parser.NewMsgClient()
	spartanChainModClient = spartanchain_mod_parser.NewMsgClient()
	MakeEncodingConfig()
	RouteClientMap = map[string]common_parser.Client{
		NftRouteKey:            irisModClient.Nft,
		MtRouteKey:             irisModClient.Mt,
		ServiceRouteKey:        irisModClient.Service,
		TokenRouteKey:          irisModClient.Token,
		HtlcRouteKey:           irisModClient.Htlc,
		CoinswapRouteKey:       irisModClient.Coinswap,
		RandomRouteKey:         irisModClient.Random,
		OracleRouteKey:         irisModClient.Oracle,
		RecordRouteKey:         irisModClient.Record,
		BankRouteKey:           cosmosModClient.Bank,
		StakingRouteKey:        cosmosModClient.Staking,
		DistributionRouteKey:   cosmosModClient.Distribution,
		CrisisRouteKey:         cosmosModClient.Crisis,
		EvidenceRouteKey:       cosmosModClient.Evidence,
		FeegrantRouteKey:       cosmosModClient.Feegrant,
		CosmosSlashingRouteKey: cosmosModClient.Slashing,
		IritaSlashingRouteKey:  iritaModClient.Slashing,
		IdentityRouteKey:       iritaModClient.Identity,
		PermRouteKey:           iritaModClient.Perm,
		EvmRouteKey:            iritaChainModClient.Evm,
		TIbcTransferRouteKey:   tibcModClient.Tibc,
		TIbcRouteKey:           tibcModClient.Tibc,
		GovRouteKey:            spartanChainModClient.Gov,
	}
}
