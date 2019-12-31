package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionNameBinanceTx = "sync_tx"
)

type (
	Tx struct {
		Time       time.Time `bson:"time"`
		Height     int64     `bson:"height"`
		TxHash     string    `bson:"tx_hash"`
		Memo       string    `bson:"memo"`
		Status     uint32    `bson:"status"`
		Log        string    `bson:"log"`
		ComplexMsg bool      `bson:"complex_msg"`

		Type      string     `bson:"type"`
		From      string     `bson:"from"`   // parse from first msg
		To        string     `bson:"to"`     // parse from first msg
		Coins     []Coin     `bson:"coins"`  // parse from first msg
		Signer    string     `bson:"signer"` // parse from first signer
		Events    []Event    `bson:"events"`
		DocTxMsgs []DocTxMsg `bson:"msgs"`
		Signers   []string   `bson:"signers"`
	}

	Event struct {
		Type       string   `bson:"type"`
		Attributes []KvPair `bson:"attributes"`
	}

	KvPair struct {
		Key   string `bson:"key"`
		Value string `bson:"value"`
	}

	DocTxMsg struct {
		Type string `bson:"type"`
		Msg  Msg    `bson:"msg"`
	}

	Msg interface {
		GetType() string
		BuildMsg(msg interface{})
	}
)

func (d Tx) Name() string {
	return CollectionNameBinanceTx
}

func (d Tx) EnsureIndexes() {
	var indexes []mgo.Index
	indexes = append(indexes, mgo.Index{
		Key:        []string{"-tx_hash"},
		Unique:     true,
		Background: true,
	})

	ensureIndexes(d.Name(), indexes)
}

func (d Tx) PkKvPair() map[string]interface{} {
	return bson.M{"tx_hash": d.TxHash}
}
