package monitor

import (
	"context"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/bianjieai/cosmos-sync/monitor/metrics"
	"github.com/qiniu/qmgo"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	NodeStatusNotReachable = 0
	NodeStatusSyncing      = 1
	NodeStatusCatchingUp   = 2

	SyncTaskFollowing  = 1
	SyncTaskCatchingUp = 0
)

type clientNode struct {
	nodeStatus          metrics.Guage
	nodeHeight          metrics.Guage
	dbHeight            metrics.Guage
	nodeTimeGap         metrics.Guage
	syncWorkWay         metrics.Guage
	syncCatchingTaskNum metrics.Guage
	syncParseBugMetric  metrics.Guage
}

func NewMetricNode(server metrics.Monitor) clientNode {
	nodeHeightMetric := metrics.NewGuage(
		"sync",
		"status",
		"node_height",
		"full node latest block height",
		nil,
	)
	dbHeightMetric := metrics.NewGuage(
		"sync",
		"status",
		"db_height",
		"sync system database max block height",
		nil,
	)
	nodeStatusMetric := metrics.NewGuage(
		"sync",
		"status",
		"node_status",
		"full node status(0:NotReachable,1:Syncing,2:CatchingUp)",
		nil,
	)
	nodeTimeGapMetric := metrics.NewGuage(
		"sync",
		"status",
		"node_seconds_gap",
		"the seconds gap between node block time with sync db block time",
		nil,
	)
	syncWorkwayMetric := metrics.NewGuage(
		"sync",
		"",
		"task_working_status",
		"sync task working status(0:CatchingUp 1:Following)",
		nil,
	)
	syncCatchingTaskNumMetric := metrics.NewGuage(
		"sync",
		"",
		"task_catching_cnt",
		"count of sync catchUping task",
		nil,
	)
	syncParseBugMetric := metrics.NewGuage(
		"sync",
		"",
		"parser_occure_bug",
		"occure bug about sync parse tx",
		nil,
	)
	//server.RegisterMetrics(nodeHeightMetric, dbHeightMetric, nodeStatusMetric, nodeTimeGapMetric, syncWorkwayMetric)
	nodeHeight, _ := metrics.CovertGuage(nodeHeightMetric)
	dbHeight, _ := metrics.CovertGuage(dbHeightMetric)
	nodeStatus, _ := metrics.CovertGuage(nodeStatusMetric)
	nodeTimeGap, _ := metrics.CovertGuage(nodeTimeGapMetric)
	syncWorkway, _ := metrics.CovertGuage(syncWorkwayMetric)
	catchingTaskNum, _ := metrics.CovertGuage(syncCatchingTaskNumMetric)
	syncParseBug, _ := metrics.CovertGuage(syncParseBugMetric)
	return clientNode{
		nodeStatus:          nodeStatus,
		nodeHeight:          nodeHeight,
		dbHeight:            dbHeight,
		nodeTimeGap:         nodeTimeGap,
		syncWorkWay:         syncWorkway,
		syncCatchingTaskNum: catchingTaskNum,
		syncParseBugMetric:  syncParseBug,
	}
}

func (node *clientNode) Report() {
	for {
		t := time.NewTimer(time.Duration(10) * time.Second)
		select {
		case <-t.C:
			node.nodeStatusReport()
			node.syncCatchUpingReport()
			node.syncBugTxsReport()
		}
	}
}
func (node *clientNode) nodeStatusReport() {
	client, err := pool.GetClientWithTimeout(10 * time.Second)
	if err != nil {
		logger.Error("rpc node connection exception", logger.String("error", err.Error()))
		node.nodeStatus.Set(float64(NodeStatusNotReachable))
		return
	}
	defer func() {
		client.Release()
	}()

	block, err := new(models.Block).GetMaxBlockHeight()
	if err != nil && err != qmgo.ErrNoSuchDocuments {
		logger.Error("query block exception", logger.String("error", err.Error()))
	}
	node.dbHeight.Set(float64(block.Height))
	status, err := client.Status(context.Background())
	if err != nil {
		logger.Error("rpc node connection exception", logger.String("error", err.Error()))
		node.nodeStatus.Set(float64(NodeStatusNotReachable))
		//return
	} else {
		if status.SyncInfo.CatchingUp {
			node.nodeStatus.Set(float64(NodeStatusCatchingUp))
		} else {
			node.nodeStatus.Set(float64(NodeStatusSyncing))
		}
		node.nodeHeight.Set(float64(status.SyncInfo.LatestBlockHeight))
	}

	follow, err := new(models.SyncTask).QueryValidFollowTasks()
	if err != nil {
		logger.Error("query valid follow task exception", logger.String("error", err.Error()))
		return
	}
	if follow && block.Time > 0 {
		timeGap := time.Now().Unix() - block.Time
		node.nodeTimeGap.Set(float64(timeGap))
	}

	if follow {
		node.syncWorkWay.Set(float64(SyncTaskFollowing))
	} else {
		node.syncWorkWay.Set(float64(SyncTaskCatchingUp))
	}
	return
}

func (node *clientNode) syncCatchUpingReport() {
	catchUpTasksNum, err := new(models.SyncTask).QueryCatchUpingTasksNum()
	if err != nil {
		logger.Error("query task exception", logger.String("error", err.Error()))
	}
	node.syncCatchingTaskNum.Set(float64(catchUpTasksNum))
}

func (node *clientNode) syncBugTxsReport() {
	exist, err := new(models.Tx).FindExistBugTxs()
	if err != nil {
		logger.Error("query task exception", logger.String("error", err.Error()))
	}
	if exist {
		node.syncParseBugMetric.Set(float64(1))
	} else {
		node.syncParseBugMetric.Set(float64(0))
	}
}

func Start() {
	c := make(chan os.Signal)
	//monitor system signal
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// start monitor
	server := metrics.NewMonitor(models.GetSrvConf().PromethousPort)
	node := NewMetricNode(server)

	server.Report(func() {
		go node.Report()
	})
	<-c
}
