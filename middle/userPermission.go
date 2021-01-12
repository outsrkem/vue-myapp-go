package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//普通权限
func General(c *gin.Context) {
	//fmt.Println(c.MustGet("role").(string))
	if role := c.MustGet("role").(string); role != "general" && role != "admin" {
		c.JSON(http.StatusOK, gin.H{
			"status": "账号权限不足1",
		})
		c.Abort()
		return
	}
}

//管理员权限
func Admin(c *gin.Context) {
	//fmt.Println(c.MustGet("role").(string))
	if c.MustGet("role").(string) != "admin" {
		c.JSON(http.StatusOK, gin.H{
			"status": "账号权限不足",
		})
		c.Abort()
		return
	}
}
