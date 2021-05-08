package controller

import (
	"github.com/jumpingcoder/quickutil4go/entity"
	"github.com/jumpingcoder/quickutil4go/utils/logutil"
	"github.com/kataras/iris/v12"
	"time"
)

func Index(ctx iris.Context) {
	ctx.HTML("<h1>Hello World!</h1>")
}

func Get(ctx iris.Context) {
	list := []int32{1}
	list = append(list, 2)
	message := make(map[string]interface{})
	message["name"] = list
	ctx.JSON(entity.ResultVO{Code: 200, Message: message})
}

func SlowGet(ctx iris.Context) {
	time.Sleep(time.Second * 30)
	ctx.JSON(entity.ResultVO{})
}

func Post(ctx iris.Context) {
	var request entity.TestBO
	err := ctx.ReadJSON(&request)
	if err != nil {
		logutil.Error(nil, err)
		ctx.StatusCode(400)
		ctx.JSON(entity.ResultVO{Code: 400, Message: "request json parsed failed"})
		return
	}
	response := entity.TestBO{Uid: request.Uid, Input: "input", Output: "output"}
	ctx.JSON(entity.ResultVO{Success: true, Code: 200, Message: response})
}

func Header(ctx iris.Context) {
	ctx.JSON(ctx.Request().Header)
}
