package config

import (
	"fmt"
	"go-web/utils"
	"os"
)

type Redis struct {
	Host            string `env:"REDIS_HOST"`
	Port            string `env:"REDIS_PORT"`
	Username        string `env:"REDIS_USERNAME"`
	Password        string `env:"REDIS_PASSWORD"`
	DB              string `env:"REDIS_DB_INDEX"`
	GlobalNamespace string `env:"REDIS_GLOBAL_NAMESPACE"`
}

func GetDefaultRedisConfig() Redis {
	utils.LoadEnv()
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	username := os.Getenv("REDIS_USERNAME")
	password := os.Getenv("REDIS_PASSWORD")
	dbIndex := os.Getenv("REDIS_DB_INDEX")
	globalNamespace := os.Getenv("REDIS_GLOBAL_NAMESPACE")
	return Redis{Host: host, Port: port, Username: username, Password: password, DB: dbIndex, GlobalNamespace: globalNamespace}
}

func CreateRedisDSN(config *Redis) string {
	//redis://<user>:<pass>@localhost:6379/<db>
	return fmt.Sprintf("redis://%s:%s@%s:%s/%s", config.Username, config.Password, config.Host, config.Port, config.DB)
}
