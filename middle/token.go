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
	maxAge := 60 * 10 //过期时间
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
		fmt.Println(err)
	}
	//fmt.Printf("token: %v\n", tokenString)
	return tokenString

	////解析token
	//ret,err :=ParseToken(tokenString)
	//if err!=nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("userinfo: %v\n", ret)
}

//解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
