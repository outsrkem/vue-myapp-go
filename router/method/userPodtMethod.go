package method

import (
	"github.com/gin-gonic/gin"
	"menu/interf"
	"menu/interf/impl"
	"net/http"
)

func UserPostMethod(c *gin.Context) {
	//接收query参数
	username := c.Query("username")
	password := c.Query("password")
	role := c.Query("role")
	status := c.Query("status")

	if username != "" && password != "" && role != "" && status != "" {
		//创建接口和结构体对象
		var face interf.UserTableI
		data := impl.UserTable{
			UserName: username,
			Password: password,
			Role:     role,
			Status:   status,
		}
		face = &data
		//调用接口add方法
		face.UserAdd()
		c.JSON(http.StatusOK, gin.H{
			"status": "用户添加成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "输入有误",
		})
	}
}
