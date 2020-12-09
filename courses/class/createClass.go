package class

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	db "nusbi-server/config"
)

type classRequest struct{
	ClassName string
	Batch string
	CourseID string
}


func CreateClass(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.SendStatus(403)
	}
	var request classRequest

	err := c.BodyParser(&request)
	if err != nil {
		return c.SendStatus(400)
	}
	_, err = db.Db.Exec("insert into nusbiam.Class (class_id, class_name, course_id, batch) VALUE (?,?,?,?)",
		uuid.New().String(),request.ClassName,request.CourseID,request.Batch)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}
