package method

import (
	"github.com/gin-gonic/gin"
	"menu/interf"
	"menu/interf/impl"
	"net/http"
)

func LinuxGetMehod(c *gin.Context) {
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
}
