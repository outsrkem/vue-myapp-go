package router

import (
	"github.com/gin-gonic/gin"
	"menu/middle"
	"menu/router/method"
)

/*
	路由总配置
*/
func Index(r *gin.Engine) {

	//登录请求
	r.GET("/login", method.LoginMethod)

	//验证请求token中间件
	r.Use(middle.AuthToken())

	//Linux服务路由
	v1Group := r.Group("/api/v1/common/resource")
	{
		//获取服务器性能列表
		v1Group.GET("/monitor", middle.General, method.LinuxGetMehod)

		//服务器列表添加
		v1Group.POST("/monitor", middle.Admin, method.LinuxPostMethod)

		//删除服务器列表数据
		v1Group.DELETE("/monitor", middle.Admin, method.LinuxDeleteMethod)
	}

	//k8s服务路由
	v2Group := r.Group("/api/v1/common/kubernetes")
	{
		//添加k8s配置文件
		v2Group.POST("/cluster", middle.General, method.K8sPostMethod)

		//查询k8s集群所有配置信息
		v2Group.GET("/cluster", middle.Admin, method.K8sGetMethod)

		//删除k8s集群配置信息
		v2Group.DELETE("/cluster", middle.Admin, method.K8sDeleteMethod)
	}

	//登录用户服务路由
	v3Group := r.Group("/api/v1/common/user")
	{
		//查询数据库用户表
		v3Group.GET("/table", middle.General, method.UserGetMethod)

		//添加用户信息
		v3Group.POST("/table", middle.Admin, method.UserPostMethod)

		//删除用户信息
		v3Group.DELETE("/table", middle.Admin, method.UserDeleteMethod)
	}
}
