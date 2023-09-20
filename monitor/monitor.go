package monitor

import (
	"context"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/qiniu/qmgo"
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
	nodeStatus  prometheus.Gauge
	nodeTimeGap prometheus.Gauge
	syncWorkWay prometheus.Gauge
}

func NewMetricNode() clientNode {

	nodeStatusMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "sync",
		Subsystem: "status",
		Name:      "node_status",
		Help:      "full node status(0:NotReachable,1:Syncing,2:CatchingUp)",
	})
	nodeTimeGapMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "sync",
		Subsystem: "status",
		Name:      "node_seconds_gap",
		Help:      "the seconds gap between node block time with sync db block time",
	})
	syncWorkwayMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "sync",
		Subsystem: "",
		Name:      "task_working_status",
		Help:      "sync task working status(0:CatchingUp 1:Following)",
	})

	prometheus.MustRegister(nodeStatusMetric)
	prometheus.MustRegister(nodeTimeGapMetric)
	prometheus.MustRegister(syncWorkwayMetric)
	return clientNode{
		nodeStatus:  nodeStatusMetric,
		nodeTimeGap: nodeTimeGapMetric,
		syncWorkWay: syncWorkwayMetric,
	}
}

func (node *clientNode) Report() {
	for {
		t := time.NewTimer(time.Duration(10) * time.Second)
		select {
		case <-t.C:
			node.nodeStatusReport()
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

func Start() {
	// start monitor
	node := NewMetricNode()
	node.Report()
}
