package major

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	db "nusbi-server/config"
)

type major struct {
	ID   string
	Name string
}

func GetMajor(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.SendStatus(403)
	}
	sqlResult, err := db.Db.Query("select major_id,major_name from nusbiam.Majors")
	if err != nil {
		return c.SendStatus(500)
	}
	if sqlResult.Err() != nil {
		return c.SendStatus(500)
	}
	result := []major{}
	for sqlResult.Next() {
		var temp major
		err = sqlResult.Scan(&temp.ID, &temp.Name)
		if err != nil {
			return c.SendStatus(500)
		}
		result = append(result, temp)
	}
	return c.JSON(fiber.Map{
		"data": result,
	})
}
