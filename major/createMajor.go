package major

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	db "nusbi-server/config"
)

type createMajorResponse struct {
	Error int
}

type createMajorRequest struct{
	Name string
}

/*
-1:Success
1: Token role invalid
2: Body parsing failed
3: Name too short
*/

func CreateMajor(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.JSON(createMajorResponse{Error: 1})
	}
	request := createMajorRequest{}
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(createMajorResponse{Error: 2})
	}
	if len(request.Name) < 3 {
		return c.JSON(createMajorResponse{Error: 3})
	}
	_, err := db.Db.Exec("insert into nusbiam.Majors (major_id, major_name) VALUE (?,?)",
		uuid.New().String(), request.Name,
	)
	if err!=nil{
		return c.SendStatus(500)
	}
	return c.JSON(createMajorResponse{Error: -1})
}
