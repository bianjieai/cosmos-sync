module github.com/bianjieai/cosmos-sync

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.45.5-0.20220523154235-2921a1c3c918
	github.com/go-redis/redis/v8 v8.11.5
	github.com/jolestar/go-commons-pool v2.0.0+incompatible
	github.com/kaifei-bianjie/common-parser v0.0.0-20220923023138-65dfc81a8ff5
	github.com/kaifei-bianjie/cosmosmod-parser v0.0.0-20220923075813-81e7e2b17c67
	github.com/kaifei-bianjie/irismod-parser v0.0.0-20220926014114-6f1e08c26bf0
	github.com/kaifei-bianjie/iritachain-mod-parser v0.0.0-20230317063603-82411421f5f2
	github.com/kaifei-bianjie/iritamod-parser v0.0.0-20220923075745-50a7044003db
	github.com/kaifei-bianjie/spartanchain-mod-parser v0.0.0-20220926103514-5f0a9d27d019
	github.com/kaifei-bianjie/tibc-mod-parser v0.0.0-20220927054648-1e71af64d126
	github.com/qiniu/qmgo v1.0.4
	github.com/spf13/viper v1.12.0
	github.com/tendermint/tendermint v0.35.0
	github.com/tharsis/ethermint v0.10.3
	github.com/weichang-bianjie/metric-sdk v1.0.0
	go.mongodb.org/mongo-driver v1.7.2
	go.uber.org/zap v1.19.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.45.1-irita-20220816.0.20220816095307-845547d9c19e
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.8-irita-210413.0.20210908054213-781a5fed16d6
	github.com/tharsis/ethermint => github.com/bianjieai/ethermint v0.10.2-irita-20230315
	golang.org/x/sys => golang.org/x/sys v0.0.0-20211210111614-af8b64212486
)
