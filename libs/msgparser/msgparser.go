package msgparser

import (
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/codec"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules/bank"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules/evm"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules/ibc"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules/identity"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules/nft"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules/record"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/modules/service"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/types"
	"strings"
)

type MsgParser interface {
	HandleTxMsg(v types.SdkMsg) MsgDocInfo
}

var (
	RouteClientMap map[string]modules.Client
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
	} else if strings.HasPrefix(data, "/csmod.nft.") {
		route = NftRouteKey
	} else if strings.HasPrefix(data, "/csmod.record.") {
		route = RecordRouteKey
	} else if strings.HasPrefix(data, "/csmod.service.") {
		route = ServiceRouteKey
	} else if strings.HasPrefix(data, "/iritamod.identity.") {
		route = IdentityRouteKey
	} else if strings.HasPrefix(data, "/cschain.ibc.") {
		route = IbcRouteKey
	} else if strings.HasPrefix(data, "/ethermint.evm.") {
		route = EvmRouteKey
	} else {
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
	codec.MakeEncodingConfig()
	RouteClientMap = map[string]modules.Client{
		BankRouteKey:     bank.NewClient(),
		NftRouteKey:      nft.NewClient(),
		RecordRouteKey:   record.NewClient(),
		ServiceRouteKey:  service.NewClient(),
		IdentityRouteKey: identity.NewClient(),
		IbcRouteKey:      ibc.NewClient(),
		EvmRouteKey:      evm.NewClient(),
	}
}
