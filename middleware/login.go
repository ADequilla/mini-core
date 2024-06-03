package middleware

import (
	"fmt"
	"time"

	"mini-core/middleware/go-utils/database"
	"mini-core/middleware/go-utils/passwordHashing"
	"mini-core/modules/approve/models/errors"
	"mini-core/modules/approve/models/response"
	createaccount "mini-core/modules/create_account"

	"github.com/gofiber/fiber/v2"
)

func LoginUser(c *fiber.Ctx) error {
	currentTime := time.Now()
	formatTime := currentTime.Format("2006-01-02 15:04:05")
	reqbody := createaccount.RequestBodyStruct{}
	respbody := createaccount.ResponseBodyStruct{}

	if bodyErr := c.BodyParser(&reqbody); bodyErr != nil {
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

	if getErr := database.DBConn.Debug().Raw(`SELECT * FROM ewallet_web.users WHERE user_name = ?`, reqbody.UserName).Scan(&respbody).Error; getErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad Request",
			Data: errors.ErrorModel{
				Message:   "Fetching Failed",
				IsSuccess: false,
				Error:     getErr,
			},
		})
	}
	if respbody.IsLogin == "true" {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Already Login",
		})
	}

	// Check if the account is locked
	if !respbody.LockoutTime.IsZero() && respbody.LockoutTime.After(time.Now()) {
		return c.JSON(response.ResponseModel{
			RetCode: "403",
			Message: "Account is locked. Try again later.",
			Data: errors.ErrorModel{
				Message:   "Account is locked",
				IsSuccess: false,
			},
		})
	}

	// Check if the provided password matches the hashed password from the database
	isPasswordValid := passwordHashing.CheckPasswordHash(reqbody.UserPassword, respbody.UserPasswd)
	if !isPasswordValid {
		respbody.FailedAttempts++
		fmt.Println("Attempts :", respbody.FailedAttempts)
		if respbody.FailedAttempts >= 3 {
			respbody.LockoutTime = time.Now().Add(15 * time.Minute) // Lock account for 15 minutes
			respbody.FailedAttempts = 0                             // Reset failed attempts after locking
		}
		database.DBConn.Debug().Exec(`UPDATE ewallet_web.users SET login_attempts = ? WHERE user_name = ?`, respbody.FailedAttempts, reqbody.UserName)

		return c.JSON(response.ResponseModel{
			RetCode: "401",
			Message: "Unauthorized",
			Data: errors.ErrorModel{
				Message:   "Invalid Credentials",
				IsSuccess: false,
			},
		})
	}
	// Reset failed attempts on successful login
	if updateErr := database.DBConn.Exec(`UPDATE ewallet_web.users SET login_attempts = 0,is_login = 'true',last_login_date = ? WHERE user_name = ?`, formatTime, reqbody.UserName).Error; updateErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad Request",
			Data: errors.ErrorModel{
				Message:   "Update Failed",
				IsSuccess: false,
				Error:     updateErr,
			},
		})
	}

	return c.JSON(response.ResponseModel{
		RetCode: "201",
		Message: "Login Successful",
	})
}
