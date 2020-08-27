package types

import (
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/irismod/nft"
	"github.com/irismod/record"
	"github.com/irismod/service"
	"github.com/irismod/token"
)

type (
	MsgSend = bank.MsgSend

	MsgNFTMint = nft.MsgMintNFT
	MsgNFTEdit = nft.MsgEditNFT
	MsgNFTTransfer = nft.MsgTransferNFT
	MsgNFTBurn = nft.MsgBurnNFT
	MsgIssueDenom = nft.MsgIssueDenom

	MsgRecordCreate = record.MsgCreateRecord


	MsgDefineService = service.MsgDefineService
	MsgBindService = service.MsgBindService
	MsgCallService = service.MsgCallService
	MsgRespondService = service.MsgRespondService
	MsgUpdateServiceBinding = service.MsgUpdateServiceBinding
	MsgSetWithdrawAddress = service.MsgSetWithdrawAddress
	MsgDisableServiceBinding = service.MsgDisableServiceBinding
	MsgEnableServiceBinding = service.MsgEnableServiceBinding
	MsgRefundServiceDeposit = service.MsgRefundServiceDeposit
	MsgPauseRequestContext = service.MsgPauseRequestContext
	MsgStartRequestContext = service.MsgStartRequestContext
	MsgKillRequestContext = service.MsgKillRequestContext
	MsgUpdateRequestContext = service.MsgUpdateRequestContext
	MsgWithdrawEarnedFees = service.MsgWithdrawEarnedFees

	MsgIssueToken = token.MsgIssueToken
	MsgEditToken = token.MsgEditToken
	MsgMintToken = token.MsgMintToken
	MsgTransferTokenOwner = token.MsgTransferTokenOwner
)
