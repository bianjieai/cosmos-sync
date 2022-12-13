package stream

import (
	"fmt"
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"testing"
)

func init() {
	conf, err := config.ReadConfig()
	if err != nil {
		logger.Fatal(err.Error())
	}

	if err = Init(conf); err != nil {
		logger.Fatal(err.Error())
	}
	InitMQClient(conf)
}

func TestStream_PutMsg(t *testing.T) {
	msg := map[string]interface{}{
		"height":      20208098,
		"evm_tx_hash": "0x8c4c89a9973bcd8ed9d64788ab629025d7c2d385352822ec987cad3e39411b3a",
	}
	_, err := GetClient().PutMsg("mq_stream", msg)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStream_GetStreamsLen(t *testing.T) {
	streamLen, err := GetClient().GetStreamLen("mq_stream")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(streamLen)
}
