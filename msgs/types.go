package msgs

import (
	"github.com/bianjieai/irita-sync/models"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/irisnet/modules/incubator/nft"
	"github.com/bianjieai/irita/modules/service"
	"github.com/bianjieai/irita/modules/record"
)

const (
	MsgTypeSend            = "send"
	MsgTypeNFTMint         = "nft_mint"
	MsgTypeNFTEdit         = "nft_edit"
	MsgTypeNFTTransfer     = "nft_transfer"
	MsgTypeNFTBurn         = "nft_burn"
	MsgTypeServiceDef      = "service_define"
	MsgTypeServiceBind     = "service_bind"
	MsgTypeServiceRequest  = "service_request"
	MsgTypeServiceResponse = "service_response"
	MsgTypeRecordCreate    = "create_record"
)

type (
	MsgDocInfo struct {
		From       string        // parse from first msg
		To         string        // parse from first msg
		Coins      []models.Coin // parse from first msg
		Signer     string        // parse from first signer
		DocTxMsg   models.DocTxMsg
		ComplexMsg bool
		Signers    []string
	}

	MsgSend = bank.MsgSend

	MsgNFTMint = nft.MsgMintNFT
	MsgNFTEdit = nft.MsgEditNFTMetadata
	MsgNFTTransfer = nft.MsgTransferNFT
	MsgNFTBurn = nft.MsgBurnNFT

	MsgServiceDef = service.MsgSvcDef
	MsgServiceBind = service.MsgSvcBind
	MsgServiceRequest = service.MsgSvcRequest
	MsgServiceResponse = service.MsgSvcResponse

	MsgRecordCreate = record.MsgCreateRecord
)
