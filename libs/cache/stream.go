package cache

import (
	"context"
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/libs/logger"
	"github.com/go-redis/redis/v8"
)

var (
	_redisStreamMQClient *RedisStreamMQClient
)

func InitMQClient(conf *config.Config) {
	_redisStreamMQClient = &RedisStreamMQClient{
		Client:         rdb,
		StreamKey:      conf.Redis.StreamKey,
		StreamMqMaxLen: conf.Redis.StreamMqMaxLen,
	}
}

func GetClient() *RedisStreamMQClient {
	return _redisStreamMQClient
}

type RedisStreamMQClient struct {
	Client         *redis.Client
	StreamKey      string //stream对应的key值
	StreamMqMaxLen int64  //stream的最大长度
}

// PutMsg 添加消息
func (mqClient *RedisStreamMQClient) PutMsg(streamKey string, msg map[string]interface{}) (string, error) {
	conn := mqClient.Client
	//*表示由Redis自己生成消息ID，设置MAXLEN可以保证消息队列的长度不会一直累加
	strMsgId, err := conn.XAdd(context.Background(), &redis.XAddArgs{
		Stream: streamKey,
		MaxLen: mqClient.StreamMqMaxLen,
		ID:     "*",
		Values: msg,
	}).Result()
	if err != nil {
		logger.Error("XADD failed", logger.String("err", err.Error()))
		return "", err
	}

	return strMsgId, nil
}

// PutMsg 批量添加消息
func (mqClient *RedisStreamMQClient) PutMsgBatch(streamKey string, evmTxHashs []string) error {
	conn := mqClient.Client
	//*表示由Redis自己生成消息ID，设置MAXLEN可以保证消息队列的长度不会一直累加
	_, err := conn.Pipelined(context.Background(), func(pipe redis.Pipeliner) error {
		for _, evmTxHash := range evmTxHashs {
			xAddArgs := redis.XAddArgs{
				Stream: streamKey,
				MaxLen: mqClient.StreamMqMaxLen,
				ID:     "*",
				Values: evmTxHash,
			}

			_, err := pipe.XAdd(context.Background(), &xAddArgs).Result()
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		logger.Error("XADD batch failed", logger.String("err", err.Error()))
		return err
	}

	return nil
}

// GetStreamLen 获取消息队列的长度
func (mqClient *RedisStreamMQClient) GetStreamLen(streamKey string) (int64, error) {
	conn := mqClient.Client

	reply, err := conn.XLen(context.Background(), streamKey).Result()
	if err != nil {
		return -1, err
	}
	return reply, nil
}
