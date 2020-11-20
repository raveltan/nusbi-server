package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func createToken(username string,role string) (string,error){
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("yigeiwoligiaogiao"))
	if err != nil {
		return "",err
	}
	return t,nil
}

func createRefreshToken(username string,role string) (string,error){
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 4380).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("aoligei"))
	if err != nil {
		return "",err
	}
	return t,nil
}