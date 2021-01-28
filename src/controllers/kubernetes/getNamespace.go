package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"mana/src/filters/util"
	"mana/src/models"
	"net/http"
)

// 获取k8s名称空间
func GetNamespace(c *gin.Context) {
	//
	cid := c.Param("cid")
	k, _ := models.FindByK8sConfigParticulars(cid)
	selfLink := k.Server + "/api/v1/namespaces"
	log.Info("get Namespace, kubernetes selfLink: ", selfLink)
	k8s := util.NewK8sResources(k.CertificateAuthorityData, k.ClientCertificateData, k.ClientKeyData, selfLink)
	// 获取信息
	body := util.K8sResourcesGet(k8s)

	// 处理返回体
	array := gjson.Get(string(*body), "items.#.metadata.name").Array()
	num := len(array)
	namespace := make([]string, num)
	for i := 0; i < num; i++ {
		namespace[i] = array[i].Str
	}

	// 接口返回结构
	pageInfo := models.NewPageInfo(1, 10, 1, num)
	msg := models.NewResMessage("200", "successfully")
	response := models.NewResponse(namespace, pageInfo)
	returns := models.NewReturns(response, msg)

	c.JSON(http.StatusOK, returns)
}
