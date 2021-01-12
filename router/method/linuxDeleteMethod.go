package method

import (
	"github.com/gin-gonic/gin"
	"menu/interf"
	"menu/interf/impl"
	"net/http"
)

/*
	删除Linux服务器账号信息
*/
func LinuxDeleteMethod(c *gin.Context) {
	ip := c.Query("ip")

	if ip != "" {
		var face interf.LinuxInterface
		var data impl.LinuxList

		face = &data
		face.Del(ip)

		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "入参有误",
		})
	}
}
