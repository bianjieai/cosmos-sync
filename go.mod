module github.com/bianjieai/cosmos-sync

go 1.16

require (
	github.com/bianjieai/iritamod v1.2.1-0.20220222035322-99168809cf24
	github.com/cosmos/cosmos-sdk v0.45.1
	github.com/gin-gonic/gin v1.7.7
	github.com/go-redis/redis/v8 v8.11.5
	github.com/jolestar/go-commons-pool v2.0.0+incompatible
	github.com/prometheus/client_golang v1.12.1
	github.com/qiniu/qmgo v1.0.4
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.8.1
	github.com/tendermint/tendermint v0.35.0
	github.com/tharsis/ethermint v0.10.3
	gitlab.cschain.tech/cschain v0.0.0-00010101000000-000000000000
	gitlab.cschain.tech/csmod v1.5.2
	go.mongodb.org/mongo-driver v1.7.4
	go.uber.org/zap v1.19.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.44.2-irita-20211102
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.0-irita-210104.0.20210112015006-57e95aa6402f
	github.com/tharsis/ethermint => github.com/bianjieai/ethermint v0.10.2-irita-20230315
	gitlab.cschain.tech/cschain => gitlab.bianjie.ai/cschain/cschain v1.1.1-0.20230920073020-5050d8bbe222
	gitlab.cschain.tech/csmod => gitlab.bianjie.ai/cschain-otcibr/csmod v0.0.0-20230914103407-ead5d33af660
	gitlab.cschain.tech/ibc-go => gitlab.bianjie.ai/cschain-otcibr/ibc-go v0.0.0-20230920071741-003ddf2d392f
	golang.org/x/sys => golang.org/x/sys v0.0.0-20211210111614-af8b64212486
	google.golang.org/grpc => google.golang.org/grpc v1.35.0
)
