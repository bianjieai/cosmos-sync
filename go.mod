module github.com/bianjieai/cosmos-sync

go 1.16

require (
	github.com/jolestar/go-commons-pool v2.0.0+incompatible
	github.com/kaifei-bianjie/msg-parser v0.0.0-20211022081833-96bb9d9bea64
	github.com/qiniu/qmgo v1.0.4
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.34.0-rc6
	github.com/weichang-bianjie/metric-sdk v1.0.0
	go.mongodb.org/mongo-driver v1.7.4
	go.uber.org/zap v1.15.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.34.4-0.20201127022001-791921d241f8
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.1-dev0.0.20201126055325-2217bc51b6c7
	google.golang.org/grpc => google.golang.org/grpc v1.35.0
)
