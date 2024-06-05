package routes

import (
	"mini-core/middleware/go-utils/database"
	"mini-core/modules/create_account/models/errors"
	"mini-core/modules/create_account/models/request"
	"mini-core/modules/create_account/models/response"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Logout(c *fiber.Ctx) error {
	currentTime := time.Now()
	formatTime := currentTime.Format("2006-01-02 15:04:05")
	reqBody := request.RequestBodyStruct{}
	respBody := response.ResponseBodyStruct{}

	payload, ok := c.Locals("payload").(map[string]interface{})
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Payload not found in context",
		})
	}

	identity, ok := payload["identity"].(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Identity not found or not a string",
		})
	}

	if bodyErr := c.BodyParser(&reqBody); bodyErr != nil {
		return c.JSON(errors.ResponseModel{
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
		return c.JSON(errors.ResponseModel{
			RetCode: "400",
			Message: "Bad Request",
			Data: errors.ErrorModel{
				Message:   "Username is empty",
				IsSuccess: false,
			},
		})
	}

	if searchErr := database.DBConn.Raw(`SELECT * FROM ewallet_web.users WHERE user_name = ? `, reqBody.UserName).Scan(&respBody).Error; searchErr != nil {
		return c.JSON(errors.ResponseModel{
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
		return c.JSON(errors.ResponseModel{
			RetCode: "404",
			Message: "Not Found",
			Data: errors.ErrorModel{
				Message:   "Error: Username Not Found",
				IsSuccess: false,
			},
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = identity
	claims["exp_stat"] = true
	claims["exp"] = formatTime

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal Server Error",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	if updateErr := database.DBConn.Exec(`UPDATE ewallet_web.users SET is_login = false, token = ? WHERE user_name = ?`, t, reqBody.UserName).Error; updateErr != nil {
		return c.JSON(errors.ResponseModel{
			RetCode: "400",
			Message: "Bad Request",
			Data: errors.ErrorModel{
				Message:   "Failed to Update Data",
				IsSuccess: false,
				Error:     updateErr,
			},
		})
	}

	return c.JSON(errors.LogoutResponseModel{
		RetCode: "200",
		Message: "Logout, Successful",
	})
}
