package msgparser

import (
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/client"
	. "github.com/bianjieai/cosmos-sync/libs/msgparser/modules/ibc/types"
	"github.com/okex/exchain/libs/cosmos-sdk/types"
	"strings"
)

type MsgParser interface {
	HandleTxMsg(v types.Msg) MsgDocInfo
	MsgType(v types.Msg) string
}

var (
	_client client.MsgClient
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
func (parser msgParser) getModule(v types.Msg) string {
	var (
		route string
	)
	data := types.MsgTypeURL(v.(types.MsgProtoAdapter))
	if strings.HasPrefix(data, "/ibc.core.") {
		route = IbcRouteKey
	} else if strings.HasPrefix(data, "/ibc.applications.") {
		route = IbcTransferRouteKey
	} else {
		route = data
	}
	return route
}

func (parser msgParser) HandleTxMsg(v types.Msg) MsgDocInfo {
	module := parser.getModule(v)
	handleFunc, err := parser.rh.GetRoute(module)
	if err != nil {
		logger.Warn(err.Error(),
			logger.String("route", module),
			logger.String("type", module))
		return MsgDocInfo{}
	}
	return handleFunc(v)
}

func (parser msgParser) MsgType(v types.Msg) string {
	switch v.(type) {
	case *MsgRecvPacket:
		return MsgTypeRecvPacket
	case *MsgTransfer:
		return MsgTypeIBCTransfer
	case *MsgUpdateClient:
		return MsgTypeUpdateClient
	case *MsgChannelOpenConfirm:
		return MsgTypeChannelOpenConfirm
	case *MsgTimeout:
		return MsgTypeTimeout
	case *MsgAcknowledgement:
		return MsgTypeAcknowledgement
	}
	return ""
}

func init() {
	_client = client.NewMsgClient()
}
func handleIbc(v types.Msg) MsgDocInfo {
	docInfo, _ := _client.Ibc.HandleTxMsg(v)
	return docInfo
}
