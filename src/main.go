package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
import "mana/src/routers"

func main() {
	r := gin.Default()

	// 服务路由请求
	routers.Index(r)

	// 服务运行
	err := r.Run(":9999")
	if err != nil {
		fmt.Println("服务运行错误")
	}
}
