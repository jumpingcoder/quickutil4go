package logutil

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

const (
	TERMINAL int8 = 0
	JSON     int8 = 1
)

const (
	DEBUG int8 = 0
	INFO  int8 = 1
	WARN  int8 = 2
	ERROR int8 = 3
)

var _logType = TERMINAL
var _logLevel = DEBUG

func SetLogType(logType int8) {
	_logType = logType
}

func SetLogLevel(logLevel int8) {
	_logLevel = logLevel
}

func Error(msg interface{}, err interface{}) {
	if _logLevel <= ERROR {
		print("ERROR", msg, err)
	}
}

func Warn(msg interface{}, err interface{}) {
	if _logLevel <= WARN {
		print("WARN", msg, err)
	}
}

func Info(msg interface{}, err interface{}) {
	if _logLevel <= INFO {
		print("INFO", msg, err)
	}
}

func Debug(msg interface{}, err interface{}) {
	if _logLevel <= DEBUG {
		print("DEBUG", msg, err)
	}
}

func print(level string, msg interface{}, err interface{}) {
	codePath, codeLine := getRuntimeInfo()
	if _logType == JSON {
		var logJson = make(map[string]interface{})
		logJson["level"] = level
		logJson["time"] = time.Now().Format("2006-01-02 15:04:05")
		logJson["msg"] = msg
		logJson["err"] = err
		logJson["path"] = fmt.Sprintf("%v:%v", codePath, codeLine)
		b, err := json.Marshal(logJson)
		if err != nil {
			Error("Json Marshal exception", nil)
		}
		fmt.Printf("%v\n", string(b))
	} else {
		fmt.Printf("[%v][%v]{%v}{%v}(%v:%v)\n", level, time.Now().Format("2006-01-02 15:04:05"), msg, err, codePath, codeLine)
	}
}

func getRuntimeInfo() (string, int) {
	_, codePath, codeLine, ok := runtime.Caller(3)
	if !ok {
		// 不ok，函数栈用尽了
		codePath = "-"
		codeLine = -1
	}
	//codePath = "." + codePath[len(fileutil.GetCurrentPath()):len(codePath)]
	return codePath, codeLine
}
