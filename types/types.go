package types

import (
	"github.com/bianjieai/irita/modules/record"
	"github.com/bianjieai/irita/modules/service"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/irisnet/modules/incubator/nft"
)

type (
	MsgSend = bank.MsgSend

	MsgNFTMint     = nft.MsgMintNFT
	MsgNFTEdit     = nft.MsgEditNFTMetadata
	MsgNFTTransfer = nft.MsgTransferNFT
	MsgNFTBurn     = nft.MsgBurnNFT

	MsgServiceDef      = service.MsgSvcDef
	MsgServiceBind     = service.MsgSvcBind
	MsgServiceRequest  = service.MsgSvcRequest
	MsgServiceResponse = service.MsgSvcResponse

	MsgRecordCreate = record.MsgCreateRecord
)
