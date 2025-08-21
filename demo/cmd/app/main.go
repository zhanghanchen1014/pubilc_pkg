package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanghanchen1014/pubilc_pkg/demo/api"
)

func main() {
	router := gin.Default()
	api.LoadRouter(router)
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
