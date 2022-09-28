module github.com/bianjieai/cosmos-sync

go 1.15

require (
	github.com/jolestar/go-commons-pool v2.0.0+incompatible
	github.com/kaifei-bianjie/common-parser v0.0.0-20220923023138-65dfc81a8ff5
	github.com/kaifei-bianjie/cosmosmod-parser v0.0.0-20220927055038-b4022b663f59
	github.com/kaifei-bianjie/irismod-parser v0.0.0-20220926014114-6f1e08c26bf0
	github.com/kaifei-bianjie/tibc-mod-parser v0.0.0-20220927054648-1e71af64d126
	github.com/qiniu/qmgo v1.0.4
	github.com/spf13/viper v1.10.1
	github.com/tendermint/tendermint v0.35.0
	github.com/weichang-bianjie/metric-sdk v1.0.0
	go.mongodb.org/mongo-driver v1.7.2
	go.uber.org/zap v1.19.1
	golang.org/x/net v0.0.0-20211208012354-db4efeb81f4b
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.45.1-irita-20220816.0.20220816095307-845547d9c19e
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.8-irita-210413.0.20211012090339-cee6e09e8ae3
)
