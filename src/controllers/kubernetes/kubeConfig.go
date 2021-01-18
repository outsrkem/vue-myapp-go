package kubernetes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"mana/src/config"
	"mana/src/models"
	"net/http"
)

var log = config.Log()

// 插入配置k8s鉴权配置文件
func InstKubeConfig(c *gin.Context) {
	var k = models.NewKubeConfig()

	// 获取token中的userid
	uid := c.MustGet("uid").(string)

	k8sConf := make(map[string]interface{})
	c.BindJSON(&k8sConf)

	// 序列化
	data, _ := json.Marshal(k8sConf)
	fmt.Printf(string(data))

	k.SERVER = "https://10.10.10.31:6443"
	k.CLIENT_KEY_DATA = k8sConf["users"]
	log.Info(k.SERVER)
	k.USERID = uid
	models.InstKubeConfig(k)

	msg := models.NewResMessage("201", "Creating a successful")
	c.JSON(http.StatusOK, &msg)
}
