package http

import (
	"fmt"
	"testing"
)

func TestCyJSHTTPGet(t *testing.T) {
	headers := map[string]string{
		"Authorization": "ewsew",
	}
	GetHttp("https://www.baidu.com", headers, 20, nil, true, "")
}

func TestCyJSHTTPGetToJSON(t *testing.T) {
	headers := map[string]string{
		"Authorization": "ewsew",
	}
	dd := GetHttp("https://tls.browserleaks.com/json", headers, 1, nil, true, "")
	fmt.Println(dd.IsError)
	ff := dd.ToJSON()
	fmt.Println(ff)
}
