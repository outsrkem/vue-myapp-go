package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"menu/interf/impl"
	"menu/middle"
	"menu/router"
)

var (
	port          int
	adminUser     string
	adminPassword string
)

//初始化函数，用于设置端口
func init() {
	flag.IntVar(&port, "port", 8080, "set server port")
	flag.StringVar(&adminUser, "user", "admin", "set admin username")
	flag.StringVar(&adminPassword, "pwd", "admin", "set admin password")
}

func main() {
	//命令行解析工具
	flag.Parse()

	//初始化登录用户，初次运行默认账号为：admin，密码：admin，如果要指定管理员账号，可以查询命令行操作
	impl.Initialize(adminUser, adminPassword)

	//使用默认中间件创建一个gin路由器
	r := gin.Default()

	//解决跨域请求中间件
	r.Use(middle.Cors())

	// 服务路由请求
	router.Index(r)

	//服务运行端口
	err := r.Run(":" + cast.ToString(port))
	if err != nil {
		fmt.Println("服务运行错误")
	}
}
