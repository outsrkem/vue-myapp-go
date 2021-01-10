package router

import (
	"github.com/gin-gonic/gin"
	"menu/interf"
	"menu/interf/impl"
	"net/http"
)

/*
	操作Linux信息路由配置
*/
func LinuxRouter(r *gin.Engine) {

	//获取服务器信息列表接口
	//路由分组
	v1Group := r.Group("/api/v1/common/resource")
	{
		//获取服务器性能列表
		v1Group.GET("/monitor", func(c *gin.Context) {
			//创建接口对象
			var face interf.LinuxInterface
			//创建结构体对象，LinuxList结构体是返回的数据格式
			var data impl.LinuxList
			//把结构体对象传给接口，让接口识别调用哪个结构体方法
			face = &data
			//调用接口方法实现
			face.GetAll()
			//返回json数据
			c.JSON(http.StatusOK, data)
		})

		//服务器列表添加
		v1Group.POST("/monitor", func(c *gin.Context) {
			//获取query的参数
			user := c.Query("user")
			ip := c.Query("ip")
			pwd := c.Query("pwd")
			port := c.Query("port")
			//创建结构体对象，LinuxListUser存入数据库的格式
			linuxData := impl.LinuxListUser{User: user, Ip: ip, Pwd: pwd, Port: port}

			if user != "" && ip != "" && pwd != "" && port != "" {
				var face interf.LinuxInterface
				var data impl.LinuxList
				face = &data
				face.Add(&linuxData)
				c.JSON(http.StatusOK, data)
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": "信息错误",
				})
			}

		})

		//删除服务器列表数据
		v1Group.DELETE("/monitor", func(c *gin.Context) {
			ip := c.Query("ip")

			if ip != "" {
				var face interf.LinuxInterface
				var data impl.LinuxList
				face = &data
				face.Del(ip)
				c.JSON(http.StatusOK, data)
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": "入参有误",
				})
			}
		})
	}
}
