package utils

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

var secret []byte

func init() {
	var err error
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")   // optionally look for config in the working directory
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("jwt fatal error config file: %w", err))
	}
	secret = []byte("sdf")
}

type Claims struct {
	jwt.StandardClaims
	ID    string `json:"id"` // Account Id
	Email string `json:"email"`
	Role  string `json:"role"`
}

func NewClaims(id, email, role string, expiredAt int64) Claims {
	var c Claims
	c.ID = id
	c.Email = email
	c.ExpiresAt = expiredAt
	c.Role = role
	return c
}

func SignJwt(id, email, role string, expiredAt int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewClaims(id, email, role, expiredAt))
	tokenString, err := token.SignedString(secret)
	return tokenString, err
}

func VerifyJwt(tokenString string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return claims, errors.New("Unvalid token")
	}
}
