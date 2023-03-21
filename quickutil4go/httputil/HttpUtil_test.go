package httputil

import (
	"io/ioutil"
	"testing"
)

func TestHttpGet(t *testing.T) {
	response, err := HttpRequest("GET", "http://localhost:9000/get2", nil, nil, 60)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}
	t.Log(response.Status)
	t.Log(string(body))
}
