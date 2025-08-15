package inits

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/zhanghanchen1014/pubilc_pkg/config"
	"log"
)

func RedisInit() {
	config.Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.AppConf.Redis.Host, config.AppConf.Redis.Port),
		Password: config.AppConf.Redis.Password, // no password set
		DB:       config.AppConf.Redis.Database, // use default DB
	})

	err := config.Rdb.Set(config.Ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
		return
	}
	log.Println("Redis init success")
}
