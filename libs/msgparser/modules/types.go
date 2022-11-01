package msgs

import (
	models "github.com/bianjieai/cosmos-sync/libs/msgparser/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctransfer "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	ibcclient "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	ibcconnect "github.com/cosmos/ibc-go/v5/modules/core/03-connection/types"
	ibc "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	ibcchannel "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
)

const (

	//ibc client
	MsgTypeCreateClient             = "create_client"
	MsgTypeUpdateClient             = "update_client"
	MsgTypeUpgradeClient            = "upgrade_client"
	MsgTypeSubmitMisbehaviourClient = "submit_misbehaviour"

	//ibc connect
	MsgTypeConnectionOpenInit    = "connection_open_init"
	MsgTypeConnectionOpenTry     = "connection_open_try"
	MsgTypeConnectionOpenAck     = "connection_open_ack"
	MsgTypeConnectionOpenConfirm = "connection_open_confirm"

	//ibc channel
	MsgTypeChannelOpenInit     = "channel_open_init"
	MsgTypeChannelOpenTry      = "channel_open_try"
	MsgTypeChannelOpenAck      = "channel_open_ack"
	MsgTypeChannelOpenConfirm  = "channel_open_confirm"
	MsgTypeChannelCloseInit    = "channel_close_init"
	MsgTypeChannelCloseConfirm = "channel_close_confirm"
	MsgTypeTimeout             = "timeout_packet"
	MsgTypeTimeoutOnClose      = "timeout_on_close_packet"
	MsgTypeAcknowledgement     = "acknowledge_packet"

	MsgTypeRecvPacket  = "recv_packet"
	MsgTypeIBCTransfer = "transfer"
)

type (
	MsgDocInfo struct {
		DocTxMsg models.TxMsg
		Addrs    []string
		Signers  []string
	}
	SdkMsg sdk.Msg
	Msg    models.Msg

	Coin models.Coin

	Coins []*Coin

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
