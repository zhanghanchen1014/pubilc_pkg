package inits

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"github.com/zhanghanchen1014/viper/config"
)

func EsInit() {
	var err error

	config.Once.Do(func() {
		cfg := elasticsearch.Config{
			Addresses: []string{
				"http://" + fmt.Sprintf("%s:%d", config.AppConf.Elasticsearch.Host, config.AppConf.Elasticsearch.Port),
			},
			// ...
		}
		config.Els, err = elasticsearch.NewClient(cfg)
		if err != nil {
			panic(err)
			return
		}
		log.Println("ES init success")
	})

}
