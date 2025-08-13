package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	cookiejar "github.com/juju/persistent-cookiejar"
)

var defaultCookieRootPath = ""

func init() {
	SetCookiesSavePath(filepath.Join(Pwd(), "cookies"))
}

func _httpClient(addr string, method string, headers map[string]string, body io.Reader, cookieFileName string) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, addr, body)
	if err != nil {
		return nil, err
	}

	_, ok := headers["User-Agent"]
	if !ok {
		_, ok = headers["user-agent"]
		if !ok {
			req.Header.Set("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 Edg/134.0.0.0`)
		}
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if len(cookieFileName) != 0 {
		jar, err := cookiejar.New(&cookiejar.Options{
			Filename: filepath.Join(defaultCookieRootPath, cookieFileName),
		})

		if err != nil {
			return nil, err
		}

		client := &http.Client{
			Jar: jar,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			}}

		resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}

		os.MkdirAll(defaultCookieRootPath, 0755)
		jar.Save()

		return resp, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}}

	return client.Do(req)
}

func SetCookiesSavePath(fpath string) {
	defaultCookieRootPath = fpath
}

func Get(addr string, params map[string]string, headers map[string]string, cookieName ...string) ([]byte, error) {
	if len(params) != 0 {
		data := url.Values{}
		for k, v := range params {
			data.Set(k, v)
		}

		u, _ := url.ParseRequestURI(addr)
		u.RawQuery = data.Encode()
		addr = u.String()
	}

	cookieFileName := ""

	if len(cookieName) >= 1 {
		cookieFileName = cookieName[0]
	}

	resp, err := _httpClient(addr, "GET", headers, nil, cookieFileName)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func PostForm(addr string, headers map[string]string, data map[string]string, cookieName ...string) ([]byte, error) {
	vals := url.Values{}
	for k, v := range data {
		vals.Set(k, v)
	}

	cookieFileName := ""

	if len(cookieName) >= 1 {
		cookieFileName = cookieName[0]
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	resp, err := _httpClient(addr, "POST", headers, bytes.NewBufferString(vals.Encode()), cookieFileName)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func PostJson(addr string, headers map[string]string, data any, cookieName ...string) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	cookieFileName := ""

	if len(cookieName) >= 1 {
		cookieFileName = cookieName[0]
	}

	headers["Content-Type"] = "application/json"
	resp, err := _httpClient(addr, "POST", headers, bytes.NewBuffer(jsonData), cookieFileName)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
