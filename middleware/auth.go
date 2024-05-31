package middleware

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte("secret"),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing or malformed JWT",
			"data":    nil})

	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid or expired JWT",
			"data":    nil})
	}
}
func decodeJWT(jwt string) (map[string]interface{}, error) {
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("Invalid JWT format")
	}

	payloadJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("Error decoding payload: %s", err.Error())
	}

	var payloadMap map[string]interface{}
	if err := json.Unmarshal(payloadJSON, &payloadMap); err != nil {
		return nil, fmt.Errorf("Error decoding/verifying payload: %s", err.Error())
	}

	return payloadMap, nil
}

func JWTMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		parts := strings.Split(token, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token format",
			})
		}

		jwtToken := parts[1]

		payloadMap, err := decodeJWT(jwtToken)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": fmt.Sprintf("Error decoding JWT: %s", err.Error()),
			})
		}

		fmt.Println("Decoded Payload:", payloadMap)

		c.Locals("payload", payloadMap)

		return c.Next()
	}
}

func HeadersLoginValidator() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		expectedHeaders := []string{"deviceId", "deviceModel", "fcmToken", "osVersion", "appVersion"}

		for _, header := range expectedHeaders {
			if c.Get(header) == "" {
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{
					"error": fmt.Sprintf("Missing header: %s", header),
				})
			}
		}

		return c.Next()
	}
}

func HeaderResponse() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Set("X-Frame-Options", "DENY")
		return c.Next()
	}
}
