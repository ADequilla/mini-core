package middleware

import (
	"mini-core/middleware/go-utils/database"
	"mini-core/modules/approve/models/errors"
	"mini-core/modules/batch_upload/models/response"
	createaccount "mini-core/modules/create_account"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {

	reqBody := createaccount.RequestBodyStruct{}
	respBody := createaccount.RequestBodyStruct{}

	if bodyErr := c.BodyParser(&reqBody); bodyErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad Request",
			Data: errors.ErrorModel{
				Message:   "Parsing Failed",
				IsSuccess: false,
				Error:     bodyErr,
			},
		})
	}
	if len(reqBody.UserName) == 0 {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad Request",
			Data: errors.ErrorModel{
				Message:   "Username is empty",
				IsSuccess: false,
			},
		})
	}

	if searchErr := database.DBConn.Raw(`SELECT * FROM ewallet_web.users WHERE user_name = ? `, reqBody.UserName).Scan(&respBody).Error; searchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad Request",
			Data: errors.ErrorModel{
				Message:   "Fetching Failed",
				IsSuccess: false,
				Error:     searchErr,
			},
		})
	}
	if reqBody.UserName != respBody.UserName {
		return c.JSON(response.ResponseModel{
			RetCode: "404",
			Message: "Not Found",
			Data: errors.ErrorModel{
				Message:   "Error: Username Not Found",
				IsSuccess: false,
			},
		})
	}
	if updateErr := database.DBConn.Exec(`UPDATE ewallet_web.users SET is_login = false WHERE user_name = ?`, reqBody.UserName).Error; updateErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad Request",
			Data: errors.ErrorModel{
				Message:   "Failed to Update Data",
				IsSuccess: false,
				Error:     updateErr,
			},
		})
	}

	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "Logout, Successful",
	})
}
