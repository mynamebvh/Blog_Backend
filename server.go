package main

import (
	"github.com/gofiber/fiber/v2"
	"mynamebvh.com/blog/database"
	"mynamebvh.com/blog/routers"
)

func main() {
  app := fiber.New()

	database.ConnectDB()
	
  routers.SetupRoutes(app)

  app.Listen(":3000")
}