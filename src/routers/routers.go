package routers

import (
	"github.com/gin-gonic/gin"
	"mana/src/controllers/user"
	"mana/src/filters/util"
	"net/http"
)

// 路由总配置
func Index(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"message": "successfully",
		})
	})

	// 用户登录
	r.POST("/api/v1/common/user/login", user.Login)
	// 用户注册
	r.POST("/api/v1/common/user/register", user.InstUser)

	// 验证请求token中间件
	r.Use(util.AuthToken())

	// Linux服务路由
	v1Group := r.Group("/api/v1")
	{

		v1Group.GET("/common/user/userinfo/:uid", user.FindByUserinfo)

		// 添加k8s配置文件
		//v1Group.POST("/cluster", method.K8sPostMethod)
	}
}
