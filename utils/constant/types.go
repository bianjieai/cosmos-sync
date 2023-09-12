package constant

const (
	TxStatusFail = iota
	TxStatusSuccess
	EnvNameConfigFilePath = "CONFIG_FILE_PATH"

	NoSupportModule     = "no_support"
	IncorrectParse      = "incorrect_parse"
	NoSupportMsgTypeTag = "no support msg parse"
	//unable to resolve type URL /cosmos.bank.v1beta1.MsgSend
	ErrNoSupportTxPrefix = "unable to resolve type URL"

	//cannot find transaction 601bf70ccdee4dde1c8be0d2_f018677a in queue for document {sync_task ObjectIdHex(\"601bdb0ccdee4dd7c214d167\")}
	ErrDbNotFindTransaction                = "cannot find transaction"
	IbcRecvPacketEventTypeWriteAcknowledge = "write_acknowledgement"
	IbcRecvPacketEventAttriKeyPacketAck    = "packet_ack"

	EthereumTxHash = "ethereumTxHash"

	UseGrantee = "use_feegrant"
	SetGrantee = "set_feegrant"
	Grantee    = "grantee"

	BlockGas = "block_gas"

	AttrKeyAmount = "amount"

	DefaultBlockGasUsed = "0"
)
