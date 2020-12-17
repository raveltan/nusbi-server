package courses

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	db "nusbi-server/config"
)

type getCourseClassResponse struct{
	ClassName string
	CourseName string
	ClassID string
}

func GetCourseClass(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.SendStatus(403)
	}
	res,err:=db.Db.Query("select class_id,course_name,class_name from Courses, Class")
	if err!=nil{
		return c.SendStatus(500)
	}
	if res.Err()!=nil{
		return c.SendStatus(500)
	}
	var result []getCourseClassResponse
	for res.Next(){
		var temp getCourseClassResponse
		err = res.Scan(&temp.ClassID,&temp.CourseName,&temp.ClassName)
		if err!=nil{
			return c.SendStatus(500)
		}
		result = append(result, temp)
	}
	return c.JSON(fiber.Map{
		"data":result,
	})
}