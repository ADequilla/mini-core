package routes

import (
	"mini-core/middleware/go-utils/database"
	"mini-core/middleware/utils"
	"mini-core/modules/update/models/request"
	"mini-core/modules/update/models/response"

	"github.com/gofiber/fiber/v2"
)

func UpdateAccount(c *fiber.Ctx) error {
	var request request.UpdateAccountsModel
	var result []response.UpdateAccount

	if err := utils.BodyParser(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"success": false,
			"error":   err.Error(),
		})
	}

	if err := database.DBConn.Raw("SELECT * FROM ewallet_web.edit_account(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", request.SetAccountNumber, request.SetAcc, request.SetAcctType, request.SetAccDesc, request.SetDOpen, request.SetStatusDesc, request.SetIIID, request.SetStatus, request.SetTitle, request.SetClassification, request.SetSubClassification, request.SetDoEntry, request.SetDoRecognized, request.SetDoResigned, request.SetInstiCode, request.SetBranchCode, request.SetUnitCode, request.SetCenterCode, request.SetUUID, request.SetCID, request.SetAreaCode, request.SetArea, request.SetBalance, request.SetWithdrawable, request.SetLedgerBalance).Find(&result).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update Account",
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"response": result,
		"success":  true,
	})
}
