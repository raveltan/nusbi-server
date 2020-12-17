package student

import (
	"database/sql"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	db "nusbi-server/config"
)

type getProfileResponse struct{
	FirstName string
	LastName string
	Gender string
	DOB string
	Email string
	GPA sql.NullInt32
	SCU sql.NullInt32
	batch int
	Major string
}

func GetStudentProfile(c *fiber.Ctx) error {
	id:=c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "s" {
		return c.SendStatus(403)
	}
	sqlResult, err := db.Db.Query(
		"select first_name,last_name,gender,dob,email,gpa,scu,batch,m.major_name from Students s,Majors m where user_id = ? and m.major_id = s.major_id",id)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	if sqlResult.Err() != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	var r getProfileResponse
	f := false
	for sqlResult.Next() {
		f = true
		err = sqlResult.Scan(&r.FirstName,&r.LastName,&r.Gender,&r.DOB,&r.Email,&r.GPA,&r.SCU,&r.batch,&r.Major)
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