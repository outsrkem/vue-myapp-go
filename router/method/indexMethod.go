package method

import (
	"github.com/gin-gonic/gin"
	"menu/interf"
	"menu/interf/impl"
	"net/http"
)

/*
	用户登录
*/
func LoginMethod(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": "登录失败，账号或密码不能为空",
		})
	} else {
		var face interf.UserTableI
		data := impl.UserTable{
			UserName: username,
			Password: password,
		}
		face = &data

		if face.UserLogin() == "成功" {
			c.JSON(http.StatusOK, data)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "登录失败，账号或密码错误",
			})
		}
	}
}
