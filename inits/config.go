package inits

import (
	"github.com/spf13/viper"
	"log"
	"viper/config"
)

func ConfigInit() {
	viper.SetConfigFile("D:\\gowork\\src\\viper\\dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
		return
	}
	err = viper.Unmarshal(&config.AppConf)
	if err != nil {
		panic(err)
		return
	}
	log.Println("动态配置读取成功", config.AppConf)
}
