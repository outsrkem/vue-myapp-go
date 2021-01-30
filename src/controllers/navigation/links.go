package navigation

import (
	"github.com/gin-gonic/gin"
	"mana/src/models"
	"net/http"
)

// 获取导航链接
func GetResourceLinks(c *gin.Context) {
	res, err := models.FindByResourceLinks()
	if err != nil {
		msg := models.NewResMessage("500", "Query database failed")
		c.JSON(http.StatusInternalServerError, msg)
		return
	}

	msg := models.NewResMessage("200", "successfully")
	returns := models.NewReturns(res, msg)
	c.JSON(http.StatusOK, returns)
}
