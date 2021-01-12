package method

import (
	"github.com/gin-gonic/gin"
	"menu/interf"
	"menu/interf/impl"
	"net/http"
)

/*
	查询登录用户
*/
func UserGetMethod(c *gin.Context) {
	//接收请求参数
	username := c.Query("username")

	var face interf.UserTableI
	var data impl.UserTable

	if username != "" {
		face = &data
		face.UserGet(username)

		c.JSON(http.StatusOK, data)
	} else {
		face = &data
		face.UserGetAll()

		c.JSON(http.StatusOK, face.UserGetAll())
	}
}
