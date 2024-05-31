package middleware

// import (
// 	"cas/struct/request"
// 	"errors"
// 	"fmt"
// 	"os"
// 	"time"
// 	"github.com/gofiber/fiber/v2"
// 	jwtware "github.com/gofiber/jwt/v2"
// 	"github.com/golang-jwt/jwt"
// 	"github.com/joho/godotenv"
// )

// // Protected protect routes
// func Protected() func(*fiber.Ctx) error {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		fmt.Printf("error loading .env: %s\n", err.Error())
// 		os.Exit(1)
// 	}

// 	skey := os.Getenv("SECRET_KEY")
// 	return jwtware.New(jwtware.Config{
// 		SigningKey:   []byte(skey),
// 		ErrorHandler: jwtError,
// 	})
// }

// func jwtError(c *fiber.Ctx, err error) error {
// 	if err.Error() == "Missing or malformed JWT" {
// 		c.Status(fiber.StatusBadRequest)
// 		return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})

// 	} else {
// 		c.Status(fiber.StatusUnauthorized)
// 		return c.JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
// 	}
// }

// func GenerateTokenClaim(req request.GenTokenRequest) (string, error) {
// 	claims := jwt.MapClaims{
// 		"appId":     req.Appid,
// 		"username":  req.Username,
// 		"firstname": req.Firstname,
// 		"lastname":  req.Lastname,
// 		"exp":       time.Now().Add(time.Minute * 15).Unix(),
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		fmt.Printf("error loading .env: %s\n", err.Error())
// 		os.Exit(1)
// 	}

// 	skey := os.Getenv("SECRET_KEY")
// 	// Generate encoded token and send it as response.
// 	t, err := token.SignedString([]byte(skey))
// 	if err != nil {
// 		return "", errors.New("Cant generate token")
// 	}
// 	return t, nil
// }
