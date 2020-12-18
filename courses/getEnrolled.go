package courses

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	db "nusbi-server/config"
)

type getEnrolledResponse struct {
	ClassName  string
	CourseName string
	ClassID    string
}

func GetEnrolled(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.SendStatus(403)
	}
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}
	res, err := db.Db.Query(
		"select e.class_id,class_name,course_name from Enrolled_Courses e,Class c,Courses cs where c.class_id = e.class_id and c.course_id = cs.course_id and e.student_id = (select student_id from Students where user_id = ?)",
		id,
	)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	if res.Err() != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	var result []getEnrolledResponse
	for res.Next() {
		var temp getEnrolledResponse
		err = res.Scan(&temp.ClassID, &temp.ClassName, &temp.CourseName)
		if err != nil {
			log.Println(err)
			return c.SendStatus(500)
		}
		result = append(result, temp)
	}
	return c.JSON(fiber.Map{
		"data": result,
	})
}
