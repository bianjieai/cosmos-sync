module github.com/bianjieai/cosmos-sync

go 1.16

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/jolestar/go-commons-pool v2.0.0+incompatible
	github.com/kaifei-bianjie/common-parser v0.0.0-20220923023138-65dfc81a8ff5
	github.com/kaifei-bianjie/cosmosmod-parser v0.0.0-20220923075813-81e7e2b17c67
	github.com/kaifei-bianjie/cschain-mod-parser v0.0.0-20221220072642-c22c42ccc4bd
	github.com/kaifei-bianjie/irismod-parser v0.0.0-20220926014114-6f1e08c26bf0
	github.com/kaifei-bianjie/iritachain-mod-parser v0.0.0-20230317063603-82411421f5f2
	github.com/kaifei-bianjie/iritamod-parser v0.0.0-20220923075745-50a7044003db
	github.com/qiniu/qmgo v1.0.4
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.8.1
	github.com/tendermint/tendermint v0.35.0
	github.com/tharsis/ethermint v0.10.3
	github.com/weichang-bianjie/metric-sdk v1.0.0
	go.mongodb.org/mongo-driver v1.7.4
	go.uber.org/zap v1.19.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.44.2-irita-20211102
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.0-irita-210104.0.20210112015006-57e95aa6402f
	github.com/tharsis/ethermint => github.com/bianjieai/ethermint v0.10.2-irita-20230315
	golang.org/x/sys => golang.org/x/sys v0.0.0-20211210111614-af8b64212486
	google.golang.org/grpc => google.golang.org/grpc v1.35.0
)
