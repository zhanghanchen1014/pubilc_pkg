package config

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"sync"
)

var (
	AppConf AppConfig
	Once    sync.Once
	Ctx     = context.Background()
	DB      *gorm.DB
	Rdb     *redis.Client
	Els     *elasticsearch.Client
)
