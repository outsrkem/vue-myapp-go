package method

import (
	"github.com/gin-gonic/gin"
	"menu/interf"
	"menu/interf/impl"
	"net/http"
)

/*
	查询k8s集群配置信息
*/
func K8sGetMethod(c *gin.Context) {
	Qtype := c.Query("type")

	if Qtype == "cluster" || Qtype == "" {

		var face interf.K8sInterface
		var json impl.K8sBodyList

		face = json

		c.JSON(http.StatusOK, face.K8sBodyGetAll())

	} else if Qtype == "namespaces" {

		address := c.Query("address")

		var face interf.K8sNamespaceGetI
		var data impl.K8sNamespace

		face = &data
		face.K8sNamespaceGet(address)
		c.JSON(http.StatusOK, data)

	} else if Qtype == "workingload" {
		//获取Query参数
		namespaces := c.Query("namespaces")
		control := c.Query("control")
		address := c.Query("address")

		//判断参数是否为空
		if namespaces == "" || control == "" || address == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": "输入有误",
			})
		} else {
			//创建一个K8sInterface类型接口
			var face interf.K8sInterface
			//创建一个K8sDeploymentList结构体
			var dpm impl.K8sDeploymentList
			//把K8sDeploymentList结构体传给接口，用于识别接口中的方法
			face = &dpm
			//调用接口中的方法
			face.K8sDeploymentGet(namespaces, control, address)
			//返回数据K8sDeploymentList结构体类型数据，会转为json格式返回
			c.JSON(http.StatusOK, dpm)
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "参数错误",
		})
	}
}
