package courses

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	db "nusbi-server/config"
)

type getCourseResponse struct {
	Data []course
}

type course struct {
	CourseID    string
	Name        string
	Scu         int
	TeacherName string
}

func GetCourse(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.SendStatus(403)
	}
	res,err:=db.Db.Query("select concat(l.first_name,' ',l.last_name) as teacher_name,c.course_name,c.scu,c.course_id from Lecturers l , Courses c where c.lecturer_id = l.lecturer_id")
	if err!=nil{
		return c.SendStatus(500)
	}
	if res.Err()!=nil{
		return c.SendStatus(500)
	}
	var result getCourseResponse
	for res.Next(){
		var temp course
		err = res.Scan(&temp.TeacherName,&temp.Name,&temp.Scu,&temp.CourseID)
		if err!=nil{
			return c.SendStatus(500)
		}
		result.Data = append(result.Data, temp)
	}
	return c.JSON(result)
}
