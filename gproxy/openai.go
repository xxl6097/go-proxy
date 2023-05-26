package gproxy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
)

type OpenAiError struct {
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Param   interface{} `json:"param"`
	Code    string      `json:"code"`
}

func CreateOpenAiProxy(port string, openAiKey string, fun func(w http.ResponseWriter, r *http.Request) bool) {
	_proxy := NewProxy()
	_proxy.SetHost("api.openai.com")
	_proxy.SetProtocol("https")
	_proxy.SetInterceptor(fun)
	_proxy.AddHeader("Authorization", fmt.Sprintf("Bearer %s", openAiKey))
	handle, err := _proxy.CreateProxy()
	if err != nil {
		panic(err)
	}
	_proxy.ListenProxy(fmt.Sprintf(":%s", port), "/", handle)
}

func In(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	//index的取值：[0,len(str_array)]
	if index < len(str_array) && str_array[index] == target { //需要注意此处的判断，先判断 &&左侧的条件，如果不满足则结束此处判断，不会再进行右侧的判断
		return true
	}
	return false
}

func ObtainEnv() (string, []string, string) {
	apikey := os.Getenv("OPENAI_APIKEY")
	if apikey != "" {
		apikey = "sk-" + apikey
	} else {
		//ai.apikey = "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"
		panic("please set apikey in env")
	}

	gocode := os.Getenv("GO_APIKEY")
	keys := make([]string, 0)
	if gocode != "" {
		codes := strings.Split(gocode, ",")
		for _, code := range codes {
			keys = append(keys, code)
		}
	} else {
		panic("please set gocode in env")
	}
	port := os.Getenv("SERVER_PORT")
	if port != "" {
		port = port
	} else {
		panic("please set port in env")
	}

	return apikey, keys, port
}

func CreateOpenAiDefaultProxy(port string, openAiKey string, userkeys []string) {
	_proxy := NewProxy()
	_proxy.SetHost("api.openai.com")
	_proxy.SetProtocol("https")
	_proxy.SetInterceptor(func(w http.ResponseWriter, r *http.Request) bool {
		authorization := r.Header.Get("Go-Authorization")
		r.Header.Del("Origin")
		if authorization == "" {
			authorization = r.Header.Get("go-authorization")
		}
		isIn := In(authorization, userkeys)
		fmt.Println("onIntercepter", isIn, r.Header)
		if authorization == "" || !isIn {
			err := OpenAiError{
				Message: "请填写有效授权码",
				Type:    "invalid_request_error",
				Code:    "invalid_api_key",
			}
			buf, _ := json.Marshal(&err)
			fmt.Println("-->", string(buf))
			_, _ = w.Write(buf)
			return true
		}
		return false
	})
	_proxy.AddHeader("Authorization", fmt.Sprintf("Bearer %s", openAiKey))
	handle, err := _proxy.CreateProxy()
	if err != nil {
		panic(err)
	}
	_proxy.ListenProxy(fmt.Sprintf(":%s", port), "/", handle)
}
