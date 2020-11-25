package courses

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	db "nusbi-server/config"
)

type getLecturerResponse struct {
	Data []lecturer
}

type lecturer struct {
	Name string
	ID   string
}

func GetLecturer(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.SendStatus(400)
	}
	res, err := db.Db.Query("select first_name,last_name,lecturer_id from Lecturers")
	if err != nil {
		return c.SendStatus(500)
	}
	var result getLecturerResponse
	for res.Next() {
		var data lecturer
		var firstName string
		var lastName string

		err = res.Scan(&firstName, &lastName, &data.ID)
		if err != nil {
			return c.SendStatus(500)
		}
		data.Name = firstName + " " + lastName
		result.Data = append(result.Data, data)
	}
	if err = res.Err(); err != nil {
		return c.SendStatus(500)
	}
	return c.JSON(result)
}
