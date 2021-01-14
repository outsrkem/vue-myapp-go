package routers

import (
	"github.com/gin-gonic/gin"
	"mana/src/method/linuxServe"
	"net/http"
)

// 路由总配置
func Index(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//Linux服务路由
	v1Group := r.Group("/api/v1/common/resource")
	{
		//获取服务器性能列表
		v1Group.GET("/monitor", linuxServe.LinuxGetMethod)

		//服务器列表添加
		//v1Group.POST("/monitor", middle.Admin, method.LinuxPostMethod)
		//
		////删除服务器列表数据
		//v1Group.DELETE("/monitor", middle.Admin, method.LinuxDeleteMethod)
	}
}
