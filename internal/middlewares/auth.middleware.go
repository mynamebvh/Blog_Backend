package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"mynamebvh.com/blog/config"
)

func Protected() fiber.Handler{
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.GetEnv("JWT_SECRET")),
		ErrorHandler: jwtError,
		AuthScheme: "Bearer",
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	fmt.Println(err)
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Lỗi xác thực", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Token hết hạn", "data": nil})
}