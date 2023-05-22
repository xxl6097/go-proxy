package openai

import (
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

type OpenAiProxy struct {
	_proxy *ApiProxy
	apikey string
	gocode string
	port   string
	keys   []string
}

func NewOpenAI() *OpenAiProxy {
	ai := &OpenAiProxy{
		_proxy: NewApiProxy(),
		keys:   make([]string, 0),
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
		codes := strings.Split(gocode, ",")
		for _, code := range codes {
			ai.keys = append(ai.keys, code)
		}
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

func in(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	//index的取值：[0,len(str_array)]
	if index < len(str_array) && str_array[index] == target { //需要注意此处的判断，先判断 &&左侧的条件，如果不满足则结束此处判断，不会再进行右侧的判断
		return true
	}
	return false
}
func (this *OpenAiProxy) onIntercepter(w http.ResponseWriter, r *http.Request) bool {
	authorization := r.Header.Get("Go-Authorization")
	if authorization == "" {
		authorization = r.Header.Get("go-authorization")
	}
	isIn := in(authorization, this.keys)
	fmt.Println("onIntercepter", isIn, r.Header)
	//if authorization == "" || !isIn {
	//	err := OpenAiError{
	//		Message: "请填写有效授权码",
	//		Type:    "invalid_request_error",
	//		Code:    "invalid_api_key",
	//	}
	//	buf, _ := json.Marshal(&err)
	//	fmt.Println("-->", string(buf))
	//	_, _ = w.Write(buf)
	//	return true
	//}
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
