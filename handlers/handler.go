package handlers

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/msgs/iservice"
	"github.com/bianjieai/irita-sync/msgs/nft"
	"github.com/bianjieai/irita-sync/msgs/record"
	"github.com/bianjieai/irita-sync/msgs/token"
	"github.com/bianjieai/irita-sync/msgs/bank"
)

func HandleTxMsg(v types.Msg) (MsgDocInfo) {
	if BankDocInfo, ok := bank.HandleTxMsg(v); ok {
		return BankDocInfo
	}
	if IServiceDocInfo, ok := iservice.HandleTxMsg(v); ok {
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
