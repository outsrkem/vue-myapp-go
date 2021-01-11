package router

import (
	"github.com/gin-gonic/gin"
	"menu/router/method"
)

/*
	操作k8s信息路由总配置
*/
func K8sRouter(r *gin.Engine) {

	//获取k8s集群信息列表接口
	v2Group := r.Group("/api/v1/common/kubernetes")
	{
		//添加k8s配置文件
		v2Group.POST("/cluster", func(c *gin.Context) {
			method.K8sPostMethod(c)
		})

		//查询k8s集群所有配置信息
		v2Group.GET("/cluster", func(c *gin.Context) {
			method.K8sGetMethod(c)
		})

		//删除k8s集群配置信息
		v2Group.DELETE("/cluster", func(c *gin.Context) {
			method.K8sDeleteMethod(c)
		})
	}
}
