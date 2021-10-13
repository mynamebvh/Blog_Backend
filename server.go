package main

import (
	"github.com/gofiber/fiber/v2"
	db "mynamebvh.com/blog/infrastructures/db"
	"mynamebvh.com/blog/internal/routes"
	"mynamebvh.com/blog/internal/web"
)

func main() {

	app := fiber.New()

	sqlServer := db.ConnectDB()

	routeStruct := routes.RouterStruct{
		RouterStruct: web.RouterStruct{
			Web:       app,
			SqlServer: sqlServer,
		},
	}

	router := routes.NewHttpRoute(routeStruct)
	router.GetRoutes()

	app.Listen(":3000")
}
