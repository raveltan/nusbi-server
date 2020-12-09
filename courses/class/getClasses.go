package class

import (
	"github.com/gofiber/fiber/v2"
)


func getClasses(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
