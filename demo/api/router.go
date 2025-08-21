package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanghanchen1014/pubilc_pkg/demo/api/handler"
)

func LoadRouter(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{
		upload := apiGroup.Group("/upload")
		{
			upload.POST("/oss", handler.AliyunUpload)  //阿里云
			upload.POST("/minio", handler.MinioUpload) //minio
		}
	}
}
