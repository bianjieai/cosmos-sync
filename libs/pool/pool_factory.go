package pool

import (
	"context"
	"fmt"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/resource"
	commonPool "github.com/jolestar/go-commons-pool"
	rpcclient "github.com/tendermint/tendermint/rpc/client/http"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type (
	PoolFactory struct {
		chainId     string
		local       bool
		retryMax    int
		autoRpcFunc *time.Timer
		peersMap    sync.Map
	}
	EndPoint struct {
		Address   string
		Available bool
	}
	Client struct {
		Id string
		*rpcclient.HTTP
	}
)

func (f *PoolFactory) MakeObject(ctx context.Context) (*commonPool.PooledObject, error) {
	endpoint := f.GetEndPoint()
	if endpoint.Available {
		logger.Info("current use node rpc info", logger.String("node_rpc", endpoint.Address))
		c, err := newClient(endpoint.Address)
		if err != nil {
			return nil, err
		} else {
			return commonPool.NewPooledObject(c), nil
		}
	} else {
		if f.local {
			return nil, fmt.Errorf("no found valid node")
		}
		//get valid nodeurl
		address, _ := resource.GetValidNodeUrl()
		var closeTimer bool
		if len(address) == 0 {
			//if no found valid node, auto update rpc nodes from githubRepo
			if f.autoRpcFunc == nil {
				f.autoRpcFunc = time.AfterFunc(time.Duration(1)*time.Minute, func() {
					f.retryMax++
					logger.Info("auto update rpc nodes from githubRepo", logger.String("chainId", f.chainId))
					nodeRpcs, err := resource.GetRpcNodesFromGithubRepo(f.chainId)
					if err != nil {
						logger.Error(err.Error())
						return
					}
					if len(nodeRpcs) > 0 {
						closeTimer = true
						nodeUrls := strings.Split(nodeRpcs, ",")
						resource.ReloadRpcResourceMap(nodeUrls)
						for _, url := range nodeUrls {
							key := generateId(url)
							endPoint := EndPoint{
								Address:   url,
								Available: true,
							}
							f.peersMap.Store(key, endPoint)
						}
					}
				})
			}
			if f.retryMax > 5 || closeTimer {
				f.StopAutoRpcTimer()
				time.Sleep(5 * time.Minute)
			}
			return nil, fmt.Errorf("no found valid node")
		} else {
			key := generateId(address)
			endPoint := EndPoint{
				Address:   address,
				Available: true,
			}
			f.peersMap.Store(key, endPoint)
			c, err := newClient(address)
			if err != nil {
				return nil, err
			} else {
				return commonPool.NewPooledObject(c), nil
			}
		}
	}
}

func (f *PoolFactory) StopAutoRpcTimer() {
	f.retryMax = 0
	if f.autoRpcFunc != nil {
		f.autoRpcFunc.Stop()
		f.autoRpcFunc = nil
	}
}

func (f *PoolFactory) DestroyObject(ctx context.Context, object *commonPool.PooledObject) error {
	c := object.Object.(*Client)
	value, ok := f.peersMap.Load(c.Id)
	//set endpoint invalid
	if ok {
		endPoint := value.(EndPoint)
		endPoint.Available = false
		f.peersMap.Store(c.Id, endPoint)
		resource.SetInvalidNode(endPoint.Address)
	}
	if c.IsRunning() {
		c.Stop()
	}
	return nil
}

func (f *PoolFactory) ValidateObject(ctx context.Context, object *commonPool.PooledObject) bool {
	// do validate
	c := object.Object.(*Client)
	if c.HeartBeat() != nil {
		value, ok := f.peersMap.Load(c.Id)
		if ok {
			endPoint := value.(EndPoint)
			endPoint.Available = true
			f.peersMap.Store(c.Id, endPoint)
		}
		return false
	}
	stat, err := c.Status(ctx)
	if err != nil {
		return false
	}
	if stat.SyncInfo.CatchingUp {
		value, ok := f.peersMap.Load(c.Id)
		if ok {
			endPoint := value.(EndPoint)
			resource.SetInvalidNode(endPoint.Address)
		}
		return false
	}
	return true
}

func (f *PoolFactory) ActivateObject(ctx context.Context, object *commonPool.PooledObject) error {
	return nil
}

func (f *PoolFactory) PassivateObject(ctx context.Context, object *commonPool.PooledObject) error {
	return nil
}

func (f *PoolFactory) GetEndPoint() EndPoint {
	var (
		keys        []string
		selectedKey string
	)

	f.peersMap.Range(func(k, value interface{}) bool {
		key := k.(string)
		endPoint := value.(EndPoint)
		if endPoint.Available {
			keys = append(keys, key)
		}
		selectedKey = key

		return true
	})

	if len(keys) > 0 {
		index := rand.Intn(len(keys))
		selectedKey = keys[index]
	}
	value, ok := f.peersMap.Load(selectedKey)
	if ok {
		return value.(EndPoint)
	} else {
		logger.Error("Can't get selected end point", logger.String("selectedKey", selectedKey))
	}
	return EndPoint{}
}

func newClient(nodeUrl string) (*Client, error) {
	client, err := rpcclient.New(nodeUrl, "/websocket")
	return &Client{
		Id:   generateId(nodeUrl),
		HTTP: client,
	}, err
}

func generateId(address string) string {
	return fmt.Sprintf("peer[%s]", address)
}
func (f *PoolFactory) PoolValidNodes() []string {
	var (
		nodes []string
	)

	f.peersMap.Range(func(k, value interface{}) bool {
		endPoint := value.(EndPoint)
		if endPoint.Available {
			nodes = append(nodes, endPoint.Address)
		}
		return true
	})
	return nodes
}

func PoolValidNodes() []string {
	return poolFactory.PoolValidNodes()
}
