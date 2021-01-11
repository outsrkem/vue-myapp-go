package router

import (
	"github.com/gin-gonic/gin"
	"menu/router/method"
)

func UserRouter(r *gin.Engine) {
	v3Group := r.Group("/api/v1/common/user")
	{
		v3Group.GET("/table", func(c *gin.Context) {
			method.UserGetMethod(c)
		})

		v3Group.POST("/table", func(c *gin.Context) {
			method.UserPostMethod(c)
		})

		v3Group.DELETE("/table", func(c *gin.Context) {
			method.UserDeleteMethod(c)
		})
	}
}
