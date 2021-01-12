package method

import (
	"github.com/gin-gonic/gin"
	"menu/interf"
	"menu/interf/impl"
	"net/http"
)

/*
	添加k8s集群配置信息
*/

func K8sPostMethod(c *gin.Context) {
	var face interf.K8sInterface
	var json impl.K8sBodyList

	//绑定获取的json数据为自定义的结构体
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": err,
		})
	} else {
		face = json
		face.K8sBodyAdd()

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}
