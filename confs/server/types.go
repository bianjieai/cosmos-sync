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
	SleepTimeCreateTaskWorker int
	BlockNumPerWorkerHandle   int

	MaxConnectionNum   int
	InitConnectionNum  int
	Bech32AccPrefix    string
	ChainId            string
	ChainBlockInterval int
}

var (
	SvrConf *ServerConf

	nodeUrls                = []string{"tcp://10.1.4.220:26657"}
	workerNumExecuteTask    = 30
	workerMaxSleepTime      = 2 * 60
	blockNumPerWorkerHandle = 100
	bech32AccPrefix         = "iaa"
	chainId                 = ""
	chainBlockInterval      = 5
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

	if v, ok := os.LookupEnv(constant.EnvNameChainBlockInterval); ok {
		if n, err := utils.ConvStrToInt(v); err != nil {
			logger.Fatal("convert str to int fail", logger.String(constant.EnvNameChainBlockInterval, v))
		} else {
			chainBlockInterval = n
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

		Bech32AccPrefix:    bech32AccPrefix,
		ChainId:            chainId,
		ChainBlockInterval: chainBlockInterval,
	}

	logger.Debug("print server config", logger.String("serverConf", utils.MarshalJsonIgnoreErr(SvrConf)))
}
