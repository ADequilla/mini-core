package utils

import "github.com/gofiber/fiber/v2"

func BodyParser(c *fiber.Ctx, in interface{}) error {
	err := c.BodyParser(in)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"Message": "Server error",
			"Error":   err,
		})
	}
	return err
}
