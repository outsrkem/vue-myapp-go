package method

import (
	"github.com/gin-gonic/gin"
	"menu/interf"
	"menu/interf/impl"
	"net/http"
)

func UserGetMethod(c *gin.Context) {
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
}
