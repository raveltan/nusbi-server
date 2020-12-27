package schedule

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	db "nusbi-server/config"
)


type getScheduleResponse struct{
	ScheduleID string
	Date string
}

func GetSchedule(c *fiber.Ctx) error {
	id:=c.Params("id")
	if id == "" {
		return c.SendStatus(400)
	}
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.SendStatus(403)
	}
	res,err:=db.Db.Query("select distinct date_time,schedule_id from Schedules where class_id = ? order by date_time desc",id)
	if err!=nil{
		return c.SendStatus(500)
	}
	if res.Err()!=nil{
		return c.SendStatus(500)
	}
	result := []getScheduleResponse{}
	for res.Next(){
		var temp getScheduleResponse
		err = res.Scan(&temp.Date,&temp.ScheduleID)
		if err!=nil{
			return c.SendStatus(500)
		}
		result = append(result, temp)
	}
	return c.JSON(fiber.Map{
		"data":result,
	})
}
