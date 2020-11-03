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
2: Username should be at least 5 characters
3: Password should be at least 8 characters
4: Role should be a,t or s
*/

type createUserRequest struct {
	Username string
	Password string
	Role     string
}


type createUserResponse struct {
	Error int
}

func CreateUser(c *fiber.Ctx) error {
	var err error
	// TODO: Allow the addition user data
	request := createUserRequest{}
	err = c.BodyParser(&request)
	if err != nil {
		return c.JSON(createUserResponse{Error: 1})
	}

	// Data validation
	if len(request.Username) < 5 {
		return c.JSON(createUserResponse{
			Error: 2,
		})
	}
	if len(request.Password) < 8 {
		return c.JSON(createUserResponse{Error: 3})
	}
	if !(request.Role == "a" || request.Role == "t" || request.Role == "s") {
		return c.JSON(createUserResponse{Error: 4})
	}

	// Add user to config
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 7)
	if err != nil {
		return c.SendStatus(500)
	}
	_, err = db.Db.Exec(
		"insert into Users (user_id, password, role) VALUES (?,?,?)",
		strings.ToLower(request.Username),
		string(hash),
		request.Role,
	)
	if err != nil {
		return c.SendStatus(500)
	}
	return c.JSON(createUserResponse{Error: -1})
}
