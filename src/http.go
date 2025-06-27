package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func _httpClient(addr string, method string, headers *map[string]string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, addr, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 Edg/134.0.0.0`)

	if headers != nil {
		for k, v := range *headers {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{}
	return client.Do(req)
}

func Get(addr string, params *map[string]string) ([]byte, error) {
	data := url.Values{}
	if params != nil {
		for k, v := range *params {
			data.Set(k, v)
		}
	}

	u, _ := url.ParseRequestURI(addr)
	u.RawQuery = data.Encode()

	resp, err := _httpClient(u.String(), "GET", nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func PostForm(addr string, params *map[string]string) ([]byte, error) {
	data := url.Values{}
	if params != nil {
		for k, v := range *params {
			data.Set(k, v)
		}
	}

	headers := map[string]string{}
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	resp, err := _httpClient(addr, "POST", &headers, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func PostJson(addr string, data any) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{}
	headers["Content-Type"] = "application/json"
	resp, err := _httpClient(addr, "POST", &headers, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
