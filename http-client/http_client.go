package http_client

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

type HttpClient struct {
	client *http.Client
}

// Http客户端
func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: &http.Client{},
	}
}

// 有代理的Http客户端
func NewHttpClientWithProxy(proxyUrl *url.URL) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		},
	}
}

// Http GET Method
// 使用x-www-urlencoded编码
func (c *HttpClient) Get(url string, params url.Values) ([]byte, error) {
	fullURL := url + "?" + params.Encode()
	resp, err := c.client.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Http POST Method
func (c *HttpClient) Post(url string, data []byte, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Http POST Method
// 发送Json
func (c *HttpClient) PostJson(url string, data []byte, headers map[string]string) ([]byte, error) {
	headers["Content-Type"] = "application/json"
	return c.Post(url, data, headers)
}

// Http POST Method
// 发送FormData
func (c *HttpClient) PostForm(url string, data map[string]string, headers map[string]string) ([]byte, error) {
	buff := new(bytes.Buffer)
	form := multipart.NewWriter(buff)
	for k, v := range data {
		form.WriteField(k, v)
	}
	form.Close()
	headers["Content-Type"] = form.FormDataContentType()
	return c.Post(url, buff.Bytes(), headers)
}

func (c *HttpClient) PostFile(url string, data map[string]string, fileKey string, file *multipart.FileHeader, headers map[string]string) ([]byte, error) {

	fh, err := file.Open()
	if err != nil {
		fmt.Println("read file error:" + err.Error())
		return nil, err
	}

	buff := new(bytes.Buffer)
	form := multipart.NewWriter(buff)
	for k, v := range data {
		form.WriteField(k, v)
	}
	formIo, err := form.CreateFormFile(fileKey, file.Filename)
	if err != nil {
		fmt.Println("creats stream:" + err.Error())
		return nil, err
	}
	io.Copy(formIo, fh)
	form.Close()
	headers["Content-Type"] = form.FormDataContentType()
	return c.Post(url, buff.Bytes(), headers)
}
