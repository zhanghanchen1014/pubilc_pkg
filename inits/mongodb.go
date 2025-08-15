package inits

import (
	"context"
	"fmt"
	"github.com/zhanghanchen1014/pubilc_pkg/config"
	"go.mongodb.org/mongo-driver/v2/event"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

func MongoDBInit() *mongo.Collection {
	// 1. 准备命令监视器
	monitor := &event.CommandMonitor{
		Started: func(ctx context.Context, startedEvent *event.CommandStartedEvent) {
			// 输出查询命令
			fmt.Println(startedEvent.Command)
		},
		Succeeded: func(ctx context.Context, succeededEvent *event.CommandSucceededEvent) {},
		Failed:    func(ctx context.Context, errEvent *event.CommandFailedEvent) {},
	}
	// 2. 设置连接地址
	mongodb := config.AppConf.MongoDB

	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", mongodb.User, mongodb.Password, mongodb.Host, mongodb.Port)).SetMonitor(monitor)
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	log.Println("MongoDB init success")

	// 3. 创建database和collection
	return client.Database("goods").Collection("goods_demo")

}
