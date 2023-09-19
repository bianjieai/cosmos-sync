package constant

const (
	TxStatusFail = iota
	TxStatusSuccess
	EnvNameConfigFilePath = "CONFIG_FILE_PATH"

	NoSupportMsgTypeTag = "no support msg parse"

	//cannot find transaction 601bf70ccdee4dde1c8be0d2_f018677a in queue for document {sync_task ObjectIdHex(\"601bdb0ccdee4dd7c214d167\")}
	ErrDbNotFindTransaction = "cannot find transaction"

	EthereumTxHash = "ethereumTxHash"

	UseGrantee = "use_feegrant"
	SetGrantee = "set_feegrant"
	Grantee    = "grantee"
)
