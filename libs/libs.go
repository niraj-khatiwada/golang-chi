package libs

import (
	"go-web/libs/db"
	"go-web/libs/redis"
)

type Libs struct {
	DB    *db.DB
	Redis *redis.Redis
}
