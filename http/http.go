package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func GetHttp(urlInfo string, headers map[string]string, timeOut int64, Params map[string]string, allowRedirects bool, proxyURL string) *Response {
	myresponse := New_response()

	if Params != nil {
		data := url.Values{}
		for key, vlue := range Params {
			data.Set(key, vlue)
		}
		urlInfo = urlInfo + "?" + data.Encode()
	}

	// 设置代理
	var proxyFunc func(*http.Request) (*url.URL, error)
	if proxyURL != "" {
		proxyURLParsed, err := url.Parse(proxyURL)
		if err != nil {
			// Handle error if the proxy URL is invalid
			return nil
		}
		proxyFunc = http.ProxyURL(proxyURLParsed)
	} else {
		proxyFunc = http.ProxyFromEnvironment
	}

	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true

	req, err := http.NewRequest("GET", urlInfo, nil)
	if err != nil {
		myresponse.ErrMessage = err.Error()
		return myresponse
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	timeOut2 := time.Duration(timeOut) * time.Second
	client := &http.Client{
		Timeout: timeOut2,
		Transport: &http.Transport{
			Proxy:           proxyFunc,
			TLSClientConfig: tlsConfig,
			//DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			//	// 建立一个原始连接
			//	dialer := &net.Dialer{Timeout: 10 * time.Second}
			//	//conn, err := dialer.DialContext(ctx, network, proxyURL.Host)
			//	conn, err := dialer.DialContext(ctx, network, addr)
			//	if err != nil {
			//		return nil, err
			//	}
			//
			//	// 使用 uTLS 扩展原始连接
			//	uconn := utls.UClient(conn, &utls.Config{
			//		InsecureSkipVerify: true,
			//	}, utls.HelloChrome_120)
			//	// 完成握手
			//	if err := uconn.Handshake(); err != nil {
			//		return nil, err
			//	}
			//	return uconn, nil
			//},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if !allowRedirects && len(via) >= 1 {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		myresponse.ErrMessage = err.Error()
		return myresponse
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		myresponse.ErrMessage = err.Error()
		return myresponse
	}
	jsonString, _ := json.Marshal(resp.Header)
	myresponse.ReHeaders = string(jsonString)
	myresponse.IsError = false
	myresponse.Text = string(body)
	myresponse.Status = resp.Status

	myresponse.Content = body

	return myresponse
}

func PostHttp(urlInfo string, headers map[string]string, timeOut int64, Params map[string]string, allowRedirects bool, proxyURL string, data string) *Response {
	myresponse := New_response()

	if Params != nil {
		data := url.Values{}
		for key, vlue := range Params {
			data.Set(key, vlue)
		}
		urlInfo = urlInfo + "?" + data.Encode()
	}

	// 设置代理
	var proxyFunc func(*http.Request) (*url.URL, error)
	if proxyURL != "" {
		proxyURLParsed, err := url.Parse(proxyURL)
		if err != nil {
			// Handle error if the proxy URL is invalid
			return nil
		}
		proxyFunc = http.ProxyURL(proxyURLParsed)
	} else {
		proxyFunc = http.ProxyFromEnvironment
	}

	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true

	req, err := http.NewRequest("POST", urlInfo, bytes.NewBuffer([]byte(data)))
	if err != nil {
		myresponse.ErrMessage = err.Error()
		return myresponse
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	timeOut2 := time.Duration(timeOut) * time.Second
	client := &http.Client{
		Timeout: timeOut2,
		Transport: &http.Transport{
			Proxy:           proxyFunc,
			TLSClientConfig: tlsConfig,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if !allowRedirects && len(via) >= 1 {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}
	resp, err := client.Do(req)

	if err != nil {
		myresponse.ErrMessage = err.Error()
		return myresponse
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		myresponse.ErrMessage = err.Error()
		return myresponse
	}
	jsonString, _ := json.Marshal(resp.Header)
	myresponse.ReHeaders = string(jsonString)
	myresponse.IsError = false
	myresponse.Text = string(body)
	myresponse.Status = resp.Status
	myresponse.Content = body

	return myresponse
}
