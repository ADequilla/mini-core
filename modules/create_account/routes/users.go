package routes

import (
	"mini-core/middleware/go-utils/database"
	"mini-core/middleware/go-utils/passwordHashing"
	"mini-core/modules/create_account/models/errors"
	"mini-core/modules/create_account/models/request"
	"mini-core/modules/create_account/models/response"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RegisterNewUser(c *fiber.Ctx) error {
	currentTime := time.Now()
	formatTime := currentTime.Format("2006-01-02 15:04:05")
	reqBody := request.RequestBodyStruct{}
	respBody := response.ResponseBodyStruct{}

	if bodyErr := c.BodyParser(&reqBody); bodyErr != nil {
		return c.JSON(errors.ResponseModel{
			RetCode: "400",
			Message: "Body Parsing Failed",
			Data: errors.ErrorModel{
				Message:   "Parsing Failed",
				IsSuccess: false,
				Error:     bodyErr,
			},
		})
	}
	reqBody.UserPhone = NormalizePhoneNumber(c, reqBody.UserPhone)
	if len(reqBody.UserPhone) != 11 {
		return c.JSON(errors.ResponseModel{
			RetCode: "400",
			Message: "BadRequest",
			Data: errors.ErrorModel{
				Message:   reqBody.UserPhone,
				IsSuccess: false,
				Error:     nil,
			},
		})
	}
	hashPassword, hashErr := passwordHashing.HashPassword(reqBody.UserPassword)
	if hashErr != nil {
		return c.JSON(errors.ResponseModel{
			RetCode: "400",
			Message: "BadRequest",
			Data: errors.ErrorModel{
				Message:   "Error: Failed to Harsh Password",
				IsSuccess: false,
				Error:     hashErr,
			},
		})
	}
	reqBody.UserName = ValidateUsername(reqBody.UserName)
	if len(reqBody.UserName) == 0 {
		return c.JSON(errors.ResponseModel{
			RetCode: "400",
			Message: "BadRequest",
			Data: errors.ErrorModel{
				Message:   "Error: Username is Empty",
				IsSuccess: false,
				Error:     nil,
			},
		})
	}
	if reqBody.UserName == "Username already exists" {
		return c.JSON(errors.ResponseModel{
			RetCode: "400",
			Message: "Bad Request",
			Data: errors.ErrorModel{
				Message:   "The Username is already Exist",
				IsSuccess: false,
			},
		})
	}
	role := GetRoleByInstiCode(reqBody.InstiCode, reqBody.UserType)

	if insertErr := database.DBConn.Debug().Raw(`SELECT * FROM ewallet_web.func_insert_users(?,?,?,?,?,?,?,?,?,?,?,?)`, hashPassword, reqBody.UserName, reqBody.UserEmail, reqBody.UserPhone, formatTime, reqBody.UserType, "Pending", role, reqBody.InstiCode, reqBody.LastName, reqBody.GivenName, reqBody.MiddleName).Scan(&respBody).Error; insertErr != nil {
		return c.JSON(errors.ResponseModel{
			RetCode: "400",
			Message: "Bad Request",
			Data: errors.ErrorModel{
				Message:   "Error: Inserting Failed",
				IsSuccess: false,
				Error:     insertErr,
			},
		})
	}
	returnData := response.ResponseBodyStruct{
		UserName:     reqBody.UserName,
		CreatedDate:  formatTime,
		UserEmail:    reqBody.UserEmail,
		UserPasswd:   reqBody.UserPassword,
		UserStatus:   "Pending",
		UserPosition: role,
		UserPhone:    reqBody.UserPhone,
		InstiCode:    reqBody.InstiCode,
		LastName:     reqBody.LastName,
		GivenName:    reqBody.GivenName,
		MiddleName:   reqBody.MiddleName,
	}
	return c.JSON(errors.ResponseModel{
		RetCode: "201",
		Message: "Success Created",
		Data:    returnData,
	})
}

// Check the phone number format
func NormalizePhoneNumber(c *fiber.Ctx, phonenumber string) string {
	var normalizedPhonenumber string

	if len(phonenumber) == 0 {
		normalizedPhonenumber = "Phone number is empty"
		return normalizedPhonenumber
	}

	if phonenumber[0:1] == "0" {
		normalizedPhonenumber = phonenumber
	} else if phonenumber[0:1] == "6" {
		normalizedPhonenumber = "0" + phonenumber[2:12]
	} else if phonenumber[0:1] == "+" || phonenumber[0:1] == " " {
		normalizedPhonenumber = "0" + phonenumber[3:13]
	} else if phonenumber[0:1] == "9" {
		normalizedPhonenumber = "0" + phonenumber
	} else {
		normalizedPhonenumber = "Invalid number format!"
	}

	if len(normalizedPhonenumber) != 11 {
		normalizedPhonenumber = "The phonenumber is invalid"
		return normalizedPhonenumber
	} else {
		return normalizedPhonenumber
	}
}

// Validate the UserName
func ValidateUsername(username string) string {
	if len(username) < 3 || len(username) > 20 {
		return "Username is too short or too long"
	}
	validateName := regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(username)
	if !validateName {
		return "Username is Invalid"
	}

	var count int64
	if err := database.DBConn.Raw(`SELECT count(*) FROM ewallet_web.users WHERE user_name = ?`, username).Count(&count).Error; err != nil {
		return "Error occurred while checking username"
	}
	if count > 0 {
		return "Username already exists"
	}

	return username
}

// Get insticode based on the user input
func GetRoleByInstiCode(insticode string, usrtype string) string {
	var role string

	switch insticode {
	case "001":
		role = "FDS"
	default:
		role = "MFI"
	}
	switch usrtype {
	case "Admin":
		return role + "-ADMIN"
	default:
		return role + "-USER"
	}
}
