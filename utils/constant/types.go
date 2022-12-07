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
	ErrDbNotFindTransaction = "cannot find transaction"

	IbcTransferEventTypeSendPacket           = "send_packet"
	IbcTransferEventAttriKeyPacketSequence   = "packet_sequence"
	IbcTransferEventAttriKeyPacketScPort     = "packet_src_port"
	IbcTransferEventAttriKeyPacketScChannel  = "packet_src_channel"
	IbcTransferEventAttriKeyPacketDcPort     = "packet_dst_port"
	IbcTransferEventAttriKeyPacketDcChannels = "packet_dst_channel"
	IbcRecvPacketEventTypeWriteAcknowledge   = "write_acknowledgement"
	IbcRecvPacketEventAttriKeyPacketAck      = "packet_ack"
	IbcUpdateClientEventTypeUpdateClient     = "update_client"
	IbcUpdateClientEventAttriKeyHeader       = "header"
	IbcRecvPacketEventTypeRecvPacket         = "recv_packet"
	IbcRecvPacketEventAttriKeyPacketDataHex  = "packet_data_hex"
)
