package handler

import (
	"anla.io/taizhou-y/models"
	"anla.io/taizhou-y/response"
	"github.com/kataras/iris"
)

type (
	// Article is
	Article struct {
		Controller
	}
)

// Create is
func (ctl Article) Create(ctx iris.Context) {
	u := &models.Article{}
	if err := ctx.ReadJSON(u); err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	if u.Content == "" {
		response.JSONError(ctx, "Content where?")
		return
	}

	user := ctl.GetUser(ctx)

	u.UserID = user.ID

	err := models.Article{}.Create(u)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	response.JSON(ctx, u)
}

// All is
func (ctl Article) All(ctx iris.Context) {
	datas, err := models.Article{}.GetAll()
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}
	response.JSON(ctx, datas)
}
