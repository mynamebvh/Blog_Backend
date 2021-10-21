package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"mynamebvh.com/blog/internal/enums"
	"mynamebvh.com/blog/internal/utils"
	"mynamebvh.com/blog/internal/web"
	"mynamebvh.com/blog/src/category/dto"
	"mynamebvh.com/blog/src/category/services"
)

type CategoryHandlerInterface interface {
	GetAllCategory(ctx *fiber.Ctx) error
	GetCategory(crx *fiber.Ctx) error
	CreateCategory(ctx *fiber.Ctx) error
	UpdateCategory(ctx *fiber.Ctx) error
	DeleteCategory(ctx *fiber.Ctx) error
}

type CategoryHandler struct {
	categoryService services.CategoryServiceInterface
}

func NewUserHttpHandler(
	categoryService services.CategoryServiceInterface,
) CategoryHandlerInterface {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (services *CategoryHandler) GetAllCategory(ctx *fiber.Ctx) error {
	return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, services.categoryService.FindByAll())
}

func (services *CategoryHandler) GetCategory(ctx *fiber.Ctx) error {

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, nil)
	}

	page, err := strconv.Atoi(ctx.Query("page"))

	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(ctx.Query("pageSize"))

	if err != nil {
		pageSize = 10
	}

	return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, services.categoryService.FindById(uint(id), page, pageSize))
}

func (services *CategoryHandler) CreateCategory(ctx *fiber.Ctx) error {

	newCategory := new(dto.Category)

	if err := ctx.BodyParser(&newCategory); err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, nil)

	}

	errors := utils.Validate(newCategory)

	if errors != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, nil)
	}

	res, err := services.categoryService.Save(*newCategory)

	if err != nil {
		web.JsonResponse(ctx, http.StatusInternalServerError, enums.ERROR_SERVER, nil)
	}
	return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, res)
}

func (services *CategoryHandler) UpdateCategory(ctx *fiber.Ctx) error {
	categoryUpdate := new(dto.Category)

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_ID_NOT_FOUND, nil)
	}

	if err := ctx.BodyParser(&categoryUpdate); err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, nil)

	}

	errors := utils.Validate(categoryUpdate)

	if errors != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, nil)
	}

	category, err := services.categoryService.Update(uint(id), *categoryUpdate)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, enums.ERROR_UPDATE, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, category)
}

func (services *CategoryHandler) DeleteCategory(ctx *fiber.Ctx) error {

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_ID_NOT_FOUND, nil)
	}

	err = services.categoryService.Delete(uint(id))

	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, enums.ERROR_DELETE, nil)
	} else {
		return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, nil)
	}
}
