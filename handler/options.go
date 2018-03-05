package handler

import (
	"anla.io/taizhou-y/response"
	"github.com/kataras/iris"
)

// OptionsHandler is
func OptionsHandler(ctx iris.Context) {
	response.JSON(ctx, "hello")
}
