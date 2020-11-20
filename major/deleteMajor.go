package major

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	db "nusbi-server/config"
)

type deleteMajorRequest struct {
	ID string
}

type deleteMajorResponse struct {
	Error int
}

/*
-1:Success
1: Token role invalid
2: Body parsing failed
3: Failed
*/

func DeleteMajor(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.JSON(createMajorResponse{Error: 1})
	}
	request := deleteMajorRequest{}
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(deleteMajorResponse{Error: 2})
	}
	if _, err := db.Db.Exec("delete  from nusbiam.Majors where major_id = ?", request.ID); err != nil {
		return c.JSON(deleteMajorResponse{Error: 3})
	}
	return c.JSON(deleteMajorResponse{Error: -1})
}
