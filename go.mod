module github.com/bianjieai/irita-sync

go 1.14

require (
	github.com/cosmos/cosmos-sdk v0.38.1
	github.com/irisnet/irismod v1.1.0
	github.com/jolestar/go-commons-pool v2.0.0+incompatible
	github.com/tendermint/tendermint v0.34.0-rc3.0.20200907055413-3359e0bf2f84
	gitlab.bianjie.ai/cschain/cschain v1.1.1-0.20201015031913-928bf6f994b0
	gitlab.bianjie.ai/irita-pro/iritamod v1.1.0
	go.uber.org/zap v1.15.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637 // indirect
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.40.0-irita-200930
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.0-irita-200930
)
