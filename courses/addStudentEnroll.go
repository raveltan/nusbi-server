package courses

import "github.com/gofiber/fiber/v2"

type studentEnrollRequest struct{
	Username string
	ClassID string
}

func AddStudentEnroll (c *fiber.Ctx) error{
	return nil
}