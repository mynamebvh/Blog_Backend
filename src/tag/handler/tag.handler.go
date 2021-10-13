package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"mynamebvh.com/blog/internal/utils"
	"mynamebvh.com/blog/internal/web"
	"mynamebvh.com/blog/src/tag/dto"
	"mynamebvh.com/blog/src/tag/services"
)

type TagHandlerInterface interface {
	GetTag(ctx *fiber.Ctx) error
	CreateTag(ctx *fiber.Ctx) error
	UpdateTag(ctx *fiber.Ctx) error
	DeleteTag(ctx *fiber.Ctx) error
}

type TagHandler struct {
	tagService services.TagServiceInterface
}

func NewUserHttpHandler(
	tagService services.TagServiceInterface,
) TagHandlerInterface {
	return &TagHandler{
		tagService: tagService,
	}
}

func (services *TagHandler) GetTag(ctx *fiber.Ctx) error {
	return web.JsonResponse(ctx, http.StatusOK, "Thành công", services.tagService.FindByAll())
}

func (services *TagHandler) CreateTag(ctx *fiber.Ctx) error {

	newTag := new(dto.Tag)

	if err := ctx.BodyParser(&newTag); err != nil {
		log.Fatal(err)
	}

	errors := utils.Validate(newTag)

	if errors != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "Lỗi validate", errors)
	}

	res, err := services.tagService.Save(*newTag)

	if err != nil {
		web.JsonResponse(ctx, http.StatusBadRequest, "Lỗi", err)
	}
	return web.JsonResponse(ctx, http.StatusOK, "Tạo thành công", res)
}

func (services *TagHandler) UpdateTag(ctx *fiber.Ctx) error {
	tagUpdate := new(dto.Tag)

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, "Lỗi", err.Error())
	}

	if err := ctx.BodyParser(&tagUpdate); err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, "Lỗi", err.Error())
	}

	errors := utils.Validate(tagUpdate)

	if errors != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "Lỗi validate", errors)
	}

	category, err := services.tagService.Update(uint(id), *tagUpdate)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, "Lỗi", err.Error())
	}

	return web.JsonResponse(ctx, http.StatusOK, "Cập nhật thành công", category)
}

func (services *TagHandler) DeleteTag(ctx *fiber.Ctx) error {

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, 404, err.Error(), err.Error())
	}

	err = services.tagService.Delete(uint(id))

	if err != nil {
		return web.JsonResponse(ctx, 404, "Lỗi", err.Error())
	} else {
		return web.JsonResponse(ctx, 200, "Xoá thành công", nil)
	}
}
