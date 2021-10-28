package pool

import (
	"context"
	"fmt"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	commonPool "github.com/jolestar/go-commons-pool"
	rpcclient "github.com/tendermint/tendermint/rpc/client/http"
	"math/rand"
	"sync"
)

type (
	PoolFactory struct {
		peersMap sync.Map
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
	c, err := newClient(endpoint.Address)
	if err != nil {
		return nil, err
	} else {
		return commonPool.NewPooledObject(c), nil
	}
}

func (f *PoolFactory) DestroyObject(ctx context.Context, object *commonPool.PooledObject) error {
	c := object.Object.(*Client)
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
