package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"mynamebvh.com/blog/controllers"
	middleware "mynamebvh.com/blog/middlewares"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/check", middleware.Hello, controllers.CheckHeath)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", controllers.Login)
	auth.Post("/signup", controllers.SignUp)
	

	//User
	user := api.Group("/user")
	user.Get("/:id", middleware.Protected(),controllers.GetUser)
	user.Put("/", middleware.Protected(), controllers.UpdateUser)
	user.Delete("/", middleware.Protected(), controllers.DeleteUser)
	
}
