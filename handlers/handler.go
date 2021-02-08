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
	if IServiceDocInfo, ok := msgparser.MsgClient.Service.HandleTxMsg(v); ok {
		return IServiceDocInfo, nil
	}
	if NftDocInfo, ok := msgparser.MsgClient.Nft.HandleTxMsg(v); ok {
		return NftDocInfo, nil
	}
	if RecordDocInfo, ok := msgparser.MsgClient.Record.HandleTxMsg(v); ok {
		return RecordDocInfo, nil
	}
	if TokenDocInfo, ok := msgparser.MsgClient.Token.HandleTxMsg(v); ok {
		return TokenDocInfo, nil
	}
	if CoinswapDocInfo, ok := msgparser.MsgClient.Coinswap.HandleTxMsg(v); ok {
		return CoinswapDocInfo, nil
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
	if HtlcDocInfo, ok := msgparser.MsgClient.Htlc.HandleTxMsg(v); ok {
		return HtlcDocInfo, nil
	}
	if StakingDocInfo, ok := msgparser.MsgClient.Staking.HandleTxMsg(v); ok {
		return StakingDocInfo, nil
	}
	if GovDocInfo, ok := msgparser.MsgClient.Gov.HandleTxMsg(v); ok {
		return GovDocInfo, nil
	}
	if IdentityDocInfo, ok := msgparser.MsgClient.Identity.HandleTxMsg(v); ok {
		return IdentityDocInfo, nil
	}
	if RandomDocInfo, ok := msgparser.MsgClient.Random.HandleTxMsg(v); ok {
		return RandomDocInfo, nil
	}
	if OracleDocInfo, ok := msgparser.MsgClient.Oracle.HandleTxMsg(v); ok {
		return OracleDocInfo, nil
	}
	if IbcDocinfo, ok := msgparser.MsgClient.Ibc.HandleTxMsg(v); ok {
		//ops := handlerIbcClient(IbcDocinfo.DocTxMsg.Type, ibcClient)
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
