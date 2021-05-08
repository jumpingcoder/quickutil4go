package interceptor

import (
	"github.com/jumpingcoder/quickutil4go/entity"
	"github.com/kataras/iris/v12"
)

type ScriptInterceptor struct {
	BasicInterceptor
}

func NewScriptInterceptor() *ScriptInterceptor {
	return &ScriptInterceptor{}
}

func (i *ScriptInterceptor) Interceptor(ctx iris.Context) {
	if !i.NeedToIntercept(ctx) {
		ctx.Next()
		return
	}
	if ctx.IsScript() {
		ctx.StatusCode(403)
		ctx.JSON(entity.ResultVO{Code: 403, Message: "Refuse to crawl"})
	} else {
		ctx.Next()
	}
}
