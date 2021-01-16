package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mana/src/models/impl"
	"net/http"
	"time"
)

// 全局token校验
func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过http header中的token解析来认证
		token := c.Request.Header.Get("X-Auth-Token")
		if token == "" {
			var metaInfo impl.GeneralErrorStruct
			fmt.Println("请求头中未携带token,403")
			metaInfo.Code = "403"
			metaInfo.Msg = "The token X-Auth-Token is not carried in the request header"
			metaInfo.RequestTime = time.Now().UnixNano()
			c.JSON(http.StatusForbidden, &metaInfo)
			c.Abort()
			return
		}

		claims, err := ParseToken(token)
		if err != nil {
			var metaInfo impl.GeneralErrorStruct
			fmt.Println("token格式无效或已过期,401")
			metaInfo.Code = "401"
			metaInfo.Msg = "The token format is invalid or expired"
			metaInfo.RequestTime = time.Now().UnixNano()
			c.JSON(http.StatusUnauthorized, &metaInfo)
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
