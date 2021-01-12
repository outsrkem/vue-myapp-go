package method

import (
	"github.com/gin-gonic/gin"
	"menu/interf"
	"menu/interf/impl"
	"net/http"
)

/*
	删除登录用户
*/
func UserDeleteMethod(c *gin.Context) {
	username := c.Query("username")

	if username != "" {
		var face interf.UserTableI
		var data impl.UserTable

		face = &data
		face.UserDel(username)

		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "用户名不能为空",
		})
	}

}
