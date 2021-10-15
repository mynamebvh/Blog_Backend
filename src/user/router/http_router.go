package router

import (
	"log"

	middlewares "mynamebvh.com/blog/internal/middlewares"
	"mynamebvh.com/blog/src/user/handlers"
	"mynamebvh.com/blog/src/user/repositories"
	"mynamebvh.com/blog/src/user/services"
)

func NewHttpRoute(
	structs RouterStruct,
) RouterStruct {
	log.Println("Setup HTTP Users Route")

	return structs
}

func (r *RouterStruct) GetRoute() {
	userSqlServerRepo := repositories.NewUserRepostiory(r.SqlServer)
	userService := services.NewUserService(userSqlServerRepo, r.jwtAuth)
	authHandlers := handlers.NewHttpHandler(userService)
	userHandlers := handlers.NewUserHttpHandler(userService)

	r.Web.Post("/api/auth/login", authHandlers.Login)
	r.Web.Post("/api/auth/signup", userHandlers.CreateUser)
	r.Web.Get("/api/user/:id", userHandlers.GetUser)
	r.Web.Delete("/api/user", middlewares.Protected(), userHandlers.DeleteUser)
	r.Web.Put("/api/user", middlewares.Protected(), userHandlers.UpdateUser)
}
