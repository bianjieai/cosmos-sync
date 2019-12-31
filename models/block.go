package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionNameBlock = "sync_block"
)

type (
	Block struct {
		Height int64     `bson:"height"`
		Hash   string    `bson:"hash"`
		Txn    int64     `bson:"txn"`
		Time   time.Time `bson:"time"`
	}
)

func (d Block) Name() string {
	return CollectionNameBlock
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
