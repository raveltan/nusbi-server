package auth

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	db "nusbi-server/config"
	"strings"
)

/*
Error message number:
-1: Success
1: Unable to parse body
2: Username and password should be at least 5 and 8 character
3: User not found
*/

type loginRequest struct {
	Username string
	Password string
}

type loginResponse struct {
	Token   string
	Refresh string
	Error   int
}

func Login(c *fiber.Ctx) error {
	req := loginRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.JSON(loginResponse{Error: 1})
	}
	if len(req.Username) < 5 || len(req.Password) < 8 {
		return c.JSON(loginResponse{Error: 2})
	}
	sqlRes, err := db.Db.Query("select role,password from Users where user_id = ?",
		strings.ToLower(req.Username),
	)
	if err != nil {
		return c.SendStatus(500)
	}
	if err = sqlRes.Err(); err != nil {
		return c.SendStatus(500)
	}
	var role string
	for sqlRes.Next(){
		var pass string
		err = sqlRes.Scan(&role,&pass)
		if err != nil {
			return c.SendStatus(500)
		}
		err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(req.Password))
		if err != nil{
			role = ""
		}
	}
	if role == "" {
		return c.JSON(loginResponse{Error: 3})
	}
	token, err := createToken(req.Username, role)
	if err != nil {
		return c.SendStatus(500)
	}
	return c.JSON(loginResponse{
		Token:   token,
		Refresh: token,
		Error:   -1,
	})
}
