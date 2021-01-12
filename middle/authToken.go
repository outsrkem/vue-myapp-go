package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过http header中的token解析来认证
		token := c.Request.Header.Get("Token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
				"data":   nil,
			})
			c.Abort()
			return
		}

		claims, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "token格式无效或已过期",
			})
			c.Abort()
			return
		}

		// 将解析后的有效载荷claims重新写入gin.Context引用对象中
		c.Set("role", claims.Audience)
	}
}
