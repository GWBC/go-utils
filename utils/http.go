package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofrs/flock"
)

var HttpTimeout = 10 * time.Second

var cookieEncry = true
var aesGCM = AesGCM{}
var defaultPathRoot = Pwd()

func init() {
	aesGCM.Init("@#%@#^!$#$%$*^&%^&*#")
}

func SetCookiesSavePath(filepath string) {
	defaultPathRoot = filepath
}

func writeFile(filePath string, data []byte) error {
	fileLock := flock.New(filePath + ".lock")
	err := fileLock.Lock()
	if err != nil {
		return err
	}
	defer fileLock.Unlock()

	if cookieEncry {
		encrpyData, err := aesGCM.Encrypt(string(data))
		if err != nil {
			return err
		}

		data = []byte(encrpyData)
	}

	return os.WriteFile(filePath, data, 0644)
}

func readFile(filePath string) ([]byte, error) {
	fileLock := flock.New(filePath + ".lock")
	err := fileLock.RLock()
	if err != nil {
		return nil, err
	}
	defer fileLock.Unlock()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if cookieEncry {
		rawData, err := aesGCM.Decrypt(string(data))
		if err != nil {
			return nil, err
		}

		data = []byte(rawData)
	}

	return data, err
}

type LocalJar struct {
	Name     string
	rootPath string
	cookies  map[string]*http.Cookie
}

func (l *LocalJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	if len(l.rootPath) == 0 {
		l.cookies = map[string]*http.Cookie{}
		l.rootPath = filepath.Join(defaultPathRoot, "cookies")
		os.MkdirAll(l.rootPath, 0755)
	}

	for _, c := range cookies {
		l.cookies[c.Name] = c
	}

	data, err := json.Marshal(l.cookies)
	if err != nil {
		return
	}

	writeFile(filepath.Join(l.rootPath, l.Name), data)
}

func (l *LocalJar) Cookies(u *url.URL) []*http.Cookie {
	if len(l.rootPath) == 0 {
		l.rootPath = filepath.Join(defaultPathRoot, "cookies")
		os.MkdirAll(l.rootPath, 0755)
		l.cookies = map[string]*http.Cookie{}
		data, err := readFile(filepath.Join(l.rootPath, l.Name))
		if err != nil {
			return nil
		}

		err = json.Unmarshal(data, &l.cookies)
		if err != nil {
			return nil
		}
	}

	cs := []*http.Cookie{}
	for _, c := range l.cookies {
		if strings.Contains(u.Path, c.Path) {
			cs = append(cs, c)
		}
	}

	return cs
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
		jar := &LocalJar{Name: cookieFileName}

		client := &http.Client{
			Jar:     jar,
			Timeout: HttpTimeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			}}

		resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != 200 {
			return nil, errors.New("status error")
		}

		for _, c := range resp.Cookies() {
			if c.MaxAge == 0 && c.Expires.IsZero() {
				c.MaxAge = 24 * 3600
			}
		}

		return resp, err
	}

	client := &http.Client{
		Timeout: HttpTimeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}}

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code:%d", resp.StatusCode)
	}

	return resp, err
}

func Get(addr string, params map[string]string, headers map[string]string, cookieName ...string) ([]byte, error) {
	if len(addr) == 0 {
		return nil, errors.New("url is empty")
	}

	if params == nil {
		params = map[string]string{}
	}

	if headers == nil {
		headers = map[string]string{}
	}

	if len(params) != 0 {
		data := url.Values{}
		for k, v := range params {
			data.Set(k, v)
		}

		u, err := url.ParseRequestURI(addr)
		if err != nil {
			return nil, err
		}

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
	if len(addr) == 0 {
		return nil, errors.New("url is empty")
	}

	if headers == nil {
		headers = map[string]string{}
	}

	if data == nil {
		data = map[string]string{}
	}

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
	if len(addr) == 0 {
		return nil, errors.New("url is empty")
	}

	if headers == nil {
		headers = map[string]string{}
	}

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
