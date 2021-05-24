package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jumpingcoder/quickutil4go/controller"
	"github.com/jumpingcoder/quickutil4go/interceptor"
	"github.com/jumpingcoder/quickutil4go/utils/dbutil"
	"github.com/jumpingcoder/quickutil4go/utils/fileutil"
	"github.com/jumpingcoder/quickutil4go/utils/logutil"
	"strings"
)

import "github.com/kataras/iris/v12"

var config map[string]interface{}
var configDecryptKey string
var port string
var env string

func configService() {
	//配置解析
	var configPath string
	flag.StringVar(&configPath, "config", "./config.json", "配置文件路径，默认为./config.json")
	flag.StringVar(&configDecryptKey, "key", "0000", "配置解密配置的key")
	flag.Parse()
	configBytes, _ := fileutil.File2Byte(configPath)
	json.Unmarshal(configBytes, &config)
	//组件初始化
	dbutil.InitFromConfig(config["GetDB"].([]interface{}), configDecryptKey, dbutil.DefaultDecryptHandler)
	//应用初始化
	env = config["env"].(string)
	port = config["port"].(string)
	if !strings.EqualFold(env, "dev") {
		logutil.SetLogType(logutil.JSON)
		logutil.SetLogLevel(logutil.INFO)
	}
	logutil.Info(fmt.Sprintf("Start with config=%v port=%v, env=%v, key=%v", configPath, port, env, configDecryptKey[0:4]), nil)
}

func configInterceptor(app *iris.Application) {
	scriptInterceptor := interceptor.NewScriptInterceptor()
	scriptInterceptor.Include("/*")
	scriptInterceptor.Exclude("/get*")
	app.Use(scriptInterceptor.Interceptor)
}

func configRouter(app *iris.Application) {
	app.HandleDir("/", "./static")
	app.OnErrorCode(iris.StatusNotFound, controller.NotFound)
	app.OnErrorCode(iris.StatusInternalServerError, controller.InternalError)
	app.Get("/", controller.Index)
	app.Get("/get", controller.Get)
	app.Get("/slowGet", controller.SlowGet)
	app.Get("/getSQL", controller.GetSQL)
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
