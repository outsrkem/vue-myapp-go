package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mana/src/models"
	"net/http"
)

// 全局token校验
func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过http header中的token解析来认证
		token := c.Request.Header.Get("X-Auth-Token")
		if token == "" {
			fmt.Println("请求头中未携带token,403")
			msg := models.NewResMessage("403", "The token X-Auth-Token is not carried in the request header")
			c.JSON(http.StatusForbidden, msg)
			c.Abort()
			return
		}

		claims, err := ParseToken(token)
		if err != nil {
			fmt.Println("token格式无效或已过期,401")
			msg := models.NewResMessage("401", "The token format is invalid or expired")
			c.JSON(http.StatusUnauthorized, msg)
			c.Abort()
			return
		}

		//if token 中的状态 != 当前用户状态 {
		//	token失效，重新获取
		//}

		// 将解析后的有效载荷claims重新写入gin.Context引用对象中
		c.Set("uid", claims.Id)
		c.Set("role", claims.Audience)
	}
}
