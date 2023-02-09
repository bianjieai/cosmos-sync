package models

import (
	"fmt"
	"github.com/bianjieai/cosmos-sync/utils/constant"
	"github.com/kaifei-bianjie/common-parser/types"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	CollectionNameTx = "sync_tx"
)

type (
	Tx struct {
		TxId            int64         `bson:"tx_id"`
		Time            int64         `bson:"time"`
		Height          int64         `bson:"height"`
		TxHash          string        `bson:"tx_hash"`
		Type            string        `bson:"type"` // parse from first msg
		Memo            string        `bson:"memo"`
		Status          uint32        `bson:"status"`
		Log             string        `bson:"log"`
		Fee             *types.Fee    `bson:"fee"`
		FeePayer        string        `bson:"fee_payer"`
		FeeGranter      string        `bson:"fee_granter"`
		FeeGrantee      string        `bson:"fee_grantee"`
		GasUsed         int64         `bson:"gas_used"`
		Types           []string      `bson:"types"`
		EventsNew       []EventNew    `bson:"events_new"`
		Signers         []string      `bson:"signers"`
		DocTxMsgs       []types.TxMsg `bson:"msgs"`
		Addrs           []string      `bson:"addrs"`
		ContractAddrs   []string      `bson:"contract_addrs"`
		EvmTxRespondRet string        `bson:"evm_tx_respond_ret"`
		TxIndex         uint32        `bson:"tx_index"`
		Ext             interface{}   `bson:"ext"`
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
	if GetSrvConf().ChainId == "" {
		return CollectionNameTx
	}
	return fmt.Sprintf("sync_%v_tx", GetSrvConf().ChainId)
}

func (d Tx) EnsureIndexes() {
	var indexes []options.IndexModel
	indexes = append(indexes, options.IndexModel{
		Key:        []string{"-tx_hash", "-height"},
		Unique:     true,
		Background: true,
	})
	indexes = append(indexes, options.IndexModel{
		Key:        []string{"-height"},
		Background: true,
	})

	ensureIndexes(d.Name(), indexes)
}

func (d Tx) PkKvPair() map[string]interface{} {
	return bson.M{"tx_hash": d.TxHash}
}

func (d Tx) FindIncorrectParseTxs() (bool, error) {
	var incorrectTxs []Tx
	q := bson.M{
		"msgs.type": constant.IncorrectParse,
		"time":      bson.M{"$lt": time.Now().Unix()},
	}

	fn := func(c *qmgo.Collection) error {
		return c.Find(_ctx, q).Select(bson.M{"msgs.type": 1}).Limit(10).All(&incorrectTxs)
	}

	err := ExecCollection(d.Name(), fn)

	if err != nil {
		return false, err
	}

	return len(incorrectTxs) > 0, nil
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
