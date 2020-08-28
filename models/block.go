package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"github.com/bianjieai/irita-sync/confs/server"
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
	}
)

func (d Block) Name() string {
	if server.SvrConf.ChainId == "" {
		return CollectionNameBlock
	}
	return fmt.Sprintf("sync_%v_block", server.SvrConf.ChainId)
}

func (d Block) EnsureIndexes() {
	var indexes []mgo.Index
	indexes = append(indexes, mgo.Index{
		Key:        []string{"-height"},
		Unique:     true,
		Background: true,
	})
	ensureIndexes(d.Name(), indexes)
}

func (d Block) PkKvPair() map[string]interface{} {
	return bson.M{"height": d.Height}
}
