package routes

import (
	"mini-core/middleware/go-utils/database"
	"mini-core/middleware/utils"
	"mini-core/modules/approve/models/request"
	"mini-core/modules/approve/models/response"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func DisapproveClients(c *fiber.Ctx) error {
	var request []request.ApproveClients
	var result []response.ApproveClientModel

	if err := utils.BodyParser(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"success": false,
			"error":   err.Error(),
		})
	}

	// Extract the id_input, cid_input, and mobile_input into separate slices
	var idInput []int64
	var cidInput []int
	var mobileInput []string

	for _, r := range request {
		idInput = append(idInput, int64(r.Id_input))
		cidInput = append(cidInput, r.Cid_input)
		mobileInput = append(mobileInput, r.Mobile_input)
	}

	if err := database.DBConn.Raw("SELECT * FROM ewallet_web.disapprove_client(?,?,?)", pq.Array(idInput), pq.Array(cidInput), pq.Array(mobileInput)).Find(&result).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to execute database query",
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"response": result,
		"success":  true,
	})
}
