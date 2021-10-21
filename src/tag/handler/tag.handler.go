package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"mynamebvh.com/blog/internal/enums"
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
	return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, services.tagService.FindByAll())
}

func (services *TagHandler) CreateTag(ctx *fiber.Ctx) error {

	newTag := new(dto.Tag)

	if err := ctx.BodyParser(&newTag); err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, nil)
	}

	errors := utils.Validate(newTag)

	if errors != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, nil)
	}

	res, err := services.tagService.Save(*newTag)

	if err != nil {
		web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR, err)
	}
	return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, res)
}

func (services *TagHandler) UpdateTag(ctx *fiber.Ctx) error {
	tagUpdate := new(dto.Tag)

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_ID_NOT_FOUND, nil)
	}

	if err := ctx.BodyParser(&tagUpdate); err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, nil)
	}

	errors := utils.Validate(tagUpdate)

	if errors != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_VALIDATE, nil)
	}

	category, err := services.tagService.Update(uint(id), *tagUpdate)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, enums.ERROR_UPDATE, nil)
	}

	return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, category)
}

func (services *TagHandler) DeleteTag(ctx *fiber.Ctx) error {

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, enums.ERROR_ID_NOT_FOUND, nil)
	}

	err = services.tagService.Delete(uint(id))

	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, enums.ERROR_DELETE, nil)
	} else {
		return web.JsonResponse(ctx, http.StatusOK, enums.MSG_SUCCESS, nil)
	}
}
