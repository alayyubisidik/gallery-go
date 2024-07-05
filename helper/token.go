package helper

import (
	"errors"
	"gallery_go/model/domain"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type MyCustomClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func signingKey() []byte {
	config := viper.New()
	config.SetConfigFile("D:/project/gallery_go/config.env")

	err := config.ReadInConfig()
	PanicIfError(err)

	mySigningKey := config.GetString("JWT_SECRET_KEY")

	return []byte(mySigningKey)
}


func CreateToken(user domain.User) (string, error) {

	claims := MyCustomClaims{
		user.ID,
		user.Username,
		user.FullName,
		user.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey())

	return ss, err
}

func VerifyToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signature method")
		}
		return signingKey(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token is invalid")
	}
}
