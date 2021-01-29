package handlers

import (
	"github.com/bianjieai/irita-sync/libs/msgsdk"
	. "github.com/weichang-bianjie/msg-sdk/modules"
	"github.com/weichang-bianjie/msg-sdk/types"
	"gopkg.in/mgo.v2/txn"
)

func HandleTxMsg(v types.SdkMsg) (MsgDocInfo, []txn.Op) {
	if BankDocInfo, ok := msgsdk.MsgClient.Bank.HandleTxMsg(v); ok {
		return BankDocInfo, nil
	}
	if CrisisDocInfo, ok := msgsdk.MsgClient.Crisis.HandleTxMsg(v); ok {
		return CrisisDocInfo, nil
	}
	if DistrubutionDocInfo, ok := msgsdk.MsgClient.Distribution.HandleTxMsg(v); ok {
		return DistrubutionDocInfo, nil
	}
	if SlashingDocInfo, ok := msgsdk.MsgClient.Slashing.HandleTxMsg(v); ok {
		return SlashingDocInfo, nil
	}
	if EvidenceDocInfo, ok := msgsdk.MsgClient.Evidence.HandleTxMsg(v); ok {
		return EvidenceDocInfo, nil
	}
	if StakingDocInfo, ok := msgsdk.MsgClient.Staking.HandleTxMsg(v); ok {
		return StakingDocInfo, nil
	}
	if GovDocInfo, ok := msgsdk.MsgClient.Gov.HandleTxMsg(v); ok {
		return GovDocInfo, nil
	}

	//if WasmDocInfo, ok := wasm.HandleTxMsg(v); ok {
	//	return WasmDocInfo, nil
	//}
	if IbcDocinfo, ok := msgsdk.MsgClient.Ibc.HandleTxMsg(v); ok {
		return IbcDocinfo, nil
	}
	return MsgDocInfo{}, nil
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
