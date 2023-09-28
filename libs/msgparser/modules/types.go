package modules

import (
	models "github.com/bianjieai/cosmos-sync/libs/msgparser/types"
	"github.com/bianjieai/iritamod/modules/identity"
	"github.com/bianjieai/iritamod/modules/node"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	evm "github.com/tharsis/ethermint/x/evm/types"
	cschainibc "gitlab.cschain.tech/cschain/modules/ibc/core/types"
	nft "gitlab.cschain.tech/csmod/modules/nft/types"
	record "gitlab.cschain.tech/csmod/modules/record/types"
	service "gitlab.cschain.tech/csmod/modules/service/types"
)

const (
	MsgTypeSend      = "send"
	MsgTypeMultiSend = "multisend"
)

type (
	MsgSend      = bank.MsgSend
	MsgMultiSend = bank.MsgMultiSend
)

const (

	//nft
	MsgTypeIssueDenom    = "issue_denom"
	MsgTypeTransferDenom = "transfer_denom"
	MsgTypeNFTMint       = "mint_nft"
	MsgTypeNFTTransfer   = "transfer_nft"
	MsgTypeNFTEdit       = "edit_nft"
	MsgTypeNFTBurn       = "burn_nft"

	//record
	MsgTypeRecordCreate = "create_record"
	//service
	MsgTypeDefineService             = "define_service"
	MsgTypeBindService               = "bind_service"
	MsgTypeUpdateServiceBinding      = "update_service_binding"
	MsgTypeServiceSetWithdrawAddress = "service/set_withdraw_address"
	MsgTypeDisableServiceBinding     = "disable_service_binding"
	MsgTypeEnableServiceBinding      = "enable_service_binding"
	MsgTypeRefundServiceDeposit      = "refund_service_deposit"
	MsgTypeCallService               = "call_service"
	MsgTypeRespondService            = "respond_service"
	MsgTypePauseRequestContext       = "pause_request_context"
	MsgTypeStartRequestContext       = "start_request_context"
	MsgTypeKillRequestContext        = "kill_request_context"
	MsgTypeUpdateRequestContext      = "update_request_context"
	MsgTypeWithdrawEarnedFees        = "withdraw_earned_fees"

	//identity
	MsgTypeUpdateIdentity = "update_identity"
	MsgTypeCreateIdentity = "create_identity"

	//node
	MsgTypeCreateValidator = "create_validator" // type for MsgCreateValidator
	MsgTypeUpdateValidator = "update_validator" // type for MsgUpdateValidator
	MsgTypeRemoveValidator = "remove_validator" // type for MsgRemoveValidator
	MsgTypeGrantNode       = "grant_node"       // type for MsgGrantNode
	MsgTypeRevokeNode      = "revoke_node"      // type for MsgRevokeNode
)

type (

	//nft
	MsgNFTMint       = nft.MsgMintNFT
	MsgNFTEdit       = nft.MsgEditNFT
	MsgNFTTransfer   = nft.MsgTransferNFT
	MsgNFTBurn       = nft.MsgBurnNFT
	MsgIssueDenom    = nft.MsgIssueDenom
	MsgTransferDenom = nft.MsgTransferDenom

	//record
	MsgRecordCreate = record.MsgCreateRecord

	//service
	MsgDefineService         = service.MsgDefineService
	MsgBindService           = service.MsgBindService
	MsgCallService           = service.MsgCallService
	MsgRespondService        = service.MsgRespondService
	MsgUpdateServiceBinding  = service.MsgUpdateServiceBinding
	MsgSetWithdrawAddress    = service.MsgSetWithdrawAddress
	MsgDisableServiceBinding = service.MsgDisableServiceBinding
	MsgEnableServiceBinding  = service.MsgEnableServiceBinding
	MsgRefundServiceDeposit  = service.MsgRefundServiceDeposit
	MsgPauseRequestContext   = service.MsgPauseRequestContext
	MsgStartRequestContext   = service.MsgStartRequestContext
	MsgKillRequestContext    = service.MsgKillRequestContext
	MsgUpdateRequestContext  = service.MsgUpdateRequestContext
	MsgWithdrawEarnedFees    = service.MsgWithdrawEarnedFees

	MsgCreateIdentity = identity.MsgCreateIdentity
	MsgUpdateIdentity = identity.MsgUpdateIdentity

	MsgNodeCreate = node.MsgCreateValidator
	MsgNodeUpdate = node.MsgUpdateValidator
	MsgNodeRemove = node.MsgRemoveValidator
	MsgNodeGrant  = node.MsgGrantNode
	MsgNodeRevoke = node.MsgRevokeNode
)

const (
	//ibc client
	MsgTypeCreateClient = "create_client"
	MsgTypeUpdateClient = "update_client"
	MsgTypeRecvPacket   = "recv_packet"
)

type (
	//ibc
	MsgRecvPacket   = cschainibc.MsgRecvPacket
	MsgCreateClient = cschainibc.MsgCreateClient
	MsgUpdateClient = cschainibc.MsgUpdateClient
)

const (
	MsgTypeEthereumTx = "ethereum_tx"
)

type (
	//evm
	MsgEthereumTx = evm.MsgEthereumTx
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
)
