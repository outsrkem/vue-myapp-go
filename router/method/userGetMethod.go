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
	username := c.Query("username")
	pagesize := c.Query("pagesize")
	pagenum := c.Query("pagenum")
	var face interf.UserTableI
	var data impl.UserTable

	// 搜索前缀
	if username != "" {
		face = &data
		face.UserGet(username)

		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   face.UserGet(username),
		})
		//分页查询
	} else if pagesize != "" && pagenum != "" {
		face = &data
		data, num, total := face.UserPageGet(pagesize, pagenum)

		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"total":   total,
			"pagenum": num,
			"data":    &data,
		})
		//查询全部
	} else {
		face = &data
		face.UserGetAll()

		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   face.UserGetAll(),
		})
	}
}
