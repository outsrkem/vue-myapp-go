package navigation

import (
	"github.com/gin-gonic/gin"
	"mana/src/config"
	"mana/src/models"
	"net/http"
	"strconv"
)

var log = config.Log()

// 获取导航链接
func GetResourceLinks(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "-1"))

	// -1 代表请求中没有传递id，则查询所有
	if id == -1 {
		res, err := models.FindByResourceLinks()
		if err != nil {
			msg := models.NewResMessage("500", "Query database failed")
			c.JSON(http.StatusInternalServerError, msg)
			return
		}

		msg := models.NewResMessage("200", "successfully")
		returns := models.NewReturns(res, msg)
		c.JSON(http.StatusOK, returns)
		return
	}

	// id为查询的具体记录，用于删除
	res, err := models.FindByResourceLinksTheId(id)
	if err != nil {
		msg := models.NewResMessage("500", "Query database failed")
		c.JSON(http.StatusInternalServerError, msg)
		return
	}
	msg := models.NewResMessage("200", "successfully")
	returns := models.NewReturns(res, msg)
	c.JSON(http.StatusOK, returns)
	return

}
