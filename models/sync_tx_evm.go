package models

import (
	"fmt"
	"github.com/kaifei-bianjie/common-parser/types"
)

const (
	CollectionNameSyncTxEvm = "ex_sync_tx_evm"
)

type (
	SyncTxEvm struct {
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
		EvmData                 interface{} `bson:"evm_data"`
		RecordStatus            int         `bson:"record_status"`
		CreateTime              int64       `bson:"create_time"`
		UpdateTime              int64       `bson:"update_time"`
	}
)

func (d SyncTxEvm) Name() string {
	if GetSrvConf().ChainId == "" {
		return CollectionNameSyncTxEvm
	}
	return fmt.Sprintf("ex_sync_%v_tx_evm", GetSrvConf().ChainId)
}

func (d SyncTxEvm) EnsureIndexes() {
}

func (d SyncTxEvm) PkKvPair() map[string]interface{} {
	return nil
}

func (d SyncTxEvm) GetStreamMap() map[string]interface{} {
	return map[string]interface{}{
		"height":      d.Height,
		"evm_tx_hash": d.EvmTxHash,
	}
}
