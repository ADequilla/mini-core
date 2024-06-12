package routes

import (
	"mini-core/middleware/go-utils/database"
	"mini-core/middleware/utils"
	"mini-core/modules/select/models/request"
	"mini-core/modules/select/models/response"

	"github.com/gofiber/fiber/v2"
)

func GetPendingClient(c *fiber.Ctx) error {
	var request request.Clients
	var result []response.GetClientModel

	if err := utils.BodyParser(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"success": false,
			"error":   err.Error(),
		})
	}

	if err := database.DBConn.Debug().Raw("SELECT * FROM ewallet_web.get_pending_client(?)", request.Insti_code_input).Find(&result).Error; err != nil {
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

func GetApprovedClient(c *fiber.Ctx) error {
	var request request.Clients
	var result []response.GetClientModel

	if err := utils.BodyParser(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"success": false,
			"error":   err.Error(),
		})
	}

	if err := database.DBConn.Debug().Raw("SELECT * FROM ewallet_web.get_approved_client(?)", request.Insti_code_input).Find(&result).Error; err != nil {
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

func GetDisapprovedClient(c *fiber.Ctx) error {
	var request request.Clients
	var result []response.GetClientModel

	if err := utils.BodyParser(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"success": false,
			"error":   err.Error(),
		})
	}

	if err := database.DBConn.Debug().Raw("SELECT * FROM ewallet_web.get_disapproved_client(?)", request.Insti_code_input).Find(&result).Error; err != nil {
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
