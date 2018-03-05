package handler

import (
	"anla.io/taizhou-y/response"
	"github.com/kataras/iris"
)

// IndexHand is
func IndexHand(ctx iris.Context) {
	response.JSON(ctx, "hello boy!")
}
