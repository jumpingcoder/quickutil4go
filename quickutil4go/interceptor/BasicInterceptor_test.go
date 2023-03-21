package interceptor

import (
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	t.Log(regexp.MatchString("/*", "/"))
	t.Log(regexp.MatchString("/*", "/hello/world/123"))
	t.Log(regexp.MatchString("/hello/*", "/hello/world/123"))
	t.Log(regexp.MatchString("/*/world/*", "/hello/world/123"))
}