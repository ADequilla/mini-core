package routes

import (
	"mini-core/middleware/go-utils/database"
	"mini-core/modules/select/models/response"

	"github.com/gofiber/fiber/v2"
)

func GetClient(c *fiber.Ctx) error {
	var result []response.GetClientModel

	if err := database.DBConn.Raw("SELECT * FROM ewallet_web.get_client()").Find(&result).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch Client list",
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"response": result,
		"success":  true,
	})
}
