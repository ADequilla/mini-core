package routes

import (
	"mini-core/middleware/go-utils/database"
	"mini-core/middleware/utils"
	"mini-core/modules/update/models/request"
	"mini-core/modules/update/models/response"

	"github.com/gofiber/fiber/v2"
)

func UpdateClient(c *fiber.Ctx) error {
	var request request.UpdateClientModel
	var result []response.UpdatetClient

	if err := utils.BodyParser(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"success": false,
			"error":   err.Error(),
		})
	}

	if err := database.DBConn.Raw("SELECT * FROM ewallet_web.edit_client(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", request.Birthday, request.Cid, request.Mobile, request.FirstName, request.LastName, request.MiddleName, request.MaidenFName, request.MaidenLName, request.MaidenMName, request.BirthPlace, request.Sex, request.CivilStatus, request.MemberMaidenFName, request.MemberMaidenLName, request.MemberMaidenMName, request.Email, request.InstiCode, request.UnitCode, request.CenterCode).Find(&result).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch Loan list",
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"response": result,
		"success":  true,
	})
}
