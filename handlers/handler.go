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
	if IServiceDocInfo, ok := msgsdk.MsgClient.Service.HandleTxMsg(v); ok {
		return IServiceDocInfo, nil
	}
	if NftDocInfo, ok := msgsdk.MsgClient.Nft.HandleTxMsg(v); ok {
		return NftDocInfo, nil
	}
	if RecordDocInfo, ok := msgsdk.MsgClient.Record.HandleTxMsg(v); ok {
		return RecordDocInfo, nil
	}
	if TokenDocInfo, ok := msgsdk.MsgClient.Token.HandleTxMsg(v); ok {
		return TokenDocInfo, nil
	}
	if CoinswapDocInfo, ok := msgsdk.MsgClient.Coinswap.HandleTxMsg(v); ok {
		return CoinswapDocInfo, nil
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
	if HtlcDocInfo, ok := msgsdk.MsgClient.Htlc.HandleTxMsg(v); ok {
		return HtlcDocInfo, nil
	}
	if StakingDocInfo, ok := msgsdk.MsgClient.Staking.HandleTxMsg(v); ok {
		return StakingDocInfo, nil
	}
	if GovDocInfo, ok := msgsdk.MsgClient.Gov.HandleTxMsg(v); ok {
		return GovDocInfo, nil
	}
	//if IdentityDocInfo, ok := identity.HandleTxMsg(v); ok {
	//	return IdentityDocInfo, nil
	//}
	if RandomDocInfo, ok := msgsdk.MsgClient.Random.HandleTxMsg(v); ok {
		return RandomDocInfo, nil
	}
	if OracleDocInfo, ok := msgsdk.MsgClient.Oracle.HandleTxMsg(v); ok {
		return OracleDocInfo, nil
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

//
//func handlerIbcClient(msgType string, client *models.IbcClient) (Ops []txn.Op) {
//	switch msgType {
//	case MsgTypeCreateClient:
//		client.ID = bson.NewObjectId()
//		op := txn.Op{
//			C:      models.CollectionNameIbcClient,
//			Id:     bson.NewObjectId(),
//			Insert: client,
//		}
//		Ops = append(Ops, op)
//	case MsgTypeUpdateClient:
//		v := client
//		mapObjId, err := client.AllIbcClientMaps()
//		if err != nil {
//			return
//		}
//		if id, ok := mapObjId[v.ClientId]; ok {
//			v.ID = id
//		}
//		if !v.ID.Valid() {
//			return
//		}
//		updateOp := txn.Op{
//			C:      models.CollectionNameIbcClient,
//			Id:     v.ID,
//			Assert: txn.DocExists,
//			Update: bson.M{
//				"$set": bson.M{
//					models.IbcClientHeaderTag:   v.Header,
//					models.IbcClientSignerTag:   v.Signer,
//					models.IbcClientUpdateAtTag: v.UpdateAt,
//				},
//			},
//		}
//		Ops = append(Ops, updateOp)
//	}
//	return
//}
