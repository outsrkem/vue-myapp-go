package router

import (
	"github.com/gin-gonic/gin"
	"menu/interf"
	"menu/interf/impl"
	"net/http"
)

func UserRouter(r *gin.Engine) {
	v3Group := r.Group("/api/v1/common/user")
	{
		v3Group.GET("/table", func(c *gin.Context) {
			username := c.Query("username")
			if username != "" {
				var face interf.UserTableI
				var data impl.UserTable
				face = &data
				face.UserGet(username)
				c.JSON(http.StatusOK, data)
			} else {
				var face interf.UserTableI
				var data impl.UserTable
				face = &data
				face.UserGetAll()
				c.JSON(http.StatusOK, face.UserGetAll())
			}
		})

		v3Group.POST("/table", func(c *gin.Context) {
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
		})

		v3Group.DELETE("/table", func(c *gin.Context) {

		})
	}
}
