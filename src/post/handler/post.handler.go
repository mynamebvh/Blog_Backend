package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"mynamebvh.com/blog/internal/utils"
	"mynamebvh.com/blog/internal/web"
	"mynamebvh.com/blog/src/post/dto"
	"mynamebvh.com/blog/src/post/services"
)

type PostHandlerInterface interface {
	GetPost(ctx *fiber.Ctx) error
	GetAllPost(ctx *fiber.Ctx) error
	CreatePost(ctx *fiber.Ctx) error
	UpdatePost(ctx *fiber.Ctx) error
	DeletePost(ctx *fiber.Ctx) error
}

type PostHandler struct {
	postService services.PostServiceInterface
}

func NewUserHttpHandler(
	postService services.PostServiceInterface,
) PostHandlerInterface {
	return &PostHandler{
		postService: postService,
	}
}
func (services *PostHandler) GetAllPost(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Query("page"))

	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(ctx.Query("pageSize"))

	if err != nil {
		pageSize = 10
	}

	postList := services.postService.FindByAll(page, pageSize)

	return web.JsonResponse(ctx, http.StatusOK, "Thành công", postList)
}

func (services *PostHandler) GetPost(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, 404, "Lỗi máy chủ", err.Error())
	}

	post := services.postService.FindById(uint(id))

	if post.UserID == 0 {
		return web.JsonResponse(ctx, http.StatusBadRequest, "Không tìm thấy bài đăng", nil)
	}
	return web.JsonResponse(ctx, http.StatusOK, "Thành công", post)
}

func (services *PostHandler) CreatePost(ctx *fiber.Ctx) error {

	newPost := new(dto.Post)

	if err := ctx.BodyParser(&newPost); err != nil {
		log.Fatal(err)
	}

	errors := utils.Validate(newPost)

	if errors != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "Lỗi validate", errors)
	}

	res, err := services.postService.Save(*newPost)

	if err != nil {
		web.JsonResponse(ctx, http.StatusBadRequest, "Lỗi", err)
	}
	return web.JsonResponse(ctx, http.StatusOK, "Tạo thành công", res)
}

func (services *PostHandler) UpdatePost(ctx *fiber.Ctx) error {
	postUpdate := new(dto.PostUpdate)

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, "Lỗi", err.Error())
	}

	if err := ctx.BodyParser(&postUpdate); err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, "Lỗi", err.Error())
	}

	errors := utils.Validate(postUpdate)

	if errors != nil {
		return web.JsonResponse(ctx, http.StatusBadRequest, "Lỗi validate", errors)
	}

	category, err := services.postService.Update(uint(id), *postUpdate)

	if err != nil {
		return web.JsonResponse(ctx, http.StatusInternalServerError, "Lỗi", err.Error())
	}

	return web.JsonResponse(ctx, http.StatusOK, "Cập nhật thành công", category)
}

func (services *PostHandler) DeletePost(ctx *fiber.Ctx) error {

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)

	if err != nil {
		return web.JsonResponse(ctx, 404, err.Error(), err.Error())
	}

	err = services.postService.Delete(uint(id))

	if err != nil {
		return web.JsonResponse(ctx, 404, "Lỗi", err.Error())
	} else {
		return web.JsonResponse(ctx, 200, "Xoá thành công", nil)
	}
}
