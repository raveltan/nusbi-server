package teacher

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	db "nusbi-server/config"
)

type getClassStudentResponse struct{
	StudentID string
	Name string
	MidScore int
	FinalScore int
}

func GetClassStudent(c *fiber.Ctx)error{
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "t" {
		return c.SendStatus(403)
	}
	id:=c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}
	res, err := db.Db.Query(
		"select distinct s.student_id,concat(s.first_name,' ',s.last_name) as name,e.mid_score,e.final_score as Name from Enrolled_Courses e, Students s where e.class_id = ? and e.student_id=s.student_id",
		id,
	)
	if err != nil {
		return c.SendStatus(500)
	}
	if res.Err() != nil {
		return c.SendStatus(500)
	}
	var result []getClassStudentResponse
	for res.Next() {
		var temp getClassStudentResponse
		err = res.Scan(&temp.StudentID,&temp.Name, &temp.MidScore, &temp.FinalScore)
		if err != nil {
			return c.SendStatus(500)
		}
		result = append(result, temp)
	}
	return c.JSON(fiber.Map{
		"ClassStudentData": result,
	})
}
