package method

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"menu/interf"
	"menu/interf/impl"
	"net/http"
)

/*
	添加登录用户
*/
func UserPostMethod(c *gin.Context) {
	//接收query参数
	username := c.Query("username")
	password := c.Query("password")
	role := c.Query("role")
	userstatus := c.Query("user_status")

	if username != "" && password != "" && role != "" && userstatus != "" {
		//创建接口和结构体对象
		var face interf.UserTableI
		data := impl.UserTable{
			UserName:   username,
			Password:   password,
			Role:       role,
			UserStatus: cast.ToBool(userstatus),
		}
		//调用接口add方法
		face = &data
		face.UserAdd()

		c.JSON(http.StatusOK, gin.H{
			"status": 200,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "输入有误",
		})
	}
}
