package server

import (
	constant "github.com/bianjieai/irita-sync/confs"
	"github.com/bianjieai/irita-sync/libs/logger"
	"github.com/bianjieai/irita-sync/utils"
	"os"
	"strings"
)

type ServerConf struct {
	NodeUrls                  []string
	WorkerNumCreateTask       int
	WorkerNumExecuteTask      int
	WorkerMaxSleepTime        int
	BlockNumPerWorkerHandle   int
	SleepTimeCreateTaskWorker int

	MaxConnectionNum   int
	InitConnectionNum  int
	Bech32ChainPrefix  string
	ChainId            string
	ChainBlockInterval int
	WaitBlockNumHandle int
}

var (
	SvrConf *ServerConf

	nodeUrls                = []string{"tcp://192.168.150.31:16657"}
	workerNumExecuteTask    = 30
	workerMaxSleepTime      = 2 * 60
	blockNumPerWorkerHandle = 100
	bech32ChainPrefix       = "i"
	chainId                 = ""
	chainBlockInterval      = 5
	waitBlockNumHandle      = 1
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

	if v, ok := os.LookupEnv(constant.EnvNameBech32ChainPrefix); ok {
		bech32ChainPrefix = v
	}

	if v, ok := os.LookupEnv(constant.EnvNameChainId); ok {
		chainId = v
	}

	if v, ok := os.LookupEnv(constant.EnvNameChainBlockInterval); ok {
		if n, err := utils.ConvStrToInt(v); err != nil {
			logger.Fatal("convert str to int fail", logger.String(constant.EnvNameChainBlockInterval, v))
		} else {
			chainBlockInterval = n
		}
	}
	if v, ok := os.LookupEnv(constant.EnvNameWaitBlockNumHandle); ok {
		if n, err := utils.ConvStrToInt(v); err != nil {
			logger.Fatal("convert str to int fail", logger.String(constant.EnvNameWaitBlockNumHandle, v))
		} else {
			if n > 1 {
				waitBlockNumHandle = n
			}
		}
	}

	// calculate sleep time of create task goroutine, time unit is second
	// 1. sleepTime must less than blockNumPerWorkerHandle*chainTimeInterval,
	//    otherwise create task goroutine can't succeed create follow task
	// 2. use value of sleepTime/5 to make create task worker do task in time
	sleepTimeCreateTaskWorker := (blockNumPerWorkerHandle * chainBlockInterval) / 5
	if sleepTimeCreateTaskWorker == 0 {
		sleepTimeCreateTaskWorker = 1
	}

	SvrConf = &ServerConf{
		NodeUrls:                  nodeUrls,
		WorkerNumCreateTask:       1,
		WorkerNumExecuteTask:      workerNumExecuteTask,
		WorkerMaxSleepTime:        workerMaxSleepTime,
		BlockNumPerWorkerHandle:   blockNumPerWorkerHandle,
		SleepTimeCreateTaskWorker: sleepTimeCreateTaskWorker,

		MaxConnectionNum:  100,
		InitConnectionNum: 5,

		Bech32ChainPrefix:  bech32ChainPrefix,
		ChainId:            chainId,
		ChainBlockInterval: chainBlockInterval,
		WaitBlockNumHandle: waitBlockNumHandle,
	}

	logger.Debug("print server config", logger.String("serverConf", utils.MarshalJsonIgnoreErr(SvrConf)))
}
