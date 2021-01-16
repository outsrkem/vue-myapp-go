package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mana/src/filters/utility"
	"mana/src/models"
	"mana/src/models/impl"
	"net/http"
	"time"
)

// 用户注册
func InstUser(c *gin.Context) {
	var userRegisterInfo impl.UserRegisterInfo

	if err := c.BindJSON(&userRegisterInfo); err != nil {
		fmt.Println("用户注册获取数据error", err)
	}
	password := userRegisterInfo.Password
	username := userRegisterInfo.Username
	// 加密密码
	encodePassword, _ := utility.PasswordBcrypt(password)
	id := models.InstUser(username, encodePassword)
	c.JSON(http.StatusCreated, gin.H{
		"id":       id,
		"username": userRegisterInfo.Username,
	})
}

// 用户登录
func Login(c *gin.Context) {
	var userLoginInfo impl.UserLoginInfo
	if err := c.BindJSON(&userLoginInfo); err != nil {
		fmt.Println("用户登录获取数据error", err)
	}
	username := userLoginInfo.Username
	loginPassword := userLoginInfo.Password
	fmt.Println(username)
	fmt.Println(loginPassword)
	result := models.SelectUserQueryRow(username)
	// 校验密码
	fmt.Println(result.PASSWD)
	err := utility.PasswordAuthentication(loginPassword, result.PASSWD)
	if err != nil {
		fmt.Println("登录错误", err)
		var user impl.MetaInfo
		user.RequestTime = time.Now().UnixNano()
		user.Msg = "Logon failed"
		user.Status = "1"
		c.JSON(http.StatusUnauthorized, &user)
	} else {
		var user impl.UserLoginStruct
		meta := &user.MetaInfo
		resp := &user.Response
		meta.RequestTime = time.Now().UnixNano()
		meta.Msg = "Successfully"
		meta.Status = "200"

		resp.Userid = result.USERID
		resp.Username = result.USERNAME
		resp.Nickname = result.NICKNAME
		resp.Role = result.ROLE
		resp.Expires = result.EXPIRES
		resp.Token = "this is token"
		c.JSON(http.StatusOK, &user)
	}
}
