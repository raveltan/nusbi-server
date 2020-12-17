package teacher

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	db "nusbi-server/config"
)

type getTeacherProfileResponse struct{
	FirstName string
	LastName string
	DOB string
	Gender string
	Email string
}

func GetTeacherProfile(c *fiber.Ctx) error {
	id:=c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "t" {
		return c.SendStatus(403)
	}
	sqlResult, err := db.Db.Query(
		"select first_name, last_name, gender, dob, email from Lecturers where user_id = ?",id)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	if sqlResult.Err() != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	var r getTeacherProfileResponse
	f := false
	for sqlResult.Next() {
		f = true
		err = sqlResult.Scan(&r.FirstName,&r.LastName,&r.Gender,&r.DOB,&r.Email)
		if err != nil {
			log.Println(err)
			return c.SendStatus(500)
		}
	}
	if !f {
		return c.SendStatus(404)
	}
	return c.JSON(r)
}