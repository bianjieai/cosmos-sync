//init client from clientPool.
//client is httpClient of tendermint

package pool

import (
	"context"
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/resource"
	commonPool "github.com/jolestar/go-commons-pool"
	"golang.org/x/time/rate"
	"strings"
	"sync"
	"time"
)

var (
	poolObject  *commonPool.ObjectPool
	poolFactory PoolFactory
	ctx         = context.Background()
)

func Init(conf *config.Config) {
	var (
		syncMap  sync.Map
		rpcAddrs string
	)
	if conf.Server.UseNodeUrls {
		rpcAddrs = conf.Server.NodeUrls
	} else {
		nodeRpcs, err := resource.GetRpcNodesFromGithubRepo(conf.Server.Chain, conf.Server.ChainId)
		if err != nil {
			//exist when get rcp node from github repo fail
			logger.Fatal("GetRpcNodesFromGithubRepo fail " + err.Error())
			return
		}
		if len(nodeRpcs) == 0 {
			//exist when no found rpc node
			logger.Fatal("no found Rpc Nodes From GithubRepo")
		}
		rpcAddrs = nodeRpcs
	}
	nodeUrls := strings.Split(rpcAddrs, ",")
	resource.ReloadRpcResourceMap(nodeUrls)
	for _, url := range nodeUrls {
		key := generateId(url)
		endPoint := EndPoint{
			Address:   url,
			Available: true,
		}

		syncMap.Store(key, endPoint)
	}

	poolFactory = PoolFactory{
		chainId:    conf.Server.ChainId,
		chain:      conf.Server.Chain,
		peersMap:   syncMap,
		local:      conf.Server.UseNodeUrls,
		retryLimit: rate.NewLimiter(rate.Every(3*time.Minute), 1),
	}

	config := commonPool.NewDefaultPoolConfig()
	config.MaxTotal = conf.Server.MaxConnectionNum
	config.MaxIdle = conf.Server.InitConnectionNum
	config.MinIdle = conf.Server.InitConnectionNum
	//config.TestOnBorrow = true
	config.TestOnCreate = true
	config.TestWhileIdle = true
	config.LIFO = false

	poolObject = commonPool.NewObjectPool(ctx, &poolFactory, config)
	poolObject.PreparePool(ctx)
}

// get client from pool
func GetClient() *Client {
	c, err := poolObject.BorrowObject(ctx)
	for err != nil {
		logger.Error("GetClient failed,will try again after 3 seconds", logger.String("err", err.Error()))
		time.Sleep(3 * time.Second)
		c, err = poolObject.BorrowObject(ctx)
	}

	return c.(*Client)
}

// release client
func (c *Client) Release() {
	err := poolObject.ReturnObject(ctx, c)
	if err != nil {
		logger.Error(err.Error())
	}
}

func (c *Client) HeartBeat() error {
	http := c.HTTP
	_, err := http.Health()
	return err
}

func (c *Client) InvalidateObject() {
	err := poolObject.InvalidateObject(ctx, c)
	if err != nil {
		logger.Error(err.Error())
	}
}

func ClosePool() {
	poolObject.Close(ctx)
}
