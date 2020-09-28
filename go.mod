module github.com/bianjieai/irita-sync

go 1.14

require (
	github.com/cosmos/cosmos-sdk v0.38.2
	github.com/irisnet/irismod v0.0.0-20200927063845-2c0b5f1b4ca6
	github.com/jolestar/go-commons-pool v2.0.0+incompatible
	github.com/tendermint/tendermint v0.34.0-rc3.0.20200907055413-3359e0bf2f84
	gitlab.bianjie.ai/irita-pro/iritamod v0.0.0-20200925015143-a512f1a33cf2
	go.uber.org/zap v1.15.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637 // indirect
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.34.4-0.20200920153336-6dd96d838b0f
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.4-irita-200703.0.20200925112439-d4196a88a285
)
