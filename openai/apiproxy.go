package openai

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ApiProxy struct {
	host        string
	protocol    string
	apipath     string
	header      map[string]string
	interceptor func(w http.ResponseWriter, r *http.Request) bool
}

func NewApiProxy() *ApiProxy {
	this := &ApiProxy{
		header: make(map[string]string),
	}
	return this
}
func (this *ApiProxy) CreateProxy() (func(http.ResponseWriter, *http.Request), error) {
	url, err := url.Parse(fmt.Sprintf("%s://%s%s", this.protocol, this.host, this.apipath))
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(url)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		this.modifyRequest(req)

	}
	proxy.ModifyResponse = this.modifyResponse()
	proxy.ErrorHandler = this.errorHandler()
	return this.proxyRequestHandler(proxy), nil
}

func (svr *ApiProxy) setAllows(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, "+
		"Go-Authorization, Go-Token, vcode, authorization, Origin, Referer, Sec-Fetch-Mode, User-Agent")
}

func (this *ApiProxy) modifyResponse() func(*http.Response) error {
	return func(resp *http.Response) error {
		return nil
	}
}

func (this *ApiProxy) modifyRequest(req *http.Request) {
	req.Host = this.host
	for k, v := range this.header {
		req.Header.Set(k, v)
	}
}

func (this *ApiProxy) ListenProxy(address, path string, handler func(http.ResponseWriter, *http.Request)) {
	// handle all requests to your server using the proxy
	http.HandleFunc(path, handler)
	log.Fatal(http.ListenAndServe(address, nil))
}

func (this *ApiProxy) proxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		this.setAllows(&w)
		isIntercept := false
		if this.interceptor != nil {
			isIntercept = this.interceptor(w, r)
		}
		this.modifyRequest(r)
		if !isIntercept {
			proxy.ServeHTTP(w, r)
		}

	}
}

func (this *ApiProxy) errorHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, req *http.Request, err error) {
		fmt.Printf("Got error while modifying response: %v \n", err)
		return
	}
}

func (this *ApiProxy) SetInterceptor(fun func(w http.ResponseWriter, r *http.Request) bool) {
	this.interceptor = fun
}
func (this *ApiProxy) SetHost(host string) {
	this.host = host
}
func (this *ApiProxy) GetHost() string {
	return this.host
}
func (this *ApiProxy) SetProtocol(protocol string) {
	this.protocol = protocol
}
func (this *ApiProxy) SetApiPath(path string) {
	this.apipath = path
}
func (this *ApiProxy) AddHeader(key, value string) {
	this.header[key] = value
}
