package middle

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SECRETKEY = "243223ffslsfsldfl412fdsfsdf" //私钥
)

//自定义Claims
type CustomClaims struct {
	jwt.StandardClaims
}

func Token(username, role string) string {
	//生成token
	maxAge := 60 * 60 //过期时间
	customClaims := &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
			Issuer:    username,                                                   // 非必须，也可以填充用户名，
			Audience:  role,                                                       // 非必须
		},
	}
	//采用HMAC SHA256加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	//私钥加密
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		fmt.Println("SignedString加密错误：", err)
		return "SignedString加密错误"
	}
	//fmt.Printf("token: %v\n", tokenString)
	return tokenString
}

//解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})

	if err != nil {
		fmt.Println("token解析错误：", err)
		return nil, err
	}

	//解析为CustomClaims结构体
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		fmt.Println("claims解析错误：", err)
		return nil, err
	}
}
