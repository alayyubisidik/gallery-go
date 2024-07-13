package helper

import (
	"errors"
	"gallery_go/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyCustomClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func CreateToken(user model.User) (string, error) {
	claims := MyCustomClaims{
		user.ID,
		user.Username,
		user.FullName,
		user.Email,
		user.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	return ss, err
}

func VerifyToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("metode tanda tangan tidak valid")
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token tidak valid")
	}
}
