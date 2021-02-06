package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 加密密码
func PasswordBcrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		fmt.Println("加密密码失败", err)
		return "", err
	}
	// 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	encodePassword := string(hash)
	return encodePassword, err
}

// PasswordAuthentication 密码校验
func PasswordAuthentication(loginPassword string, password string) bool {
	// 密码验证
	// password 加密的密码
	// 用户登录时输入的密码 loginPassword
	// loginPassword = "123456"
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(loginPassword)); err != nil { //验证（对比）
		return false
	}
	// 密码校验通过
	return true
}
