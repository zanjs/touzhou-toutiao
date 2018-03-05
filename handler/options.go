package handler

import (
	"anla.io/taizhou-ir/response"
	"github.com/kataras/iris"
)

// OptionsHandler is
func OptionsHandler(ctx iris.Context) {
	response.JSON(ctx, "hello")
}
