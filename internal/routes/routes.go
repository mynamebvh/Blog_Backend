package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"mynamebvh.com/blog/internal/web"
	categoryRouter "mynamebvh.com/blog/src/category/router"
	tagRouter "mynamebvh.com/blog/src/tag/router"
	userRouter "mynamebvh.com/blog/src/user/router"
)

type RouterStruct struct {
	web.RouterStruct
}

func NewHttpRoute(r RouterStruct) RouterStruct {
	log.Println("Loading the HTTP Router")

	return r
}

func (c *RouterStruct) GetRoutes() {
	c.Web.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Hello this is my first route in go fiber"))
	})

	webRouterConfig := web.RouterStruct{
		Web:       c.Web,
		SqlServer: c.SqlServer,
	}
	// registering route from another modules
	userRouterStruct := userRouter.RouterStruct{
		RouterStruct: webRouterConfig,
	}
	userRouter := userRouter.NewHttpRoute(userRouterStruct)
	userRouter.GetRoute()

	categoryRouterStruct := categoryRouter.RouterStruct{
		RouterStruct: webRouterConfig,
	}
	categoryRouter := categoryRouter.NewHttpRoute(categoryRouterStruct)
	categoryRouter.GetRoute()

	tagRouterStruct := tagRouter.RouterStruct{
		RouterStruct: webRouterConfig,
	}
	tagRouter := tagRouter.NewHttpRoute(tagRouterStruct)
	tagRouter.GetRoute()

	// handling 404 error
	c.Web.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
}
