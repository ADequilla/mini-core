package routes

import (
	"mini-core/middleware/go-utils/database"
	"mini-core/middleware/utils"
	"mini-core/modules/search/models/request"
	"mini-core/modules/search/models/response"

	"github.com/gofiber/fiber/v2"
)

func ViewClientAccount(c *fiber.Ctx) error {
	var view request.ViewClientAccount
	var result []response.GetAccountsModel

	if err := utils.BodyParser(c, &view); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"success": false,
			"error":   err.Error(),
		})
	}

	if err := database.DBConn.Raw("SELECT * FROM ewallet_web.view_client_account(?,?,?)", view.Id_input, view.Cid_input, view.Mobile_input).Find(&result).Error; err != nil {
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
