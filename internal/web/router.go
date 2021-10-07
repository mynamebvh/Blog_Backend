package web

import (
	"github.com/gofiber/fiber/v2"
	db "mynamebvh.com/blog/infrastructures/db"
)

type RouterStruct struct {
	Web         *fiber.App
	SqlServer   db.SqlServer
}
