package routes

import (
	"mini-core/middleware/go-utils/database"
	"mini-core/modules/select/models/response"

	"github.com/gofiber/fiber/v2"
)

func GetAccount(c *fiber.Ctx) error {
	// var page request.Accounts
	var result []response.GetAccountsModel

	// if err := utils.BodyParser(c, &page); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "Failed to parse request body",
	// 		"success": false,
	// 		"error":   err.Error(),
	// 	})
	// }

	if err := database.DBConn.Raw("SELECT * FROM ewallet_web.get_account()").Find(&result).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch Account list",
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"response": result,
		"success":  true,
	})
}
