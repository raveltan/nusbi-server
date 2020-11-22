package auth

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	db "nusbi-server/config"
)

type User struct {
	Username string
	Role     string
}

func GetUserList(c *fiber.Ctx) error {
	if db.GetRoleFromToken(c.Locals("user").(*jwt.Token)) != "a" {
		return c.SendStatus(403)
	}
	var userList []User
	res, err := db.Db.Query("select user_id,role from nusbiam.Users")
	if err != nil {
		log.Println(err.Error())
		return c.SendStatus(500)
	}
	if res.Err() != nil {
		log.Println(err.Error())
		return c.SendStatus(500)
	}
	for res.Next() {
		var data User
		err = res.Scan(&data.Username,&data.Role)
		if err != nil {
			return c.SendStatus(500)
		}
		userList = append(userList, data)
	}
	return c.JSON(fiber.Map{
		"data": userList,
	})
}
