package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"mynamebvh.com/blog/config"
)

type TokenStruct struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

type JwtTokenInterface interface {
	Sign(claims jwt.MapClaims) TokenStruct
}

func Sign(claims jwt.MapClaims) TokenStruct{
	timeNow := time.Now()

	// jwtExpired, _:= time.ParseDuration(config.GetEnv("JWT_EXPIRED"))
	tokenExpired := timeNow.Add(time.Hour * 2).Unix()
 
	if claims["id"] == nil {
		return TokenStruct{}
	}

	token := jwt.New(jwt.SigningMethodHS256)

	// setup userdata
	var _, checkExp = claims["exp"]

	if !checkExp {
		claims["exp"] = tokenExpired
	}

	claims["token_type"] = "access_token"
	
	token.Claims = claims
	
	t, err := token.SignedString([]byte(config.GetEnv("JWT_SECRET")))

	if(err != nil){
		return TokenStruct{}
	}

	return TokenStruct{
		Type:  "Bearer",
		Token: t,
	}
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}