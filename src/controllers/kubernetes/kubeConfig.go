package kubernetes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"mana/src/config"
	"mana/src/models"
	"net/http"
	"strconv"
)

var log = config.Log()

// 插入配置k8s鉴权配置文件
func InstKubeConfig(c *gin.Context) {
	var k = models.NewKubeConfig()

	// 获取token中的userid
	uid := c.MustGet("uid").(string)

	k8sConf := make(map[string]interface{})
	c.ShouldBind(&k8sConf)

	// 序列化
	data, err := json.Marshal(k8sConf)
	if err != nil {
		msg := models.NewResMessage("406", "The configuration file failed to parse JSON")
		c.JSON(http.StatusNotAcceptable, &msg)
		return
	}

	// 获取json数据
	k.USERID = uid
	k.CLUSTER_ALIAS = gjson.Get(string(data), "clusters.0.cluster.server").String()
	k.CLUSTER_USER = gjson.Get(string(data), "contexts.0.context.user").String()
	k.CURRENT_CONTEXT = gjson.Get(string(data), "current-context").String()
	k.SERVER = gjson.Get(string(data), "clusters.0.cluster.server").String()
	k.CERTIFICATE_AUTHORITY_DATA = gjson.Get(string(data), "clusters.0.cluster.certificate-authority-data").String()
	k.CLIENT_CERTIFICATE_DATA = gjson.Get(string(data), "users.0.user.client-certificate-data").String()
	k.CLIENT_KEY_DATA = gjson.Get(string(data), "users.0.user.client-key-data").String()

	if k.CLUSTER_USER == "" || k.CURRENT_CONTEXT == "" || k.SERVER == "" ||
		k.CERTIFICATE_AUTHORITY_DATA == "" || k.CLIENT_CERTIFICATE_DATA == "" || k.CLIENT_KEY_DATA == "" {
		msg := models.NewResMessage("406", "Missing configuration data")
		c.JSON(http.StatusNotAcceptable, &msg)
		return
	}

	models.InstKubeConfig(k)

	msg := models.NewResMessage("201", "Creating KubeConfig a successful")
	c.JSON(http.StatusCreated, &msg)
}

// 返回集群配置文件信息
func GetKubeConfig(c *gin.Context) {

	// 获取Query参数，转换为int类型
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	// 获取uid，(token 中获取)
	uid := c.MustGet("uid").(string)

	res := models.FindByKubeConfigs(uid, pageSize, page)
	msg := models.NewResMessage("200", "successfully")
	returns := models.NewReturns(res, msg)
	c.JSON(http.StatusOK, returns)
}

// 删除集群配置
func DelKubeConfig(c *gin.Context) {
	configId, _ := strconv.Atoi(c.Query("config_id"))
	log.Info(configId)
	row := models.DeleteKubeConfig(configId)
	if row < 1 {
		if row == 0 {
			msg := models.NewResMessage("200", "Delete 0 records")
			c.JSON(http.StatusOK, msg)
			return
		} else if row == -1 {
			msg := models.NewResMessage("202", "Remove abnormal")
			c.JSON(http.StatusAccepted, msg)
			return
		}
	}
	msg := models.NewResMessage("200", "Delete successful")
	c.JSON(http.StatusOK, msg)
}
