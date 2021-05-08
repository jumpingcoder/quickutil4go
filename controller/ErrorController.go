package controller

import (
	"github.com/jumpingcoder/quickutil4go/entity"
	"github.com/kataras/iris/v12"
	"strings"
)

func InternalError(ctx iris.Context) {
	code := ctx.GetStatusCode()
	ctx.JSON(entity.ResultVO{Code: code})
}

func NotFound(ctx iris.Context) {
	if strings.Contains(strings.ToLower(ctx.GetHeader("accept")), "html") {
		ctx.StatusCode(200)
		Index(ctx)
	} else {
		ctx.JSON(entity.ResultVO{Code: 404})
	}
}
