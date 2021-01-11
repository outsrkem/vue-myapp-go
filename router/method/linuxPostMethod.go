package method

import (
	"github.com/gin-gonic/gin"
	"menu/interf"
	"menu/interf/impl"
	"net/http"
)

func LinuxPostMethod(c *gin.Context) {
	//获取query的参数
	user := c.Query("user")
	ip := c.Query("ip")
	pwd := c.Query("pwd")
	port := c.Query("port")
	//创建结构体对象，LinuxListUser存入数据库的格式
	linuxData := impl.LinuxListUser{User: user, Ip: ip, Pwd: pwd, Port: port}

	if user != "" && ip != "" && pwd != "" && port != "" {
		var face interf.LinuxInterface
		var data impl.LinuxList
		face = &data
		face.Add(&linuxData)
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "信息错误",
		})
	}
}
