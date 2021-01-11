package router

import (
	"github.com/gin-gonic/gin"
	"menu/router/method"
)

/*
	操作Linux信息路由配置
*/
func LinuxRouter(r *gin.Engine) {

	//获取服务器信息列表接口
	//路由分组
	v1Group := r.Group("/api/v1/common/resource")
	{
		//获取服务器性能列表
		v1Group.GET("/monitor", func(c *gin.Context) {
			method.LinuxGetMehod(c)
		})

		//服务器列表添加
		v1Group.POST("/monitor", func(c *gin.Context) {
			method.LinuxPostMethod(c)
		})

		//删除服务器列表数据
		v1Group.DELETE("/monitor", func(c *gin.Context) {
			method.LinuxDeleteMethod(c)
		})
	}
}
