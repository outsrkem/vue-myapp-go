package kubernetes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetKubernetesResource(c *gin.Context) {
	cid := c.Param("cid")
	resourceName := c.Param("resource")
	k8sLink := c.DefaultQuery("k8s_link", "-1")
	if k8sLink == "-1" {
		log.Error("selfLink 无效")
		return
	}
	log.Info(cid)
	if resourceName == "namespaces" {
		returns := GetNamespace(cid, k8sLink)
		c.JSON(http.StatusOK, returns)

	} else if resourceName == "deployments" {
		// deployments
		returns := GetWorkingLoad(cid, k8sLink)
		c.JSON(http.StatusOK, returns)

	} else if resourceName == "daemonset" {
		// daemonset
		returns := GetWorkingLoad(cid, k8sLink)
		c.JSON(http.StatusOK, returns)
	}
}
