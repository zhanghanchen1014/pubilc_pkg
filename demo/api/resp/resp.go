package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  msg,
		"data": data,
	})
}
