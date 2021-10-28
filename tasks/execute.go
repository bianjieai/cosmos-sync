package tasks

import (
	"context"
	"fmt"
	"github.com/bianjieai/cosmos-sync/handlers"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/bianjieai/cosmos-sync/utils"
	"github.com/bianjieai/cosmos-sync/utils/constant"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
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
	workerId := fmt.Sprintf("%v@%v", s.hostname, bson.NewObjectId().Hex())
	err := s.syncTaskModel.TakeOverTask(task, workerId)
	if err != nil {
		if err == mgo.ErrNotFound {
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
	workerHealthCheck := func(taskId bson.ObjectId, currentWorker string) {
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
						if err == mgo.ErrNotFound {
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
		blockDoc, txDocs, ops, err := handlers.ParseBlockAndTxs(inProcessBlock, client)
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

			err := saveDocsWithTxn(blockDoc, txDocs, taskDoc, ops)
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
	return blockChainLatestHeight, flag
}

// assert task worker unchanged
func assertTaskWorkerUnchanged(taskId bson.ObjectId, workerId string) (bool, error) {
	var (
		syncTaskModel models.SyncTask
	)
	// check task owner
	task, err := syncTaskModel.GetTaskById(taskId)
	if err != nil {
		return false, err
	}

	if task.WorkerId == workerId {
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

func saveDocsWithTxn(blockDoc *models.Block, txDocs []*models.Tx, taskDoc models.SyncTask, opsDoc []txn.Op) error {
	var (
		ops, binanceTxsOps []txn.Op
	)

	if blockDoc.Height == 0 {
		return fmt.Errorf("invalid block, height equal 0")
	}

	blockOp := txn.Op{
		C:      models.BlockModel.Name(),
		Id:     bson.NewObjectId(),
		Insert: blockDoc,
	}

	if length := len(txDocs); length > 0 {

		binanceTxsOps = make([]txn.Op, 0, length)
		for _, v := range txDocs {
			op := txn.Op{
				C:      models.TxModel.Name(),
				Id:     bson.NewObjectId(),
				Insert: v,
			}
			binanceTxsOps = append(binanceTxsOps, op)
		}
	}

	updateOp := txn.Op{
		C:      models.SyncTaskModel.Name(),
		Id:     taskDoc.ID,
		Assert: txn.DocExists,
		Update: bson.M{
			"$set": bson.M{
				"current_height":   taskDoc.CurrentHeight,
				"status":           taskDoc.Status,
				"last_update_time": taskDoc.LastUpdateTime,
			},
		},
	}

	ops = make([]txn.Op, 0, len(binanceTxsOps)+2)
	ops = append(append(ops, blockOp, updateOp), binanceTxsOps...)
	if len(opsDoc) > 0 {
		ops = append(ops, opsDoc...)
	}

	if len(ops) > 0 {
		err := models.Txn(ops)
		if err != nil {
			return err
		}
	}

	return nil
}
