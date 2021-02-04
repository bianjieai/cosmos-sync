// interface for a document

package models

const (
	CollectionNameTxn = "sync_txn"
)

var (
	SyncTaskModel SyncTask
	BlockModel    Block
	TxModel       Tx

	Collections = []Docs{
		SyncTaskModel,
		BlockModel,
		TxModel,
	}
)

type (
	Docs interface {
		// collection name
		Name() string
		// ensure indexes
		EnsureIndexes()
		// primary key pair(used to find a unique record)
		PkKvPair() map[string]interface{}
	}
)


func ensureDocsIndexes() {
	if len(Collections) > 0 {
		for _, v := range Collections {
			v.EnsureIndexes()
		}
	}
}