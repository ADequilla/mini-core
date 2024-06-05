package routes

import (
	"fmt"
	"time"

	"mini-core/middleware/go-utils/database"
	"mini-core/middleware/go-utils/passwordHashing"
	"mini-core/modules/create_account/models/errors"
	"mini-core/modules/create_account/models/request"
	"mini-core/modules/create_account/models/response"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Getusertype(usrtype string) string {
	switch usrtype {
	case "MFI-ADMIN":
		return "true"
	default:
		return "flase"
	}
}

func LoginUser(c *fiber.Ctx) error {
	currentTime := time.Now()
	formatTime := currentTime.Format("2006-01-02 15:04:05")
	reqbody := request.RequestBodyStruct{}
	respbody := response.ResponseBodyStruct{}

	if bodyErr := c.BodyParser(&reqbody); bodyErr != nil {
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

	if getErr := database.DBConn.Debug().Raw(`SELECT * FROM ewallet_web.users WHERE user_name = ?`, reqbody.UserName).Scan(&respbody).Error; getErr != nil {
		return c.JSON(errors.ResponseModel{
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
		return c.JSON(errors.ResponseModel{
			RetCode: "400",
			Message: "Already Login",
		})
	}

	if respbody.UserStatus != "Approved" {
		return c.JSON(errors.ResponseModel{
			RetCode: "400",
			Message: "Contact System Administration for approval.",
		})
	}

	// Check if the account is locked
	if !respbody.LockoutTime.IsZero() && respbody.LockoutTime.After(time.Now()) {
		return c.JSON(errors.ResponseModel{
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

		return c.JSON(errors.ResponseModel{
			RetCode: "401",
			Message: "Unauthorized",
			Data: errors.ErrorModel{
				Message:   "Invalid Credentials",
				IsSuccess: false,
			},
		})
	}

	adminclaims := Getusertype(respbody.UserPosition)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = reqbody.UserName
	claims["insti"] = respbody.InstiCode
	claims["role"] = respbody.UserPosition
	claims["admin"] = adminclaims
	claims["exp_stat"] = false
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

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

	// Reset failed attempts on successful login
	if updateErr := database.DBConn.Exec(`UPDATE ewallet_web.users SET login_attempts = 0,is_login = 'true',last_login_date = ?, token = ? WHERE user_name = ?`, formatTime, t, reqbody.UserName).Error; updateErr != nil {
		return c.JSON(errors.ResponseModel{
			RetCode: "400",
			Message: "Bad Request",
			Data: errors.ErrorModel{
				Message:   "Update Failed",
				IsSuccess: false,
				Error:     updateErr,
			},
		})
	}

	returnData := response.ResposBodyStruct{
		UserName:     respbody.UserName,
		UserEmail:    respbody.UserEmail,
		UserPosition: respbody.UserPosition,
		UserPhone:    respbody.UserPhone,
		InstiCode:    respbody.InstiCode,
		LastName:     respbody.LastName,
		GivenName:    respbody.GivenName,
		MiddleName:   respbody.MiddleName,
	}

	return c.JSON(errors.LoginResponseModel{
		RetCode: "201",
		Message: "Login Successful",
		Token:   t,
		Data:    returnData,
	})
}
