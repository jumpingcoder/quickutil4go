package interceptor

import (
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
		result := make(map[string]interface{})
		result["success"] = false
		result["code"] = 403
		result["message"] = "Refuse to crawl"
		ctx.JSON(result)
	} else {
		ctx.Next()
	}
}
