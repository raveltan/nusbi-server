package student

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	db "nusbi-server/config"
)

type getStudentCourseResponse struct{
	Name string
	TeacherName string
	SCU int
}

func GetStudentCourse(c *fiber.Ctx) error{
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "s" {
		return c.SendStatus(403)
	}
	id:=c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}
	res, err := db.Db.Query(
		"select concat(cs.course_name,' - ',c.class_name) as class,concat(l.first_name,' ',l.last_name) as teacher,cs.scu from Enrolled_Courses e,Class c, Courses cs, Lecturers l where student_id = (select student_id from Students where user_id = ?) and e.class_id = c.class_id and c.course_id = cs.course_id and l.lecturer_id = cs.lecturer_id",
		id,
	)
	if err != nil {
		return c.SendStatus(500)
	}
	if res.Err() != nil {
		return c.SendStatus(500)
	}
	var result []getStudentCourseResponse
	for res.Next() {
		var temp getStudentCourseResponse
		err = res.Scan(&temp.Name, &temp.TeacherName, &temp.SCU)
		if err != nil {
			return c.SendStatus(500)
		}
		result = append(result, temp)
	}
	return c.JSON(fiber.Map{
		"data": result,
	})
}