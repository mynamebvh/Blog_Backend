package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
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
	return web.JsonResponse(ctx, http.StatusOK, "Thành công", services.categoryService.FindByAll())
}

func (services *CategoryHandler) GetCategory(ctx *fiber.Ctx) error {

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, 404, err.Error(), err.Error())
	}

	page, err := strconv.Atoi(ctx.Query("page"))

	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(ctx.Query("pageSize"))

	if err != nil {
		pageSize = 10
	}

	return web.JsonResponse(ctx, http.StatusOK, "Thành công", services.categoryService.FindById(uint(id), page, pageSize))
}

func (services *CategoryHandler) CreateCategory(ctx *fiber.Ctx) error {

	newCategory := new(dto.Category)

	if err := ctx.BodyParser(&newCategory); err != nil {
		log.Fatal(err)
	}

	errors := utils.Validate(newCategory)

	if errors != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "Lỗi validate", errors)
	}

	res, err := services.categoryService.Save(*newCategory)

	if err != nil {
		web.JsonResponse(ctx, http.StatusBadRequest, "Lỗi", err)
	}
	return web.JsonResponse(ctx, http.StatusOK, "Tạo thành công", res)
}

func (services *CategoryHandler) UpdateCategory(ctx *fiber.Ctx) error {
	categoryUpdate := new(dto.Category)

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, "Lỗi", err.Error())
	}

	if err := ctx.BodyParser(&categoryUpdate); err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, "Lỗi", err.Error())
	}

	errors := utils.Validate(categoryUpdate)

	if errors != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "Lỗi validate", errors)
	}

	category, err := services.categoryService.Update(uint(id), *categoryUpdate)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, "Lỗi", err.Error())
	}

	return web.JsonResponse(ctx, http.StatusOK, "Cập nhật thành công", category)
}

func (services *CategoryHandler) DeleteCategory(ctx *fiber.Ctx) error {

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, 404, err.Error(), err.Error())
	}

	err = services.categoryService.Delete(uint(id))

	if err != nil {
		return web.JsonResponse(ctx, 404, "Lỗi", err.Error())
	} else {
		return web.JsonResponse(ctx, 200, "Xoá thành công", nil)
	}
}
