package main

import (
	"flag"
	"fmt"
	"github.com/jumpingcoder/quickutil4go/controller"
	"github.com/jumpingcoder/quickutil4go/interceptor"
	"github.com/jumpingcoder/quickutil4go/utils/logutil"
	"strings"
)

import "github.com/kataras/iris/v12"

var port string
var env string
var decryptKey string

func configService() {
	flag.StringVar(&port, "port", "9000", "端口号，默认为9000")
	flag.StringVar(&env, "env", "dev", "环境，默认为dev")
	flag.StringVar(&decryptKey, "key", "", "配置解密配置的key")
	flag.Parse()
	if !strings.EqualFold(env, "dev") {
		logutil.SetLogType(logutil.JSON)
		logutil.SetLogLevel(logutil.INFO)
	}
	logutil.Info(fmt.Sprintf("Start with port=%v, env=%v, key=%v", port, env, decryptKey), nil)
}

func configInterceptor(app *iris.Application) {
	scriptInterceptor := interceptor.NewScriptInterceptor()
	scriptInterceptor.Include("/*")
	scriptInterceptor.Exclude("/get*")
	app.Use(scriptInterceptor.Interceptor)
}

func configRouter(app *iris.Application) {
	app.OnErrorCode(iris.StatusNotFound, controller.NotFound)
	app.OnErrorCode(iris.StatusInternalServerError, controller.InternalError)
	app.Get("/", controller.Index)
	app.Get("/get", controller.Get)
	app.Get("/slowGet", controller.SlowGet)
	app.Post("/post", controller.Post)
	app.Get("/header", controller.Header)
}

func main() {
	app := iris.New()
	configService()
	configInterceptor(app)
	configRouter(app)
	app.Listen(":" + port)
}
