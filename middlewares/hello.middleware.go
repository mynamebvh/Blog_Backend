package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	if c.Query("q") == "hoangdz" {
		return c.Next()
	}	else {
		return c.Status(404).JSON(fiber.Map{"err": "khong dz"})
	}
}
