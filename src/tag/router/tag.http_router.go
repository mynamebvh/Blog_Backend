package router

import (
	"log"

	middlewares "mynamebvh.com/blog/internal/middlewares"
	"mynamebvh.com/blog/src/tag/handler"
	"mynamebvh.com/blog/src/tag/repositories"
	"mynamebvh.com/blog/src/tag/services"
)

func NewHttpRoute(
	structs RouterStruct,
) RouterStruct {
	log.Println("Setup HTTP Tag Route")

	return structs
}

func (r *RouterStruct) GetRoute() {
	tagSqlServerRepo := repositories.NewUserRepostiory(r.SqlServer)
	tagService := services.NewUserService(tagSqlServerRepo)
	tagHandler := handler.NewUserHttpHandler(tagService)

	r.Web.Get("/api/tag", tagHandler.GetTag)
	r.Web.Post("/api/tag", middlewares.Protected(), tagHandler.CreateTag)
	r.Web.Put("/api/tag/:id", middlewares.Protected(), tagHandler.UpdateTag)
	r.Web.Delete("/api/tag/:id", middlewares.Protected(), tagHandler.DeleteTag)
}
