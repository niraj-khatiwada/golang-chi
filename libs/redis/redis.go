package redis

import (
	"github.com/redis/go-redis/v9"
	"go-web/config"
	"go-web/utils"
)

type Redis struct {
	Client *redis.Client
}

func (r *Redis) Init(conf config.Redis) error {
	var redisConf config.Redis
	if conf != (config.Redis{}) {
		redisConf = conf
	} else {
		redisConf = config.GetDefaultRedisConfig()
	}
	dsn := config.CreateRedisDSN(&redisConf)
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		utils.CatchRuntimeErrors(err)
		return err
	}
	client := redis.NewClient(opt)
	r.Client = client
	return nil
}
