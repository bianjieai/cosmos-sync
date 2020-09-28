package server

import (
	constant "github.com/bianjieai/irita-sync/confs"
	"github.com/bianjieai/irita-sync/libs/logger"
	"github.com/bianjieai/irita-sync/utils"
	"os"
	"strings"
)

type ServerConf struct {
	NodeUrls                []string
	WorkerNumCreateTask     int
	WorkerNumExecuteTask    int
	WorkerMaxSleepTime      int
	BlockNumPerWorkerHandle int

	MaxConnectionNum  int
	InitConnectionNum int
	Bech32AccPrefix   string
	ChainId           string
}

var (
	SvrConf *ServerConf

	nodeUrls                = []string{"tcp://58.33.6.114:26657"}
	workerNumExecuteTask    = 30
	workerMaxSleepTime      = 2 * 60
	blockNumPerWorkerHandle = 100
	bech32AccPrefix         = "iaa"
	chainId                 = ""
)

// get value of env var
func init() {
	if v, ok := os.LookupEnv(constant.EnvNameSerNetworkFullNodes); ok {
		nodeUrls = strings.Split(v, ",")
	}


	if v, ok := os.LookupEnv(constant.EnvNameWorkerNumExecuteTask); ok {
		if n, err := utils.ConvStrToInt(v); err != nil {
			logger.Fatal("convert str to int fail", logger.String(constant.EnvNameWorkerNumExecuteTask, v))
		} else {
			workerNumExecuteTask = n
		}
	}

	if v, ok := os.LookupEnv(constant.EnvNameWorkerMaxSleepTime); ok {
		if n, err := utils.ConvStrToInt(v); err != nil {
			logger.Fatal("convert str to int fail", logger.String(constant.EnvNameWorkerMaxSleepTime, v))
		} else {
			workerMaxSleepTime = n
		}
	}

	if v, ok := os.LookupEnv(constant.EnvNameBlockNumPerWorkerHandle); ok {
		if n, err := utils.ConvStrToInt(v); err != nil {
			logger.Fatal("convert str to int fail", logger.String(constant.EnvNameBlockNumPerWorkerHandle, v))
		} else {
			blockNumPerWorkerHandle = n
		}
	}

	if v, ok := os.LookupEnv(constant.EnvNameBech32AccPrefix); ok {
		bech32AccPrefix = v
	}

	if v, ok := os.LookupEnv(constant.EnvNameChainId); ok {
		chainId = v
	}

	SvrConf = &ServerConf{
		NodeUrls:                nodeUrls,
		WorkerNumCreateTask:     1,
		WorkerNumExecuteTask:    workerNumExecuteTask,
		WorkerMaxSleepTime:      workerMaxSleepTime,
		BlockNumPerWorkerHandle: blockNumPerWorkerHandle,

		MaxConnectionNum:  100,
		InitConnectionNum: 5,

		Bech32AccPrefix: bech32AccPrefix,
		ChainId:         chainId,
	}

	logger.Debug("print server config", logger.String("serverConf", utils.MarshalJsonIgnoreErr(SvrConf)))
}
