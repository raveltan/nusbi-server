package courses

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	db "nusbi-server/config"
)

type getCourseClassResponse struct {
	ClassName  string
	CourseName string
	ClassID    string
}

func GetCourseClass(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.SendStatus(403)
	}
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}
	res, err := db.Db.Query(
		"select class_id,course_name,class_name from Courses, Class where Class.course_id = Courses.course_id and class_id not in (select class_id from Enrolled_Courses where student_id = (select student_id from Students where user_id = ?)) and Courses.course_id not in (select c.course_id from Enrolled_Courses e,Class c where e.class_id = c.class_id and e.student_id = (select student_id from Students where user_id = ?))",
		id,id,
	)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	if res.Err() != nil {
		return c.SendStatus(500)
	}
	result := []getCourseClassResponse{}
	for res.Next() {
		var temp getCourseClassResponse
		err = res.Scan(&temp.ClassID, &temp.CourseName, &temp.ClassName)
		if err != nil {
			return c.SendStatus(500)
		}
		result = append(result, temp)
	}
	return c.JSON(fiber.Map{
		"data": result,
	})
}
