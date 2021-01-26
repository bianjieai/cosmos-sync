package handlers

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/msgs/bank"
	"github.com/bianjieai/irita-sync/msgs/crisis"
	"github.com/bianjieai/irita-sync/msgs/distribution"
	"github.com/bianjieai/irita-sync/msgs/evidence"
	"github.com/bianjieai/irita-sync/msgs/gov"
	"github.com/bianjieai/irita-sync/msgs/ibc"
	"github.com/bianjieai/irita-sync/msgs/slashing"
	"github.com/bianjieai/irita-sync/msgs/staking"
	"github.com/cosmos/cosmos-sdk/types"
	"gopkg.in/mgo.v2/txn"
)

func HandleTxMsg(v types.Msg) (MsgDocInfo, []txn.Op) {
	if BankDocInfo, ok := bank.HandleTxMsg(v); ok {
		return BankDocInfo, nil
	}
	if CrisisDocInfo, ok := crisis.HandleTxMsg(v); ok {
		return CrisisDocInfo, nil
	}
	if DistrubutionDocInfo, ok := distribution.HandleTxMsg(v); ok {
		return DistrubutionDocInfo, nil
	}
	if SlashingDocInfo, ok := slashing.HandleTxMsg(v); ok {
		return SlashingDocInfo, nil
	}
	if EvidenceDocInfo, ok := evidence.HandleTxMsg(v); ok {
		return EvidenceDocInfo, nil
	}
	if StakingDocInfo, ok := staking.HandleTxMsg(v); ok {
		return StakingDocInfo, nil
	}
	if GovDocInfo, ok := gov.HandleTxMsg(v); ok {
		return GovDocInfo, nil
	}

	//if WasmDocInfo, ok := wasm.HandleTxMsg(v); ok {
	//	return WasmDocInfo, nil
	//}
	if IbcDocinfo, ok := ibc.HandleTxMsg(v); ok {
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

