package routes

import (
	"mini-core/middleware/go-utils/database"
	"mini-core/modules/create_account/models/response"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var result []response.GetUserModel

	if err := database.DBConn.Raw("SELECT * FROM ewallet_web.get_registred_user()").Find(&result).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch User list",
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"response": result,
		"success":  true,
	})
}
