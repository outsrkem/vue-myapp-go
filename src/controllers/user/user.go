package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mana/src/filters/util"
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
	// 用户名或密码不能为空
	if password == "" || username == "" {
		fmt.Println("注册错误,用户名或密码为空")
		var metaInfo impl.GeneralErrorStruct
		metaInfo.Code = "2014"
		metaInfo.Msg = "The user name or password cannot be empty"
		metaInfo.RequestTime = time.Now().UnixNano()
		c.JSON(http.StatusForbidden, &metaInfo)
		return
	}
	// 加密密码
	encodePassword, _ := util.PasswordBcrypt(password)
	userId := models.InstUser(username, encodePassword)

	var user impl.UserRegisterStruct
	meta := &user.MetaInfo
	resp := &user.Response

	(*meta).RequestTime = time.Now().UnixNano()
	(*meta).Msg = "registered successfully"
	(*meta).Code = "201"
	resp.Userid = userId
	resp.Username = username
	c.JSON(http.StatusCreated, &user)
}

// 用户登录
func Login(c *gin.Context) {
	var userLoginInfo impl.UserLoginInfo
	if err := c.BindJSON(&userLoginInfo); err != nil {
		fmt.Println("用户登录获取数据error", err)
	}
	username := userLoginInfo.Username
	loginPassword := userLoginInfo.Password
	// 用户名或密码不能为空
	if loginPassword == "" || username == "" {
		fmt.Println("登录错误,用户名或密码为空")
		var metaInfo impl.GeneralErrorStruct
		metaInfo.Code = "2014"
		metaInfo.Msg = "The user name or password cannot be empty"
		metaInfo.RequestTime = time.Now().UnixNano()
		c.JSON(http.StatusForbidden, &metaInfo)
		return
	}

	result, _ := models.SelectUserQueryRow(username)
	fmt.Println("登录用户", username)
	// 校验密码
	err := util.PasswordAuthentication(loginPassword, result.PASSWD)
	if err != nil {
		fmt.Println("登录错误", err)
		var user impl.MetaInfo
		user.RequestTime = time.Now().UnixNano()
		user.Msg = "Logon failed"
		user.Code = "1"
		c.JSON(http.StatusUnauthorized, &user)
	} else {
		var user impl.UserLoginStruct
		meta := &user.MetaInfo
		resp := &user.Response

		// 生成token
		token := util.EncodeAuthToken(result.USERID, result.USERNAME, result.ROLE)

		// 构造返回数据
		meta.RequestTime = time.Now().UnixNano()
		meta.Msg = "login successfully"
		meta.Code = "200"
		resp.Userid = result.USERID
		resp.Username = result.USERNAME
		resp.Nickname = result.NICKNAME
		resp.Role = result.ROLE
		resp.Expires = result.EXPIRES
		resp.Token = token
		c.JSON(http.StatusOK, &user)
	}
}
