module github.com/bianjieai/cosmos-sync

go 1.16

require (
	github.com/jolestar/go-commons-pool v2.0.0+incompatible
	github.com/kaifei-bianjie/msg-parser v0.0.0-20220916031544-0074da0488e4
	github.com/qiniu/qmgo v1.0.4
	github.com/spf13/viper v1.12.0
	github.com/tendermint/tendermint v0.35.0
	github.com/tharsis/ethermint v0.10.3
	github.com/weichang-bianjie/metric-sdk v1.0.0
	go.mongodb.org/mongo-driver v1.7.2
	go.uber.org/zap v1.19.1
	golang.org/x/net v0.0.0-20220520000938-2e3eb7b945c2
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.45.1-irita-20220816.0.20220816095307-845547d9c19e
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.8-irita-210413.0.20210908054213-781a5fed16d6
	github.com/tharsis/ethermint => github.com/bianjieai/ethermint v0.10.2-irita-20220826
)
