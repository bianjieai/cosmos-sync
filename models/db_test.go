package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
	"testing"
	"time"
)

func TestTxn(t *testing.T) {
	var ops []txn.Op

	type SyncTask struct {
		ID             bson.ObjectId `bson:"_id"`
		StartHeight    int64         `bson:"start_height"`     // task start height
		EndHeight      int64         `bson:"end_height"`       // task end height
		CurrentHeight  int64         `bson:"current_height"`   // task current height
		Status         string        `bson:"status"`           // task status
		WorkerId       string        `bson:"worker_id"`        // worker id
		LastUpdateTime int64         `bson:"last_update_time"` // unix timestamp
	}

	syncIrisTasks := []SyncTask{
		{
			StartHeight:    34339000,
			EndHeight:      34340000,
			Status:         "unhandled",
			CurrentHeight:  0,
			WorkerId:       "yyyyyyyy1@xxxxx",
			LastUpdateTime: time.Now().Unix(),
		},
	}

	for _, v := range syncIrisTasks {
		objectId := bson.NewObjectId()
		v.ID = objectId
		op := txn.Op{
			C:      CollectionNameSyncTask,
			Id:     objectId,
			Assert: nil,
			Insert: v,
		}

		ops = append(ops, op)
	}
	if err := Txn(ops); err != nil {
		t.Fatal(err)
	}
}
