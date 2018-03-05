package handler

import (
	"anla.io/taizhou-ir/models"
	"anla.io/taizhou-ir/response"
	"github.com/kataras/iris"
)

// InitTable is 初始化表结构
func InitTable(ctx iris.Context) {
	models.CreateTable()
	response.JSON(ctx, "are you ok")
}
