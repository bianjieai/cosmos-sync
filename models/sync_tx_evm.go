package models

import (
	"github.com/bianjieai/cosmos-sync/libs/msgparser/types"
)

const (
	CollectionNameExEvmTx = "ex_evm_tx"
)

type (
	EvmTx struct {
		Height                  int64       `bson:"height"`
		Types                   []string    `bson:"types"`
		TxHash                  string      `bson:"tx_hash"`
		EvmTxHash               string      `bson:"evm_tx_hash"`
		Time                    int64       `bson:"time"`
		Status                  uint32      `bson:"status"`
		Memo                    string      `bson:"memo"`
		Signers                 []string    `bson:"signers"`
		TxId                    int64       `bson:"tx_id"`
		Payer                   string      `bson:"payer"`
		GasPrice                int64       `bson:"gas_price"`
		GasLimit                int64       `bson:"gas_limit"`
		GasUsed                 int64       `bson:"gas_used"`
		Nonce                   int64       `bson:"nonce"`
		IsTransfer              int64       `bson:"is_transfer"`
		ContractAddress         string      `bson:"contract_address"`
		RelationContractAddress []string    `bson:"relation_contract_address"`
		Fee                     *types.Fee  `bson:"fee"`
		EvmDatas                interface{} `bson:"evm_datas"`
		RecordStatus            int         `bson:"record_status"`
		CreateTime              int64       `bson:"create_time"`
		UpdateTime              int64       `bson:"update_time"`
	}
)

func (d EvmTx) Name() string {
	return CollectionNameExEvmTx
}

func (d EvmTx) EnsureIndexes() {
}

func (d EvmTx) PkKvPair() map[string]interface{} {
	return nil
}

func (d EvmTx) GetStreamMap() map[string]interface{} {
	return map[string]interface{}{
		"height":      d.Height,
		"evm_tx_hash": d.EvmTxHash,
	}
}
