package class

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	db "nusbi-server/config"
)

type getClassResponse struct{
	ClassName string
	Batch int
	ClassID string
}

func GetClasses(c *fiber.Ctx) error {
	id:=c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.SendStatus(403)
	}
	res,err:=db.Db.Query("select class_id,class_name,batch from Class where course_id = ?",id)
	if err!=nil{
		return c.SendStatus(500)
	}
	if res.Err()!=nil{
		return c.SendStatus(500)
	}
	result := []getClassResponse{}
	for res.Next(){
		var temp getClassResponse
		err = res.Scan(&temp.ClassID,&temp.ClassName,&temp.Batch)
		if err!=nil{
			return c.SendStatus(500)
		}
		result = append(result, temp)
	}
	return c.JSON(fiber.Map{
		"data":result,
	})
}
