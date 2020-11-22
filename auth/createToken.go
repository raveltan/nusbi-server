package auth

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"time"
)

func createToken(username string, role string) (string, error) {
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
		return "", err
	}
	return t, nil
}

func createRefreshToken(username string, role string) (string, error) {
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
		return "", err
	}
	return t, nil
}

type tokenRefreshResponse struct {
	Token   string
	Refresh string
}

func RefreshToken(c *fiber.Ctx) error {
	data := c.Locals("user").(*jwt.Token)
	claims := data.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	role := claims["role"].(string)

	token, err := createToken(name, role)
	if err != nil {
		return c.SendStatus(500)
	}
	refresh, err := createRefreshToken(name, role)
	if err != nil {
		return c.SendStatus(500)
	}
	var result = tokenRefreshResponse{
		Token:   token,
		Refresh: refresh,
	}
	return c.JSON(result)
}
