package kubernetes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取k8s控制器资源
func GetKubernetesControl(c *gin.Context) {
	cid := c.Param("cid")
	control := c.Param("control")
	namespaces := c.Param("namespaces")

	// selfLink
	// /apis/apps/v1/namespaces/kube-system/daemonsets/kube-flannel-ds-amd64
	k8sLink := "/apis/apps/v1/namespaces/" + namespaces + "/" + control

	log.Info("get ", control, ", k8sLink: ", k8sLink)

	// 根据不同的控制器加载不同的资源
	switch control {
	case "deployments":
		returns := GetWorkingLoad(cid, k8sLink)
		c.JSON(http.StatusOK, returns)
	case "daemonset":
		returns := GetWorkingLoad(cid, k8sLink)
		c.JSON(http.StatusOK, returns)
	}
}

// 获取k8s的pods资源
func GetKubernetesPods(c *gin.Context) {
	c.JSON(http.StatusOK, "获取k8s的pods资源")
}
