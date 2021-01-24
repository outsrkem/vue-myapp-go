package kubernetes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"mana/src/config"
	"mana/src/filters/util"
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

// 获取工作负载

func GetWorkingLoad(c *gin.Context) {
	cid := c.Param("cid")              // 获取路径参数
	namespace := c.Param("namespaces") // 获取路径参数
	control := c.Param("control")      // 获取路径参数

	// 调用函数，获取负载信息
	k, _ := models.FindByK8sConfigParticulars(cid)
	// 获取名称空间下 deploy
	url := k.Server + "/apis/apps/v1/namespaces/" + namespace + "/" + control
	// 传递证书及url
	k8s := util.NewK8sResources(k.CertificateAuthorityData, k.ClientCertificateData, k.ClientKeyData, url)
	// 获取信息
	body := util.K8sResourcesGet(k8s)



	fmt.Println(string(*body))
	/**
	遍历json数据中item数组，使用gjson模块
	*/
	////定义索引i，gjson遍历数组，没有索引，需要自己设置
	//i := 0
	//
	////计算数组长度
	//n := len(gjson.Get(string(*body), "items").Array())
	////初始化一个ItemsDpm结构体数组，用于存放遍历数据
	//item := make([]map[string]string, n)

	//foreach遍历item数组，如果遍历完成，返回false

	// 返回数据
	msg := models.NewResMessage("200", "Successful")
	c.JSON(http.StatusOK, msg)
}
