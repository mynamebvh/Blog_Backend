package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"mynamebvh.com/blog/internal/web"
	"mynamebvh.com/blog/src/user/dto"
	"mynamebvh.com/blog/src/user/services"
)

type AuthHandlers interface {
	Login(ctx *fiber.Ctx) error
}

type authHandlers struct {
	UserService services.UserService
}

func NewHttpHandler(
	userService services.UserService,
) AuthHandlers {
	return &authHandlers{
		UserService: userService,
	}
}

func (services *authHandlers) Login(ctx *fiber.Ctx) error{
	userData := new(dto.UserLogin)

	if err := ctx.BodyParser(userData); err != nil {
		log.Fatal(err)
	}

	res, err := services.UserService.Login(userData)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "", res)
}

