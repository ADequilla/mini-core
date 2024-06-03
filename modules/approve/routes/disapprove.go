package routes

import (
	"mini-core/middleware/go-utils/database"
	"mini-core/middleware/utils"
	"mini-core/modules/approve/models/request"
	"mini-core/modules/approve/models/response"

	"github.com/gofiber/fiber/v2"
)

func DisapproveClients(c *fiber.Ctx) error {
	var request request.ApproveClients
	var result []response.ApproveClientModel

	if err := utils.BodyParser(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"success": false,
			"error":   err.Error(),
		})
	}

	if err := database.DBConn.Raw("SELECT * FROM ewallet_web.disapprove_client(?,?,?)", request.Id_input, request.Cid_input, request.Mobile_input).Find(&result).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result,
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"response": result,
		"success":  true,
	})
}
