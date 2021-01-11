package router

import (
	"github.com/gin-gonic/gin"
	"menu/router/method"
)

/*
	路由总配置
*/
func Index(r *gin.Engine) {

	r.GET("/", method.IndexMethod)

	//Linux服务路由
	v1Group := r.Group("/api/v1/common/resource")
	{
		//获取服务器性能列表
		v1Group.GET("/monitor", method.LinuxGetMehod)

		//服务器列表添加
		v1Group.POST("/monitor", method.LinuxPostMethod)

		//删除服务器列表数据
		v1Group.DELETE("/monitor", method.LinuxDeleteMethod)
	}

	//k8s服务路由
	v2Group := r.Group("/api/v1/common/kubernetes")
	{
		//添加k8s配置文件
		v2Group.POST("/cluster", method.K8sPostMethod)

		//查询k8s集群所有配置信息
		v2Group.GET("/cluster", method.K8sGetMethod)

		//删除k8s集群配置信息
		v2Group.DELETE("/cluster", method.K8sDeleteMethod)
	}

	//登录用户服务路由
	v3Group := r.Group("/api/v1/common/user")
	{
		v3Group.GET("/table", method.UserGetMethod)

		v3Group.POST("/table", method.UserPostMethod)

		v3Group.DELETE("/table", method.UserDeleteMethod)
	}
}
