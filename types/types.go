package types

import (
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/irisnet/modules/incubator/nft"
	"github.com/bianjieai/irita/modules/service"
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
)
