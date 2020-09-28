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
	"github.com/bianjieai/irita-sync/msgs/ibc"
	"github.com/bianjieai/irita-sync/models"
	"gopkg.in/mgo.v2/txn"
	"gopkg.in/mgo.v2/bson"
)

func HandleTxMsg(v types.Msg, timestamp int64) (MsgDocInfo, []txn.Op) {
	if BankDocInfo, ok := bank.HandleTxMsg(v); ok {
		return BankDocInfo, nil
	}
	if IServiceDocInfo, ok := service.HandleTxMsg(v); ok {
		return IServiceDocInfo, nil
	}
	if NftDocInfo, ok := nft.HandleTxMsg(v); ok {
		return NftDocInfo, nil
	}
	if RecordDocInfo, ok := record.HandleTxMsg(v); ok {
		return RecordDocInfo, nil
	}
	if TokenDocInfo, ok := token.HandleTxMsg(v); ok {
		return TokenDocInfo, nil
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
		return IdentityDocInfo, nil
	}
	if IbcDocinfo, ibcClient, ok := ibc.HandleTxMsg(v, timestamp); ok {
		ops := handlerIbcClient(IbcDocinfo.DocTxMsg.Type, ibcClient)
		return IbcDocinfo, ops
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

func handlerIbcClient(msgType string, client *models.IbcClient) (Ops []txn.Op) {
	switch msgType {
	case MsgTypeCreateClient:
		client.ID = bson.NewObjectId()
		op := txn.Op{
			C:      models.CollectionNameIbcClient,
			Id:     bson.NewObjectId(),
			Insert: client,
		}
		Ops = append(Ops, op)
	case MsgTypeUpdateClient:
		v := client
		mapObjId, err := client.AllIbcClientMaps()
		if err != nil {
			return
		}
		if id, ok := mapObjId[v.ClientId]; ok {
			v.ID = id
		}
		if !v.ID.Valid() {
			return
		}
		updateOp := txn.Op{
			C:      models.CollectionNameIbcClient,
			Id:     v.ID,
			Assert: txn.DocExists,
			Update: bson.M{
				"$set": bson.M{
					models.IbcClientHeaderTag:   v.Header,
					models.IbcClientSignerTag:   v.Signer,
					models.IbcClientUpdateAtTag: v.UpdateAt,
				},
			},
		}
		Ops = append(Ops, updateOp)
	}
	return
}
