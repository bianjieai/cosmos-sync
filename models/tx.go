package models

import (
	"github.com/kaifei-bianjie/common-parser/types"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	CollectionNameTx = "sync_tx"
)

type (
	Tx struct {
		TxId          int64         `bson:"tx_id"`
		Time          int64         `bson:"time"`
		Height        int64         `bson:"height"`
		TxHash        string        `bson:"tx_hash"`
		Type          string        `bson:"type"` // parse from first msg
		Memo          string        `bson:"memo"`
		Status        uint32        `bson:"status"`
		Log           string        `bson:"log"`
		Fee           *types.Fee    `bson:"fee"`
		FeePayer      string        `bson:"fee_payer"`
		FeeGranter    string        `bson:"fee_granter"`
		FeeGrantee    string        `bson:"fee_grantee"`
		Types         []string      `bson:"types"`
		EventsNew     []EventNew    `bson:"events_new"`
		Signers       []string      `bson:"signers"`
		DocTxMsgs     []types.TxMsg `bson:"msgs"`
		Addrs         []string      `bson:"addrs"`
		ContractAddrs []string      `bson:"contract_addrs"`
		TxIndex       uint32        `bson:"tx_index"` // sequence tx of this block
		Ext           interface{}   `bson:"ext"`
		GasUsed       int64         `bson:"gas_used"`
	}

	Event struct {
		Type       string   `bson:"type"`
		Attributes []KvPair `bson:"attributes"`
	}

	KvPair struct {
		Key   string `bson:"key"`
		Value string `bson:"value"`
	}

	EventNew struct {
		MsgIndex uint32  `bson:"msg_index" json:"msg_index"`
		Events   []Event `bson:"events"`
	}
)

func (d Tx) Name() string {
	return CollectionNameTx
}

func (d Tx) EnsureIndexes() {
	var indexes []options.IndexModel
	indexes = append(indexes, options.IndexModel{
		Key:        []string{"-tx_hash"},
		Unique:     true,
		Background: true,
	})

	indexes = append(indexes, options.IndexModel{
		Key:        []string{"-height", "-tx_index"},
		Unique:     true,
		Background: true,
	})

	ensureIndexes(d.Name(), indexes)
}

func (d Tx) PkKvPair() map[string]interface{} {
	return bson.M{"tx_hash": d.TxHash}
}

func (e EventNew) GetValue(eventType, attributeKey string) string {
	for _, event := range e.Events {
		if event.Type == eventType {
			for _, attribute := range event.Attributes {
				if attribute.Key == attributeKey {
					return attribute.Value
				}
			}
		}
	}
	return ""
}
