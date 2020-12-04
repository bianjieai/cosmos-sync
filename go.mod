module github.com/bianjieai/irita-sync

go 1.14

require (
	github.com/cosmos/cosmos-sdk v0.40.0-rc3
	github.com/irisnet/irismod v1.1.1-0.20201126013702-4999558204d6
	github.com/jolestar/go-commons-pool v2.0.0+incompatible
	github.com/tendermint/tendermint v0.34.0-rc6
	gitlab.bianjie.ai/cschain/cschain v1.1.1-0.20201203013652-8cce6b01a507
	github.com/bianjieai/iritamod v0.0.0-20201202112849-cdce8a8df2d0
	go.uber.org/zap v1.15.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637 // indirect
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.34.4-0.20201127022001-791921d241f8
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.1-dev0.0.20201126055325-2217bc51b6c7
)
