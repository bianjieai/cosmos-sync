module github.com/bianjieai/irita-sync

go 1.14

require (
	github.com/cosmos/cosmos-sdk v0.34.4-0.20200914022129-c26ef79ed0a2
	github.com/irisnet/irismod v1.1.1-0.20200930082942-4882cf4fd17a
	github.com/jolestar/go-commons-pool v2.0.0+incompatible
	github.com/tendermint/tendermint v0.34.0-rc4
	go.uber.org/zap v1.15.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637 // indirect
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/irisnet/cosmos-sdk v0.28.2-0.20200930073142-1e1bec2e65a6
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
)
