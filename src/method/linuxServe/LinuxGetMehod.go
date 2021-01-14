package linuxServe

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	查询Linux服务器账号信息
*/

func LinuxGetMethod(c *gin.Context) {

	c.JSON(http.StatusOK, "ok")
}
