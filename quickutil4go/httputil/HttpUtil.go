package httputil

import (
	"bytes"
	"net/http"
	"time"
)

func HttpGet(url string, headers map[string]string) (*http.Response, error) {
	return HttpRequest("GET", url, headers, nil, 0)
}

func HttpPost(url string, headers map[string]string, body []byte) (*http.Response, error) {
	return HttpRequest("POST", url, headers, body, 0)
}

func HttpPut(url string, headers map[string]string, body []byte) (*http.Response, error) {
	return HttpRequest("PUT", url, headers, body, 0)
}

func HttpDelete(url string, headers map[string]string) (*http.Response, error) {
	return HttpRequest("PUT", url, headers, nil, 0)
}

func HttpRequest(method string, url string, headers map[string]string, body []byte, timeoutSeconds int64) (*http.Response, error) {
	client := &http.Client{}
	if timeoutSeconds > 0 {
		client.Timeout = time.Duration(time.Second.Nanoseconds() * timeoutSeconds)
	}
	request, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	if headers != nil {
		for key := range headers {
			request.Header.Add(key, headers[key])
		}
	}
	return client.Do(request)
}
