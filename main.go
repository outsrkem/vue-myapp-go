package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"menu/middle"
	"menu/router"
)

func main() {
	//使用默认中间件创建一个gin路由器
	r := gin.Default()
	//解决跨域请求中间件
	r.Use(middle.Cors())

	// "/"请求路由
	router.Index(r)

	//Linux功能路由
	router.LinuxRouter(r)

	//k8s功能路由
	router.K8sRouter(r)

	//服务运行端口
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("服务运行错误")
	}
}
