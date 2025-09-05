package jsengine

import (
	"encoding/base64"
	"fmt"
	"path/filepath"

	"github.com/GWBC/go-utils/utils"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

/*
	内置函数
	log(string, string)
	atob(string) => string
	btoa(string) => string
	Get(url, header{}) => string
	PostForm(url, header{}, data{}) => string
	PostJson(url, header{}, data{}) => string
	BiliBili2MPD(biliData, proxy) => string
*/

type JSEngine struct {
	vm   *goja.Runtime
	name string
}

func (j *JSEngine) init(name string) error {
	j.name = name
	j.vm = goja.New()
	require.NewRegistry().Enable(j.vm)

	err := j.initLog()
	if err != nil {
		return err
	}

	err = j.initBase64()
	if err != nil {
		return err
	}

	err = j.initHttp()
	if err != nil {
		return err
	}

	err = j.initBiliBili()
	if err != nil {
		return err
	}

	return nil
}

// 获取引擎名称
func (j *JSEngine) Name() string {
	return j.name
}

// 添加函数或变量
func (j *JSEngine) Set(name string, val any) error {
	return j.vm.Set(name, val)
}

// 执行js
func (j *JSEngine) Require(file string, name string) error {
	file = filepath.ToSlash(file)
	_, err := j.RunString(fmt.Sprintf("const %s = require('%s')", name, file))
	return err
}

func (j *JSEngine) RunString(content string) (any, error) {
	v, err := j.vm.RunString(content)
	if err != nil {
		return nil, err
	}

	return v.Export(), nil
}

func (j *JSEngine) initLog() error {
	return j.vm.Set("log", func(msgs ...string) {
		for _, msg := range msgs {
			fmt.Print(msg, " ")
		}

		fmt.Println()
	})
}

func (j *JSEngine) initBase64() error {
	err := j.vm.Set("atob", func(data string) string {
		dcodeData, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			panic(j.vm.ToValue(err.Error()))
		}

		return string(dcodeData)
	})

	if err != nil {
		return err
	}

	return j.vm.Set("btoa", func(data string) string {
		return base64.StdEncoding.EncodeToString([]byte(data))
	})
}

func (j *JSEngine) initHttp() error {
	err := j.vm.Set("Get", func(url string, headers map[string]string, cookieName string) string {
		rsp, err := utils.Get(url, nil, headers, cookieName)
		if err != nil {
			panic(j.vm.ToValue(err.Error()))
		}

		return string(rsp)
	})

	if err != nil {
		return err
	}

	err = j.vm.Set("PostForm", func(url string, headers map[string]string, data map[string]string, cookieName string) string {
		rsp, err := utils.PostForm(url, headers, data, cookieName)
		if err != nil {
			panic(j.vm.ToValue(err.Error()))
		}

		return string(rsp)
	})

	if err != nil {
		return err
	}

	err = j.vm.Set("PostJson", func(url string, headers map[string]string, data map[string]any, cookieName string) string {
		rsp, err := utils.PostJson(url, headers, data, cookieName)
		if err != nil {
			panic(j.vm.ToValue(err.Error()))
		}

		return string(rsp)
	})

	if err != nil {
		return err
	}

	return nil
}

func (j *JSEngine) initBiliBili() error {
	err := j.vm.Set("BiliBili2MPD", func(biliData string,
		proxy string, selectType SelectType) string {
		return BlibiliData2MPD(biliData, proxy, selectType)
	})

	if err != nil {
		return err
	}

	return nil
}

func New(engineName string) (*JSEngine, error) {
	js := &JSEngine{}
	err := js.init(engineName)

	if err != nil {
		return nil, err
	}

	return js, nil
}
