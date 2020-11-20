package config

import "github.com/form3tech-oss/jwt-go"

func GetRoleFromToken(data *jwt.Token) string{
	claims := data.Claims.(jwt.MapClaims)
	return claims["role"].(string)
}
