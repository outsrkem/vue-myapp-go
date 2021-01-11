package router

import (
	"github.com/gin-gonic/gin"
	"menu/router/method"
)

/*
	路由总配置
*/
func Index(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		method.IndexMethod(c)
	})
}
