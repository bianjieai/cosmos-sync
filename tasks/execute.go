package tasks

import (
	"context"
	"fmt"
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/handlers"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/libs/stream"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/bianjieai/cosmos-sync/utils"
	"github.com/bianjieai/cosmos-sync/utils/constant"
	. "github.com/kaifei-bianjie/iritachain-mod-parser/modules"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strings"
	"time"
)

func (s *syncTaskService) StartExecuteTask() {
	defer func() {
		pool.ClosePool()
	}()

	var (
		blockNumPerWorkerHandle = int64(s.conf.Server.BlockNumPerWorkerHandle)
		workerMaxSleepTime      = int64(s.conf.Server.WorkerMaxSleepTime)
	)
	if workerMaxSleepTime <= 1*60 {
		logger.Fatal("workerMaxSleepTime shouldn't less than 1 minute")
	}

	logger.Info("init execute task")
	s.hostname, _ = os.Hostname()

	// buffer channel to limit goroutine num
	chanLimit := make(chan bool, s.conf.Server.WorkerNumExecuteTask)

	for {
		chanLimit <- true
		go s.executeTask(blockNumPerWorkerHandle, workerMaxSleepTime, chanLimit)
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func (s *syncTaskService) executeTask(blockNumPerWorkerHandle, maxWorkerSleepTime int64, chanLimit chan bool) {
	//var (
	//	workerId, taskType     string
	//	blockChainLatestHeight int64
	//)
	//genWorkerId := func() string {
	//	// generate worker id use hostname@xxx
	//	hostname, _ := os.Hostname()
	//	return fmt.Sprintf("%v@%v", hostname, bson.NewObjectId().Hex())
	//}

	healthCheckQuit := make(chan bool)
	//workerId = genWorkerId()
	client := pool.GetClient()

	defer func() {
		if r := recover(); r != nil {
			logger.Error("execute task fail", logger.Any("err", r))
		}
		close(healthCheckQuit)
		<-chanLimit
		client.Release()
	}()

	// check whether exist executable task
	// status = unhandled or
	// status = underway and now - lastUpdateTime > confTime
	tasks, err := s.syncTaskModel.GetExecutableTask(maxWorkerSleepTime)
	if err != nil {
		logger.Error("Get executable task fail", logger.String("err", err.Error()))
		return
	}
	if len(tasks) == 0 {
		// there is no executable tasks
		return
	}

	// take over sync task
	// attempt to update status, worker_id and worker_logs
	task := tasks[utils.RandInt(len(tasks))]
	s.TakeOverTaskAndExecute(task, client, healthCheckQuit, blockNumPerWorkerHandle)
}
func (s *syncTaskService) TakeOverTaskAndExecute(task models.SyncTask, client *pool.Client, healthCheckQuit chan bool, blockNumPerWorkerHandle int64) {
	var taskType string
	workerId := fmt.Sprintf("%v@%v", s.hostname, primitive.NewObjectID().Hex())
	err := s.syncTaskModel.TakeOverTask(task, workerId)
	if err != nil {
		if err == qmgo.ErrNoSuchDocuments {
			// this task has been take over by other goroutine
			logger.Info("Task has been take over by other goroutine")
		} else {
			logger.Error("Take over task fail", logger.String("err", err.Error()))
		}
		return
	} else {
		// task over task success, update task worker to current worker
		task.WorkerId = workerId
	}

	if task.EndHeight != 0 {
		taskType = models.SyncTaskTypeCatchUp
	} else {
		taskType = models.SyncTaskTypeFollow
	}
	logger.Info("worker begin execute task",
		logger.String("curWorker", workerId), logger.Any("taskId", task.ID),
		logger.String("from-to", fmt.Sprintf("%v-%v", task.StartHeight, task.EndHeight)))

	// worker health check, if worker is alive, then update last update time every minute.
	// health check will exit in follow conditions:
	// 1. task is not owned by current worker
	// 2. task is invalid
	workerHealthCheck := func(taskId primitive.ObjectID, currentWorker string) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("worker health check err", logger.Any("err", r))
			}
		}()

		func() {
			for {
				select {
				case <-healthCheckQuit:
					logger.Info("get health check quit signal, now exit health check")
					return
				default:
					task, err := s.syncTaskModel.GetTaskByIdAndWorker(taskId, workerId)
					if err == nil {
						if _, valid := assertTaskValid(task, blockNumPerWorkerHandle); valid {
							// update task last update time
							if err := s.syncTaskModel.UpdateLastUpdateTime(task); err != nil {
								logger.Error("update last update time fail", logger.String("err", err.Error()))
							}
						} else {
							logger.Info("task is invalid, exit health check", logger.String("taskId", taskId.Hex()))
							return
						}
					} else {
						if err == qmgo.ErrNoSuchDocuments {
							logger.Info("task may be task over by other goroutine, exit health check",
								logger.String("taskId", taskId.Hex()), logger.String("curWorker", workerId))
							return
						} else {
							logger.Error("get task by id and worker fail", logger.String("taskId", taskId.Hex()),
								logger.String("curWorker", workerId))
						}
					}
				}
				time.Sleep(1 * time.Minute)
			}
		}()
	}
	go workerHealthCheck(task.ID, workerId)

	// check task is valid
	// valid catch up task: current_height < end_height
	// valid follow task: current_height + blockNumPerWorkerHandle > blockChainLatestHeight
	blockChainLatestHeight, isValid := assertTaskValid(task, blockNumPerWorkerHandle)
	for isValid {
		var inProcessBlock int64
		if task.CurrentHeight == 0 {
			inProcessBlock = task.StartHeight
		} else {
			inProcessBlock = task.CurrentHeight + 1
		}

		// if inProcessBlock > blockChainLatestHeight, should wait blockChainLatestHeight update
		if taskType == models.SyncTaskTypeFollow &&
			inProcessBlock+int64(s.conf.Server.BehindBlockNum) > blockChainLatestHeight {
			logger.Info(fmt.Sprintf("wait blockChain latest height update, must interval %v block",
				s.conf.Server.BehindBlockNum),
				logger.Int64("curSyncedHeight", inProcessBlock-1),
				logger.Int64("blockChainLatestHeight", blockChainLatestHeight))
			time.Sleep(2 * time.Second)
			// continue to assert task is valid
			blockChainLatestHeight, isValid = assertTaskValid(task, blockNumPerWorkerHandle)
			continue
		}

		// parse data from block
		blockDoc, txDocs, err := handlers.ParseBlockAndTxs(inProcessBlock, client)
		if err != nil {
			logger.Error("Parse block fail",
				logger.Int64("height", inProcessBlock),
				logger.String("errTag", utils.GetErrTag(err)),
				logger.String("err", err.Error()))
			//continue to assert task is valid
			blockChainLatestHeight, isValid = assertTaskValid(task, blockNumPerWorkerHandle)
			continue
		}

		// check task owner
		workerUnchanged, err := assertTaskWorkerUnchanged(task.ID, task.WorkerId)
		if err != nil {
			logger.Error("assert task worker is unchanged fail", logger.String("err", err.Error()))
		}
		if workerUnchanged {
			// save data and update sync task
			taskDoc := task
			taskDoc.CurrentHeight = inProcessBlock
			taskDoc.LastUpdateTime = time.Now().Unix()
			taskDoc.Status = models.SyncTaskStatusUnderway
			if inProcessBlock == task.EndHeight {
				taskDoc.Status = models.SyncTaskStatusCompleted
			}

			err := saveDocsWithTxn(blockDoc, txDocs, &taskDoc)
			if err != nil {
				if !strings.Contains(err.Error(), constant.ErrDbNotFindTransaction) {
					logger.Error("save docs fail",
						logger.Int64("height", inProcessBlock),
						logger.String("err", err.Error()))
					//continue to assert task is valid
					blockChainLatestHeight, isValid = assertTaskValid(task, blockNumPerWorkerHandle)
					continue
				}
			} else {
				task.CurrentHeight = inProcessBlock
			}

			// continue to assert task is valid
			blockChainLatestHeight, isValid = assertTaskValid(task, blockNumPerWorkerHandle)
		} else {
			logger.Info("task worker changed", logger.Any("task_id", task.ID),
				logger.String("origin worker", workerId), logger.String("current worker", task.WorkerId))
			return
		}
	}

	logger.Info("worker finish execute task",
		logger.String("task_worker", task.WorkerId), logger.Any("task_id", task.ID),
		logger.String("from-to-current", fmt.Sprintf("%v-%v-%v", task.StartHeight, task.EndHeight, task.CurrentHeight)))
}

// assert task is valid
// valid task status: "underway" or "unhandled"
// valid catch up task: current_height < end_height
// valid follow task: current_height + blockNumPerWorkerHandle > blockChainLatestHeight
func assertTaskValid(task models.SyncTask, blockNumPerWorkerHandle int64) (int64, bool) {
	var (
		taskType               string
		flag                   = false
		blockChainLatestHeight int64
		err                    error
	)
	if task.EndHeight != 0 {
		taskType = models.SyncTaskTypeCatchUp
	} else {
		taskType = models.SyncTaskTypeFollow
	}
	currentHeight := task.CurrentHeight
	if currentHeight == 0 {
		currentHeight = task.StartHeight - 1
	}

	switch taskType {
	case models.SyncTaskTypeCatchUp:
		if currentHeight < task.EndHeight {
			flag = true
		}
		break
	case models.SyncTaskTypeFollow:
		blockChainLatestHeight, err = getBlockChainLatestHeight()
		if err != nil {
			logger.Error("get blockChain latest height err", logger.String("err", err.Error()))
			return blockChainLatestHeight, flag
		}
		if currentHeight+blockNumPerWorkerHandle > blockChainLatestHeight {
			flag = true
		}
		break
	}
	//judge task status is "underway" or "unhandled" when flag is true
	if flag {
		taskStatusUnderway := task.Status == models.SyncTaskStatusUnderway
		taskStatusUnhandled := task.Status == models.SyncTaskStatusUnHandled
		flag = taskStatusUnderway || taskStatusUnhandled
	}
	return blockChainLatestHeight, flag
}

// assert task worker unchanged
func assertTaskWorkerUnchanged(taskId primitive.ObjectID, workerId string) (bool, error) {
	var (
		syncTaskModel models.SyncTask
	)
	// check task owner
	task, err := syncTaskModel.GetTaskById(taskId)
	if err != nil {
		return false, err
	}

	//workid is not change and status is "underway"
	if task.WorkerId == workerId && task.Status == models.SyncTaskStatusUnderway {
		return true, nil
	} else {
		return false, nil
	}
}

// get current block height
func getBlockChainLatestHeight() (int64, error) {
	client := pool.GetClient()
	defer func() {
		client.Release()
	}()
	status, err := client.Status(context.Background())
	if err != nil {
		return 0, err
	}

	return status.SyncInfo.LatestBlockHeight, nil
}

func saveDocsWithTxn(blockDoc *models.Block, txDocs []*models.Tx, taskDoc *models.SyncTask) error {
	if blockDoc.Height == 0 {
		return fmt.Errorf("invalid block, height equal 0")
	}

	txEvms := dealEvmTx(txDocs)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := models.GetClient().DoTransaction(ctx, func(sessCtx context.Context) (interface{}, error) {
		blockCli := models.GetClient().Database(models.GetDbConf().Database).Collection(models.Block{}.Name())
		txCli := models.GetClient().Database(models.GetDbConf().Database).Collection(models.Tx{}.Name())
		taskCli := models.GetClient().Database(models.GetDbConf().Database).Collection(models.SyncTask{}.Name())
		txEvmCli := models.GetClient().Database(models.GetDbConf().Database).Collection(models.SyncTxEvm{}.Name())
		if _, err := blockCli.InsertOne(sessCtx, blockDoc); err != nil {
			return nil, err
		}
		sizeTxDocs := len(txDocs)
		if sizeTxDocs == 1 {
			if _, err := txCli.InsertOne(sessCtx, txDocs[0]); err != nil {
				return nil, err
			}
		} else if sizeTxDocs > 1 {
			if _, err := txCli.InsertMany(sessCtx, txDocs); err != nil {
				return nil, err
			}
		}

		if len(txEvms) == 1 {
			if _, err := txEvmCli.InsertOne(sessCtx, txEvms[0]); err != nil {
				return nil, err
			}
		} else if len(txEvms) > 1 {
			if _, err := txEvmCli.InsertMany(sessCtx, txEvms); err != nil {
				return nil, err
			}
		}

		cond := bson.M{"_id": taskDoc.ID, "status": bson.M{"$in": []string{models.SyncTaskStatusUnderway, models.SyncTaskStatusUnHandled}}}
		update := bson.M{
			"$set": bson.M{
				"current_height":   taskDoc.CurrentHeight,
				"status":           taskDoc.Status,
				"last_update_time": taskDoc.LastUpdateTime,
			},
		}
		if err := taskCli.UpdateOne(sessCtx, cond, update); err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		logger.Error("saveDocsWithTxn DoTransaction fail",
			logger.Int64("height", blockDoc.Height),
			logger.String("err", err.Error()))
		return err
	}

	if len(txEvms) > 0 {
		if err = productMsgToMq(txEvms); err != nil {
			logger.Error("saveDocsWithTxn productMsgToMq fail",
				logger.Int64("height", blockDoc.Height),
				logger.String("err", err.Error()))
			return err
		}
	}

	return nil
}

const (
	RecordStatusUnprocessed = iota
	RecordStatusProcessing
	RecordStatusCompleted
)

func dealEvmTx(txDocs []*models.Tx) []*models.SyncTxEvm {
	txEvms := make([]*models.SyncTxEvm, 0, len(txDocs))
	for _, doc := range txDocs {
		if doc.Type == MsgTypeEthereumTx {
			if doc.Status == constant.TxStatusSuccess {
				var txEvm models.SyncTxEvm
				txEvm.Height = doc.Height
				txEvm.Types = doc.Types
				txEvm.TxHash = doc.TxHash

				var evmTxHash string
				for _, eventNew := range doc.EventsNew {
					evmTxHash = eventNew.GetValue(MsgTypeEthereumTx, constant.EthereumTxHash)
				}

				txEvm.EvmTxHash = evmTxHash
				txEvm.Time = doc.Time
				txEvm.Status = doc.Status
				txEvm.Memo = doc.Memo
				txEvm.Signers = doc.Signers
				txEvm.TxId = doc.TxId
				txEvm.GasUsed = doc.GasUsed
				txEvm.ContractAddress = doc.ContractAddrs[0]
				txEvm.Fee = doc.Fee
				txEvm.RecordStatus = RecordStatusUnprocessed
				txEvm.CreateTime = time.Now().Unix()
				txEvm.UpdateTime = time.Now().Unix()
				txEvms = append(txEvms, &txEvm)
			}
		}

	}
	return txEvms
}

func productMsgToMq(txEvms []*models.SyncTxEvm) error {
	_, err := models.GetClient().DoTransaction(context.Background(), func(sessCtx context.Context) (interface{}, error) {
		for _, txEvm := range txEvms {
			streamLen, err := stream.GetClient().GetStreamLen(config.GetConfig().Redis.StreamTxEvmKey)
			if err != nil {
				logger.Error("productMsgToMq stream GetStreamLen fail",
					logger.String("err", err.Error()))
				return nil, err
			}

			if streamLen >= config.GetConfig().Redis.StreamMqMaxLen {
				//TODO 队列满了，需要监控告警
				logger.Debug("productMsgToMq streamLen >= StreamMqMaxLen", logger.Int64("maxLen", config.GetConfig().Redis.StreamMqMaxLen))
				return nil, nil
			}

			_, err = stream.GetClient().PutMsg(config.GetConfig().Redis.StreamTxEvmKey, txEvm.GetStreamMap())
			if err != nil {
				logger.Error("productMsgToMq stream PutMsg fail",
					logger.Int64("height", txEvm.Height),
					logger.String("evm_tx_hash", txEvm.EvmTxHash),
					logger.String("err", err.Error()))
				return nil, err
			}

			txEvmCli := models.GetClient().Database(models.GetDbConf().Database).Collection(models.SyncTxEvm{}.Name())

			query := bson.M{"height": txEvm.Height, "evm_tx_hash": txEvm.EvmTxHash, "record_status": RecordStatusUnprocessed}
			update := bson.M{
				"$set": bson.M{
					"record_status": RecordStatusProcessing,
					"update_time":   time.Now().Unix(),
				},
			}
			if err := txEvmCli.UpdateOne(sessCtx, query, update); err != nil {
				return nil, err
			}
		}

		return nil, nil
	})

	return err
}
