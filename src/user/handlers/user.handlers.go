package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"mynamebvh.com/blog/internal/web"
	dto "mynamebvh.com/blog/src/user/dto"
	"mynamebvh.com/blog/src/user/services"
)

type UserHandlers interface{
	GetUser(ctx *fiber.Ctx) error
	CreateUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
}

type userHandlers struct {
	UserService services.UserService
}

func NewUserHttpHandler(
	userService services.UserService,
) UserHandlers {
	return &userHandlers{
		UserService: userService,
	}
}

func (services *userHandlers) GetUser(ctx *fiber.Ctx) error{

	return ctx.SendString("hi")
}

func (services *userHandlers) CreateUser(ctx *fiber.Ctx) error{
	newUser := new(dto.UserRequest)

	if err := ctx.BodyParser(&newUser); err != nil {
		log.Fatal(err)
	}

	res, err := services.UserService.Signup(newUser)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, "Tạo tài khoản thành công", res)
}

func (services *userHandlers) UpdateUser(ctx *fiber.Ctx) error{
	return ctx.SendString("hi")
}

func (services *userHandlers) DeleteUser(ctx *fiber.Ctx) error{
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	_, err := services.UserService.Delete(uint(id))

	if err != nil {
		return web.JsonResponse(ctx, 404, err.Error() , nil)
	}	else {
		return web.JsonResponse(ctx, 200, "Xoá thành công", nil)
	}
	
		
}