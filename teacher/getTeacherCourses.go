package teacher


import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	db "nusbi-server/config"
)

type getTeacherCourseResponse struct{
	ClassName string
	CourseName string
	ClassID string
	SCU int
}

func GetTeacherCourse(c *fiber.Ctx) error{
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "t" {
		return c.SendStatus(403)
	}
	id:=c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}
	res, err := db.Db.Query(
		"select distinct c.class_id,c.class_name,cs.course_name,cs.scu from Class c, Courses cs, Lecturers l where cs.lecturer_id = (select lecturer_id from Lecturers where user_id = ?) and c.course_id = cs.course_id",
		id,
	)
	if err != nil {
		return c.SendStatus(500)
	}
	if res.Err() != nil {
		return c.SendStatus(500)
	}
	var result []getTeacherCourseResponse
	for res.Next() {
		var temp getTeacherCourseResponse
		err = res.Scan(&temp.ClassID,&temp.ClassName, &temp.CourseName, &temp.SCU)
		if err != nil {
			return c.SendStatus(500)
		}
		result = append(result, temp)
	}
	return c.JSON(fiber.Map{
		"TeacherCourseData": result,
	})
}