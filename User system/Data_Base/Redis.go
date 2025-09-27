package Data_Base

import (
	"User_system/Error_Handling"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ClientRedis *redis.Client

func RunRedisServer() string {
	Error_Handling.InitLogger(zapcore.InfoLevel)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ClientRedis = redis.NewClient(&redis.Options{
		Addr:     Error_Handling.LoadConfigFile().Redis.Addr,
		Password: Error_Handling.LoadConfigFile().Redis.Password,
		DB:       Error_Handling.LoadConfigFile().Redis.DB,
	})
	ping, err := ClientRedis.Ping(ctx).Result()
	if err != nil {
		zap.L().Error("Redis could not be connected", zap.String("err", err.Error()))
		return "error ping"
	}
	zap.L().Warn("Redis connected", zap.String("ping", ping))
	zap.L().Info("RedisServer is running")
	return "ok"

}
