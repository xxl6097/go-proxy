package openai

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type OpenAiError struct {
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Param   interface{} `json:"param"`
	Code    string      `json:"code"`
}

type Error struct {
	OpenAiError
}

type OpenAiProxy struct {
	_proxy *ApiProxy
	apikey string
	gocode string
	port   string
}

func NewOpenAI() *OpenAiProxy {
	ai := &OpenAiProxy{
		_proxy: NewApiProxy(),
	}
	apikey := os.Getenv("OPENAI_APIKEY")
	if apikey != "" {
		ai.apikey = "sk-" + apikey
	} else {
		//ai.apikey = "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"
		panic("please set apikey in env")
	}

	gocode := os.Getenv("GO_APIKEY")
	if gocode != "" {
		ai.gocode = "clife-" + gocode
	} else {
		panic("please set gocode in env")
	}
	port := os.Getenv("SERVER_PORT")
	if port != "" {
		ai.port = port
	} else {
		panic("please set port in env")
	}

	fmt.Println("NewOpenAI", ai)
	return ai
}

func (this *OpenAiProxy) onIntercepter(w http.ResponseWriter, r *http.Request) bool {
	fmt.Println("onIntercepter", r.Header)
	authorization := r.Header.Get("Go-Authorization")
	if authorization == "" || authorization != this.gocode {
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
}

func (this *OpenAiProxy) CreateOpenAiProxy() {
	this._proxy.SetHost("api.openai.com")
	this._proxy.SetProtocol("https")
	this._proxy.SetInterceptor(this.onIntercepter)
	this._proxy.AddHeader("Authorization", fmt.Sprintf("Bearer %s", this.apikey))
	handle, err := this._proxy.CreateProxy()
	if err != nil {
		panic(err)
	}
	this._proxy.ListenProxy(fmt.Sprintf(":%s", this.port), "/", handle)
}

func (this *OpenAiProxy) GetTestHandle() {
	this._proxy.SetHost("api.uuxia.cn")
	this._proxy.SetInterceptor(this.onIntercepter)
	this._proxy.SetProtocol("http")
	this._proxy.AddHeader("Go-Authorization", "xiaxiaoli")
	this._proxy.AddHeader("Go-Token", "5e1df951de82a151068a8a96c35fc84d")
	this._proxy.AddHeader("Host", this._proxy.GetHost())
	handle, err := this._proxy.CreateProxy()
	if err != nil {
		panic(err)
	}
	this._proxy.ListenProxy(fmt.Sprintf(":%s", this.port), "/", handle)
}
