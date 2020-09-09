module github.com/bianjieai/irita-sync

go 1.14

require (
	github.com/cosmos/cosmos-sdk v0.38.2
	github.com/irismod/nft v1.2.0
	github.com/irismod/record v1.1.0
	github.com/irismod/service v1.1.0
	github.com/irismod/token v1.1.0
	github.com/jolestar/go-commons-pool v2.0.0+incompatible
	github.com/onsi/ginkgo v1.14.0 // indirect
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.6
	gitlab.bianjie.ai/cschain/cschain v1.0.1-0.20200826054044-98c73fc75924
	go.uber.org/zap v1.15.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637 // indirect
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.39.0-irita-200703
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.4-irita-200703
)
