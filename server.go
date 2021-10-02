package main

import (
	"github.com/gofiber/fiber/v2"
	config "mynamebvh.com/blog/config"
	db "mynamebvh.com/blog/db"
)

func main() {
  app := fiber.New()

	db.ConnectDB()
	
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString(config.GetEnv("JWT_SECRET"))
  })

  app.Listen(":3000")
}