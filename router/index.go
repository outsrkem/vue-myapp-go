package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	路由总配置
*/
func Index(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "没有请求参数")
	})
}
