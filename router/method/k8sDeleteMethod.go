package method

import (
	"github.com/gin-gonic/gin"
	"menu/interf"
	"menu/interf/impl"
	"net/http"
)

/*
	删除k8s集群配置信息
*/
func K8sDeleteMethod(c *gin.Context) {
	address := c.Query("address")

	if address != "" {
		var face interf.K8sInterface
		var json impl.K8sBodyList

		face = json
		face.K8sBodyDel(address)

		c.JSON(http.StatusOK, gin.H{
			"status": "删除成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "输入错误",
		})
	}
}
