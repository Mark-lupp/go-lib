package db

import (
	"fmt"
	"time"

	"github.com/Mark-lupp/go-lib/lib/config"
	"github.com/Mark-lupp/go-lib/log"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func initRedis() {

	redisDb = redis.NewClient(
		&redis.Options{
			Addr:         fmt.Sprintf("%s:%s", config.GetRedisConfig().GetIP(), config.GetRedisConfig().GetPort()),
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			Password:     config.GetRedisConfig().GetPass(),
			PoolSize:     config.GetRedisConfig().GetMaxOpen(),
		},
	)
	err = redisDb.Ping().Err()
	if nil != err {
		log.NewLogger().Error("ping redis err:", zap.Error(err))
		panic(err)
	}
	log.NewLogger().Debug("redis success : " + fmt.Sprintf("%s:%s", config.GetRedisConfig().GetIP(), config.GetRedisConfig().GetPort()))

}

func closeRedis() {
	if redisDb != nil {
		_ = redisDb.Close()
	}
}
