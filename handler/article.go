package handler

import (
	"fmt"
	"strconv"

	"anla.io/taizhou-ir/models"
	"anla.io/taizhou-ir/response"
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
	pageNoStr := ctx.Request().FormValue("page_no")
	var pageNo int
	var err error
	if pageNo, err = strconv.Atoi(pageNoStr); err != nil {
		pageNo = 1
	}

	page := models.PageModel{}

	page.Num = pageNo

	datas, err := models.Article{}.GetAll(&page)
	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}
	fmt.Println(datas)
	response.JSONPage(ctx, datas, page)
}
