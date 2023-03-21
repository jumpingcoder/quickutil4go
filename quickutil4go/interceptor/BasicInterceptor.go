package interceptor

import (
	"github.com/jumpingcoder/quickutil4go/quickutil4go/logutil"
	"github.com/kataras/iris/v12"
	"regexp"
)

type BasicInterceptor struct {
	includeSlice []string
	excludeSlice []string
}

func NewBasicInterceptor() *BasicInterceptor {
	return &BasicInterceptor{}
}

func (i *BasicInterceptor) Include(path string) {
	i.includeSlice = append(i.includeSlice, path)
}

func (i *BasicInterceptor) Exclude(path string) {
	i.excludeSlice = append(i.excludeSlice, path)
}

func (i *BasicInterceptor) NeedToIntercept(ctx iris.Context) bool {
	for _, excludeRegexp := range i.excludeSlice {
		matched, err := regexp.MatchString(excludeRegexp, ctx.Request().RequestURI)
		if err != nil {
			logutil.Error(nil, err)
		}
		return !matched
	}
	for _, includeRegexp := range i.includeSlice {
		matched, err := regexp.MatchString(includeRegexp, ctx.Request().RequestURI)
		if err != nil {
			logutil.Error(nil, err)
		}
		return matched
	}
	return false
}
