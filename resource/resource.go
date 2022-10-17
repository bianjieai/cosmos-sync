package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/bianjieai/cosmos-sync/models"
	rpcclient "github.com/tendermint/tendermint/rpc/client/http"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	nodeUrlMap            = map[string]bool{}
	nodeEarliestHeightMap = map[string]int64{}
	mutex                 sync.RWMutex
)

func GetValidNodeUrl() (string, int64) {
	for nodeUri, val := range nodeUrlMap {
		if val {
			return nodeUri, nodeEarliestHeightMap[nodeUri]
		}
	}
	//logger.Debug("GetValidNodeUrl:", logger.Any("nodeUrlMap<nodeurl,valid>:", nodeUrlMap))
	return "", 0
}

func SetInvalidNode(nodeUri string) {
	if valid, ok := nodeUrlMap[nodeUri]; ok && valid {
		logger.Debug("set node rpc invalid", logger.String("node_rpc", nodeUri))
		mutex.Lock()
		nodeUrlMap[nodeUri] = false
		mutex.Unlock()
	}
}

func getData(chainId string) (string, error) {
	chainRegistry, err := new(models.ChainRegistry).FindOne(chainId)
	if err != nil {
		//logger.Error("loadRpcResource error: " + err.Error())
		return "", err
	}

	bz, err := HttpGet(chainRegistry.ChainJsonUrl)
	if err != nil {
		//logger.Error("rpc resource get chain json error: " + err.Error())
		return "", err
	}

	return string(bz), nil

}

func checkRpcValid(nodeUrl string, chainId string) error {
	client, err := rpcclient.New(nodeUrl, "/websocket")
	if err != nil {
		return err
	}
	defer client.Quit()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	retStatus, err := client.Status(ctx)
	if err != nil {
		return err
	}
	//node if catching up
	if retStatus.SyncInfo.CatchingUp {
		return fmt.Errorf("node is catchingUp")
	}

	//node check tx_index
	if strings.Compare(strings.ToLower(retStatus.NodeInfo.Other.TxIndex), "off") == 0 {
		return fmt.Errorf("transaction indexing is disabled")
	}

	//network no match
	network := strings.ReplaceAll(retStatus.NodeInfo.Network, "-", "_")
	if network != chainId {
		return fmt.Errorf("network(%s) not match config network:%s", network, chainId)
	}
	mutex.Lock()
	if _, ok := nodeEarliestHeightMap[nodeUrl]; !ok {
		nodeEarliestHeightMap[nodeUrl] = retStatus.SyncInfo.EarliestBlockHeight
	}
	mutex.Unlock()
	return nil
}

func LoadRpcResource(bz string, chainId string) (string, error) {
	var chainRegisterResp ChainRegisterResp
	err := json.Unmarshal([]byte(bz), &chainRegisterResp)
	if err != nil {
		//logger.Error("rpc resource get chain json error: " + err.Error())
		return "", err
	}
	var rpcAddrs []string
	for _, v := range chainRegisterResp.Apis.Rpc {
		nodeUrl := HandleUri(v.Address)
		if err := checkRpcValid(nodeUrl, chainId); err == nil {
			rpcAddrs = append(rpcAddrs, nodeUrl)
		} else {
			logger.Debug("invalid nodeurl:"+nodeUrl, logger.String("err", err.Error()))
		}
	}
	nodeEarliestHeightMap = make(map[string]int64, len(rpcAddrs))

	return strings.Join(rpcAddrs, ","), nil
}

func HandleUri(rpcaddr string) string {
	var nodeUrl string
	nodehttp := strings.Split(rpcaddr, "://")[0]
	nodeuri := strings.Split(rpcaddr, "://")[1]
	if strings.Count(nodeuri, "/") <= 1 {
		if strings.Contains(nodeuri, ":") {
			nodeuri = strings.ReplaceAll(nodeuri, "/", "")
		} else {
			if strings.Contains(nodeuri, "/") {
				nodeuri = strings.ReplaceAll(nodeuri, "/", ":443")
			} else {
				nodeuri = nodeuri + ":443"
			}
		}
		nodeUrl = nodehttp + "://" + nodeuri
	} else {
		nodeuri = strings.Replace(nodeuri, "/", ":443/", 1)
		nodeUrl = nodehttp + "://" + nodeuri
	}
	return nodeUrl
}

func ReloadRpcResourceMap(rpcAddrs []string) {
	mutex.Lock()
	nodeUrlMap = make(map[string]bool, len(rpcAddrs))
	for _, val := range rpcAddrs {
		nodeUrlMap[val] = true
	}
	mutex.Unlock()
}

func HttpGet(url string) (bz []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode(%s) != 200, url: %s", resp.Status, url)
	}

	bz, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return
}

func GetRpcNodesFromGithubRepo(chainId string) (string, error) {
	data, err := getData(chainId)
	if err != nil {
		return "", fmt.Errorf("%v %s", err, "GetRpcNodesFromGithubRepo fail")
	}
	nodeurl, err := LoadRpcResource(data, chainId)
	if err != nil {
		return "", fmt.Errorf("%v %s", err, "LoadRpcResource fail")
	}
	logger.Info("valid Rpc Nodes From Github Repo: " + nodeurl)
	return nodeurl, nil
}
