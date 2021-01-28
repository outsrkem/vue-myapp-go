package kubernetes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetKubernetesResource(c *gin.Context) {
	cid := c.Param("cid")
	control := c.Param("control")
	namespaces := c.Param("namespaces")

	// selfLink
	// /apis/apps/v1/namespaces/kube-system/daemonsets/kube-flannel-ds-amd64
	k8sLink := "/apis/apps/v1/namespaces/" + namespaces + "/" + control

	if control == "nil" || namespaces == "" {
		log.Error("namespaces namespaces 无效")
		return
	}
	log.Info(cid)
	if control == "deployments" {
		// deployments

		returns := GetWorkingLoad(cid, k8sLink)
		c.JSON(http.StatusOK, returns)

	} else if control == "daemonset" {
		// daemonset
		returns := GetWorkingLoad(cid, k8sLink)
		c.JSON(http.StatusOK, returns)
	}
}
