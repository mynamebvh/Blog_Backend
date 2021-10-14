package router

import (
	"log"

	middlewares "mynamebvh.com/blog/internal/middlewares"
	"mynamebvh.com/blog/src/post/handler"
	"mynamebvh.com/blog/src/post/repositories"
	"mynamebvh.com/blog/src/post/services"
)

func NewHttpRoute(
	structs RouterStruct,
) RouterStruct {
	log.Println("Setup HTTP Tag Route")

	return structs
}

func (r *RouterStruct) GetRoute() {
	postSqlServerRepo := repositories.NewUserRepostiory(r.SqlServer)
	postService := services.NewUserService(postSqlServerRepo)
	postHandler := handler.NewUserHttpHandler(postService)

	r.Web.Get("/api/post", middlewares.Protected(), postHandler.GetPost)
	r.Web.Post("/api/post", middlewares.Protected(), postHandler.CreatePost)
	r.Web.Put("/api/post/:id", middlewares.Protected(), postHandler.UpdatePost)
	r.Web.Delete("/api/post/:id", middlewares.Protected(), postHandler.DeletePost)
}
