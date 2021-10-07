package router

import (
	auth "mynamebvh.com/blog/internal/utils"
	"mynamebvh.com/blog/internal/web"
)

type RouterStruct struct {
	web.RouterStruct
	jwtAuth auth.JwtTokenInterface
}