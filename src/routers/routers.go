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
	r.POST("/api/v1/common/login", user.Login)
	// 用户注册
	r.POST("/api/v1/common/register", user.InstUser)

	// 验证请求token中间件
	r.Use(util.AuthToken())

	// Linux服务路由
	v1Group := r.Group("/api/v1")
	{

		v1Group.GET("/common/usercenter", user.FindByUserinfo)
		//获取服务器性能列表
		//v1Group.GET("/common/resource/monitor", linuxServe.LinuxGetMethod)

		//服务器列表添加
		//v1Group.POST("/monitor", middle.Admin, method.LinuxPostMethod)
		//
		////删除服务器列表数据
		//v1Group.DELETE("/monitor", middle.Admin, method.LinuxDeleteMethod)
	}
}
