package resource

import (
	"github.com/gin-gonic/gin"
	"mana/src/models"
	"net/http"
)

func GetHostMonitorInfo(c *gin.Context) {
	// 临时返回
	msg := models.NewResMessage("200", "successfully")
	pageInfo := models.NewPageInfo(1, 1, 1, 1)
	items := make([]map[string]string, 0)
	res := models.NewResponse(items, pageInfo)
	returns := models.NewReturns(res, msg)
	c.JSON(http.StatusOK, returns)
}
