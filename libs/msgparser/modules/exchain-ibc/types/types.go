package types

import (
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	ibctransfer "github.com/okex/exchain/libs/ibc-go/modules/apps/transfer/types"
	ibcclient "github.com/okex/exchain/libs/ibc-go/modules/core/02-client/types"
	ibcconnect "github.com/okex/exchain/libs/ibc-go/modules/core/03-connection/types"
	ibc "github.com/okex/exchain/libs/ibc-go/modules/core/04-channel/types"
	ibcchannel "github.com/okex/exchain/libs/ibc-go/modules/core/04-channel/types"
)

type (
	SdkMsg                  sdk.Msg
	FungibleTokenPacketData = ibctransfer.FungibleTokenPacketData
	MsgRecvPacket           = ibc.MsgRecvPacket
	MsgTransfer             = ibctransfer.MsgTransfer

	MsgCreateClient       = ibcclient.MsgCreateClient
	MsgUpdateClient       = ibcclient.MsgUpdateClient
	MsgSubmitMisbehaviour = ibcclient.MsgSubmitMisbehaviour
	MsgUpgradeClient      = ibcclient.MsgUpgradeClient

	MsgConnectionOpenInit    = ibcconnect.MsgConnectionOpenInit
	MsgConnectionOpenAck     = ibcconnect.MsgConnectionOpenAck
	MsgConnectionOpenConfirm = ibcconnect.MsgConnectionOpenConfirm
	MsgConnectionOpenTry     = ibcconnect.MsgConnectionOpenTry

	Acknowledgement        = ibc.Acknowledgement
	MsgChannelOpenInit     = ibcchannel.MsgChannelOpenInit
	MsgChannelOpenTry      = ibcchannel.MsgChannelOpenTry
	MsgChannelOpenAck      = ibcchannel.MsgChannelOpenAck
	MsgChannelOpenConfirm  = ibcchannel.MsgChannelOpenConfirm
	MsgChannelCloseConfirm = ibcchannel.MsgChannelCloseConfirm
	MsgChannelCloseInit    = ibcchannel.MsgChannelCloseInit
	MsgAcknowledgement     = ibcchannel.MsgAcknowledgement
	MsgTimeout             = ibcchannel.MsgTimeout
	MsgTimeoutOnClose      = ibcchannel.MsgTimeoutOnClose
)
