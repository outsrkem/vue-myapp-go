package method

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexMethod(c *gin.Context) {
	c.JSON(http.StatusOK, "没有请求参数")
}
