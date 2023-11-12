package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-web/config"
	"go-web/utils"
	"time"
)

type Redis struct {
	Client          *redis.Client
	GlobalNamespace string
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, pingErr := client.Ping(ctx).Result()
	if pingErr != nil {
		utils.CatchRuntimeErrors(pingErr)
		return pingErr
	}
	r.GlobalNamespace = redisConf.GlobalNamespace
	r.Client = client
	return nil
}
