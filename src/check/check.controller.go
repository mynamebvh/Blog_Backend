package controllers

import "github.com/gofiber/fiber/v2"

func CheckHeath(c *fiber.Ctx) error {
	return c.SendString("hello")
}