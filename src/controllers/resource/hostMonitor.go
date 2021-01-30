package resource

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHostMonitorInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "successfully",
	})
}
