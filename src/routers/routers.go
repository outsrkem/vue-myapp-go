package routers

import (
	"github.com/gin-gonic/gin"
	"mana/src/controllers/kubernetes"
	"mana/src/controllers/navigation"
	"mana/src/controllers/resource"
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

	// 路由
	v1Group := r.Group("/api/v1")
	{

		v1Group.GET("/common/user/userinfo/:uid", user.FindByUserinfo)
		//v1Group.GET("/common/user/userinfo/:uid/test", user.FindByUserInfos)

		// 获取导航链接列表，添加链接，编辑，删除，获取单条导航链接记录
		v1Group.GET("/common/navigation/links", navigation.GetResourceLinks)
		v1Group.POST("/common/navigation/links", navigation.AddResourceLink)
		v1Group.GET("/common/navigation/links/:id", navigation.GetResourceLinks)
		v1Group.PATCH("/common/navigation/links/:id", navigation.UpdateResourceLink)
		v1Group.DELETE("/common/navigation/links/:id", navigation.DeleteResourceLink)


		// 主机监控
		v1Group.GET("/common/resource/monitor", resource.GetHostMonitorInfo)

		// 添加k8s配置文件
		v1Group.POST("/common/kubernetes/cluster", kubernetes.InstKubeConfig)

		// 获取集群配置列表
		v1Group.GET("/common/kubernetes/cluster", kubernetes.GetKubeConfig)

		// 删除集群配置
		v1Group.DELETE("/common/kubernetes/cluster", kubernetes.DelKubeConfig)

		// 获取k8s名称空间/common/kubernetes/cluster/:cid/work/namespaces
		v1Group.GET("/common/kubernetes/cluster/:cid/namespaces", kubernetes.GetNamespace)

		// 获取k8s控制器资源
		v1Group.GET("/common/kubernetes/cluster/:cid/control/:namespaces/:control", kubernetes.GetKubernetesControl)

		// 获取k8s工作负载详细信息
		// /apis/apps/v1/namespaces/kube-system/daemonsets/kube-flannel-ds-amd64
		v1Group.GET("/common/kubernetes/cluster/:cid/pods/:namespaces/:control/:podsName", kubernetes.GetKubernetesPods)
	}
}
