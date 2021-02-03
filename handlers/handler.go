package handlers

import (
	"github.com/bianjieai/irita-sync/libs/msgparser"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/kaifei-bianjie/msg-parser/types"
	"gopkg.in/mgo.v2/txn"
)

func HandleTxMsg(v types.SdkMsg) (MsgDocInfo, []txn.Op) {
	if BankDocInfo, ok := msgparser.MsgClient.Bank.HandleTxMsg(v); ok {
		return BankDocInfo, nil
	}
	if CrisisDocInfo, ok := msgparser.MsgClient.Crisis.HandleTxMsg(v); ok {
		return CrisisDocInfo, nil
	}
	if DistrubutionDocInfo, ok := msgparser.MsgClient.Distribution.HandleTxMsg(v); ok {
		return DistrubutionDocInfo, nil
	}
	if SlashingDocInfo, ok := msgparser.MsgClient.Slashing.HandleTxMsg(v); ok {
		return SlashingDocInfo, nil
	}
	if EvidenceDocInfo, ok := msgparser.MsgClient.Evidence.HandleTxMsg(v); ok {
		return EvidenceDocInfo, nil
	}
	if StakingDocInfo, ok := msgparser.MsgClient.Staking.HandleTxMsg(v); ok {
		return StakingDocInfo, nil
	}
	if GovDocInfo, ok := msgparser.MsgClient.Gov.HandleTxMsg(v); ok {
		return GovDocInfo, nil
	}

	//if WasmDocInfo, ok := wasm.HandleTxMsg(v); ok {
	//	return WasmDocInfo, nil
	//}
	if IbcDocinfo, ok := msgparser.MsgClient.Ibc.HandleTxMsg(v); ok {
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
