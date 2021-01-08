module github.com/bianjieai/irita-sync

go 1.15

require (
    //github.com/CosmWasm/wasmd v0.13.1-0.20201217131318-53bbf96e9e87
	github.com/cosmos/cosmos-sdk v0.40.0-rc5
	github.com/irisnet/irismod v1.1.1-0.20201229063925-7d7dad20f951
	github.com/jolestar/go-commons-pool v2.0.0+incompatible
	github.com/tendermint/tendermint v0.34.0
	go.uber.org/zap v1.15.0
	golang.org/x/net v0.0.0-20201209123823-ac852fbbde11
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
