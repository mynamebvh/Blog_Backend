package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"mynamebvh.com/blog/internal/enums"
	"mynamebvh.com/blog/internal/utils"
	"mynamebvh.com/blog/internal/web"
	dto "mynamebvh.com/blog/src/user/dto"
	"mynamebvh.com/blog/src/user/services"
)

type UserHandlers interface {
	GetUser(ctx *fiber.Ctx) error
	GetAllUser(ctx *fiber.Ctx) error
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

func (services *userHandlers) GetUser(ctx *fiber.Ctx) error {

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusOK, enums.ERROR_ID_NOT_FOUND, nil)
	}

	user := services.UserService.FindByID(uint(id))

	if user.ID == 0 {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_USER_NOT_FOUND, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, user)
}

func (services *userHandlers) GetAllUser(ctx *fiber.Ctx) error {

	return ctx.SendString("hi")
}

func (services *userHandlers) CreateUser(ctx *fiber.Ctx) error {
	newUser := new(dto.UserRequest)

	if err := ctx.BodyParser(&newUser); err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, nil)

	}

	res, err := services.UserService.Signup(newUser)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, res)
}

func (services *userHandlers) UpdateUser(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	userUpdate := new(dto.UserUpdate)

	if err := ctx.BodyParser(&userUpdate); err != nil {
		web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, err)
	}

	errors := utils.Validate(userUpdate)

	if errors != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, errors)
	}

	userUpdated, err := services.UserService.Update(uint(id), *userUpdate)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, enums.ERROR_UPDATE, userUpdated)
	} else {
		return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, userUpdate)
	}
}

func (services *userHandlers) DeleteUser(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	err := services.UserService.Delete(uint(id))

	if err != nil {
		return web.JsonResponse(ctx, 404, enums.ERROR_DELETE, nil)
	} else {
		return web.JsonResponse(ctx, 200, enums.MSG_SUCCESS, nil)
	}
}
