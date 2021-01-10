package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"menu/middle"
	"menu/router"
)

var (
	port int
)

//初始化函数，用于设置端口
func init() {
	flag.IntVar(&port, "p", 8080, "set server port")
}

func main() {
	//命令行解析工具
	flag.Parse()
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

	//用户功能路由
	router.UserRouter(r)

	//服务运行端口
	err := r.Run(":" + cast.ToString(port))
	if err != nil {
		fmt.Println("服务运行错误")
	}
}
