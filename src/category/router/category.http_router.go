package router

import (
	"log"

	middlewares "mynamebvh.com/blog/internal/middlewares"
	"mynamebvh.com/blog/src/category/handler"
	"mynamebvh.com/blog/src/category/repositories"
	"mynamebvh.com/blog/src/category/services"
)

func NewHttpRoute(
	structs RouterStruct,
) RouterStruct {
	log.Println("Setup HTTP Category Route")

	return structs
}

func (r *RouterStruct) GetRoute() {
	categorySqlServerRepo := repositories.NewUserRepostiory(r.SqlServer)
	categoryService := services.NewUserService(categorySqlServerRepo)
	categoryHandler := handler.NewUserHttpHandler(categoryService)

	r.Web.Get("/api/category", categoryHandler.GetCategory)
	r.Web.Post("/api/category", middlewares.Protected(), categoryHandler.CreateCategory)
	r.Web.Put("/api/category/:id", middlewares.Protected(), categoryHandler.UpdateCategory)
	r.Web.Delete("/api/category/:id", middlewares.Protected(), categoryHandler.DeleteCategory)
}
