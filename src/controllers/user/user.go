package user

import (
	"github.com/gin-gonic/gin"
	"mana/src/filters/utility"
	"mana/src/models"
	"net/http"
)

// 注册
func InstUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	// 加密密码
	encodePassword := utility.PasswordBcrypt(c.Query("password"))
	id := models.InstUser(c.Query("username"), encodePassword)
	c.JSON(http.StatusOK, gin.H{
		"id":       id,
		"num":      id,
		"password": password,
		"username": username,
	})
}

// 登录
func Login(c *gin.Context) {
	username := c.Query("username")
	loginPassword := c.Query("password")
	result := models.SelectUserQueryRow(username)
	utility.PasswordAuthentication(loginPassword, result.PASSWD)
	c.JSON(http.StatusOK, gin.H{
		"password": loginPassword,
		"username": username,
	})
}
