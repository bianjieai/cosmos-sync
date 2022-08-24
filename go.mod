module github.com/bianjieai/cosmos-sync

go 1.16

require (
	github.com/jolestar/go-commons-pool v2.0.0+incompatible
	github.com/kaifei-bianjie/msg-parser v0.0.0-20220824054827-d4143c5e9b40
	github.com/qiniu/qmgo v1.0.4
	github.com/spf13/viper v1.9.0
	github.com/tendermint/tendermint v0.35.0
	github.com/weichang-bianjie/metric-sdk v1.0.0
	go.mongodb.org/mongo-driver v1.7.2
	go.uber.org/zap v1.17.0
	golang.org/x/net v0.0.0-20211111160137-58aab5ef257a
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.44.2-irita-20211102
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.8-irita-210413.0.20211012090339-cee6e09e8ae3
	github.com/tharsis/ethermint => github.com/bianjieai/ethermint v0.8.2-0.20220211020007-9ec25dde74d4
)
