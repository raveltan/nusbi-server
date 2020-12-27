package student

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	db "nusbi-server/config"
)

type studentAbsentRecordResponse struct{
	ClassName string
	CourseName string
	Absent int
}

func GetStudentAbsentRecord(c *fiber.Ctx)error{
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "s" {
		return c.SendStatus(403)
	}
	id:=c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}
	res, err := db.Db.Query(
		"select c.class_name,x.course_name,count(x.course_name) as absent from Absence a,Schedules s,Class c,Courses x where a.schedule_id = s.schedule_id and c.class_id = s.class_id and x.course_id = c.course_id and a.student_id = (select student_id from Students where user_id = ? limit 1) group by x.course_name, c.class_name",
		id,
	)
	if err != nil {
		return c.SendStatus(500)
	}
	if res.Err() != nil {
		return c.SendStatus(500)
	}
	var result []studentAbsentRecordResponse
	for res.Next() {
		var temp studentAbsentRecordResponse
		err = res.Scan(&temp.ClassName, &temp.CourseName, &temp.Absent)
		if err != nil {
			return c.SendStatus(500)
		}
		result = append(result, temp)
	}
	return c.JSON(fiber.Map{
		"GetStudentAbsentRecordData": result,
	})
}