package router

import (
	"log"

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
	userMysqlRepo := repositories.NewUserRepostiory(r.SqlServer)
	userService := services.NewUserService(userMysqlRepo, r.jwtAuth)
	authHandlers := handlers.NewHttpHandler(userService)
	userHandlers := handlers.NewUserHttpHandler(userService)

	r.Web.Post("api/auth/login", authHandlers.Login)
	r.Web.Post("api/user/signup", userHandlers.CreateUser)
}