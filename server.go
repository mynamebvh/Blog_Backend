package main

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	db "mynamebvh.com/blog/infrastructures/db"
	"mynamebvh.com/blog/internal/routes"
	"mynamebvh.com/blog/internal/web"
)

func main() {

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {

			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			if err != nil {
				return web.JsonResponse(ctx, code, "Lỗi máy chủ", nil)
			}

			return nil
		},
	})

	app.Use(recover.New())

	app.Use(limiter.New(limiter.Config{
		Next:         func(c *fiber.Ctx) bool { return false },
		Max:          60,
		Expiration:   10 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string { return c.Get("x-forwarded-for") },
		LimitReached: func(c *fiber.Ctx) error {
			return web.JsonResponse(c, http.StatusTooManyRequests, "Bạn đang truy cập quá nhanh", nil)
		},
	}))

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
