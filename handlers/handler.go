package handlers

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/msgs/service"
	"github.com/bianjieai/irita-sync/msgs/nft"
	"github.com/bianjieai/irita-sync/msgs/record"
	"github.com/bianjieai/irita-sync/msgs/token"
	"github.com/bianjieai/irita-sync/msgs/bank"
	//"github.com/bianjieai/irita-sync/msgs/distribution"
	//"github.com/bianjieai/irita-sync/msgs/crisis"
	//"github.com/bianjieai/irita-sync/msgs/evidence"
	//"github.com/bianjieai/irita-sync/msgs/staking"
	//"github.com/bianjieai/irita-sync/msgs/gov"
	"github.com/bianjieai/irita-sync/msgs/identity"
)

func HandleTxMsg(v types.Msg) (MsgDocInfo) {
	if BankDocInfo, ok := bank.HandleTxMsg(v); ok {
		return BankDocInfo
	}
	if IServiceDocInfo, ok := service.HandleTxMsg(v); ok {
		return IServiceDocInfo
	}
	if NftDocInfo, ok := nft.HandleTxMsg(v); ok {
		return NftDocInfo
	}
	if RecordDocInfo, ok := record.HandleTxMsg(v); ok {
		return RecordDocInfo
	}
	if TokenDocInfo, ok := token.HandleTxMsg(v); ok {
		return TokenDocInfo
	}
	//if CoinswapDocInfo, ok := coinswap.HandleTxMsg(v); ok {
	//	return CoinswapDocInfo
	//}
	//if CrisisDocInfo, ok := crisis.HandleTxMsg(v); ok {
	//	return CrisisDocInfo
	//}
	//if DistrubutionDocInfo, ok := distribution.HandleTxMsg(v); ok {
	//	return DistrubutionDocInfo
	//}
	//if EvidenceDocInfo, ok := evidence.HandleTxMsg(v); ok {
	//	return EvidenceDocInfo
	//}
	//if HtlcDocInfo, ok := htlc.HandleTxMsg(v); ok {
	//	return HtlcDocInfo
	//}
	//if StakingDocInfo, ok := staking.HandleTxMsg(v); ok {
	//	return StakingDocInfo
	//}
	//if GovDocInfo, ok := gov.HandleTxMsg(v); ok {
	//	return GovDocInfo
	//}
	if IdentityDocInfo, ok := identity.HandleTxMsg(v); ok {
		return IdentityDocInfo
	}
	return MsgDocInfo{}
}

func removeDuplicatesFromSlice(data []string) (result []string) {
	tempSet := make(map[string]string, len(data))
	for _, val := range data {
		if _, ok := tempSet[val]; ok || val == "" {
			continue
		}
		tempSet[val] = val
	}
	for one := range tempSet {
		result = append(result, one)
	}
	return
}
