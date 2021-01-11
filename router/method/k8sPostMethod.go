package method

import (
	"github.com/gin-gonic/gin"
	"menu/interf"
	"menu/interf/impl"
	"net/http"
)

func K8sPostMethod(c *gin.Context) {
	var face interf.K8sInterface
	var json impl.K8sBodyList

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
