package models

import (
	"fmt"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	CollectionNameBlock = "sync_block"
)

type (
	Block struct {
		Height   int64  `bson:"height"`
		Hash     string `bson:"hash"`
		Txn      int64  `bson:"txn"`
		Time     int64  `bson:"time"`
		Proposer string `bson:"proposer"`
		GasUsed  string `bson:"gas_used"`
	}
)

func (d Block) Name() string {
	if GetSrvConf().ChainId == "" {
		return CollectionNameBlock
	}
	return fmt.Sprintf("sync_%v_block", GetSrvConf().ChainId)
}

func (d Block) EnsureIndexes() {
	var indexes []options.IndexModel
	indexes = append(indexes, options.IndexModel{
		Key:        []string{"-height"},
		Unique:     true,
		Background: true,
	})
	ensureIndexes(d.Name(), indexes)
}

func (d Block) PkKvPair() map[string]interface{} {
	return bson.M{"height": d.Height}
}

func (d Block) GetMaxBlockHeight() (Block, error) {
	var result Block

	getMaxBlockHeightFn := func(c *qmgo.Collection) error {
		return c.Find(_ctx, bson.M{}).Select(bson.M{"height": 1, "time": 1}).Sort("-height").Limit(1).One(&result)
	}

	if err := ExecCollection(d.Name(), getMaxBlockHeightFn); err != nil {
		return result, err
	}

	return result, nil
}
