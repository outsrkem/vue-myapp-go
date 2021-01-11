package method

import (
	"github.com/gin-gonic/gin"
	"menu/interf"
	"menu/interf/impl"
	"net/http"
)

func UserDeleteMethod(c *gin.Context) {
	username := c.Query("username")
	var face interf.UserTableI
	var data impl.UserTable
	face = &data
	face.UserDel(username)
	c.JSON(http.StatusOK, data)
}
