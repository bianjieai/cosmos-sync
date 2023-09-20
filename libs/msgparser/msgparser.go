package msgparser

import (
	"github.com/bianjieai/cosmos-sync/libs/logger"
	common_parser "github.com/kaifei-bianjie/common-parser"
	. "github.com/kaifei-bianjie/common-parser/modules"
	"github.com/kaifei-bianjie/common-parser/types"
	cosmosmod_parser "github.com/kaifei-bianjie/cosmosmod-parser"
	cschain_mod_parser "github.com/kaifei-bianjie/cschain-mod-parser"
	irismod_parser "github.com/kaifei-bianjie/irismod-parser"
	iritamod_parser "github.com/kaifei-bianjie/iritamod-parser"
	"strings"
)

type MsgParser interface {
	HandleTxMsg(v types.SdkMsg) MsgDocInfo
}

var (
	cosmosModClient  cosmosmod_parser.MsgClient
	chainModClient   irismod_parser.MsgClient
	iritaModClient   iritamod_parser.MsgClient
	cschainModClient cschain_mod_parser.MsgClient

	RouteClientMap map[string]common_parser.Client
	RoutePrefixMap = map[string]string{
		"/cosmos.bank.":       BankRouteKey,
		"/irismod.nft.":       NftRouteKey,
		"/irismod.record.":    RecordRouteKey,
		"/irismod.service.":   ServiceRouteKey,
		"/iritamod.identity.": IdentityRouteKey,
		"/cschain.ibc.":       IbcRouteKey,
	}
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
	for prefix, r := range RoutePrefixMap {
		if strings.HasPrefix(data, prefix) {
			route = r
		}
	}
	if _, ok := RouteClientMap[route]; !ok {
		route = data
	}
	return route
}

func (parser *msgParser) HandleTxMsg(v types.SdkMsg) MsgDocInfo {
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
	cosmosModClient = cosmosmod_parser.NewMsgClient()
	chainModClient = irismod_parser.NewMsgClient()
	iritaModClient = iritamod_parser.NewMsgClient()
	cschainModClient = cschain_mod_parser.NewMsgClient()

	RouteClientMap = map[string]common_parser.Client{
		BankRouteKey:     cosmosModClient.Bank,
		NftRouteKey:      chainModClient.Nft,
		RecordRouteKey:   chainModClient.Record,
		ServiceRouteKey:  chainModClient.Service,
		IdentityRouteKey: iritaModClient.Identity,
		IbcRouteKey:      cschainModClient.Ibc,
	}
}
