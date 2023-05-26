package gproxy

import (
	"fmt"
	"go-openai/middle"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy struct {
	host        string
	protocol    string
	apipath     string
	header      map[string]string
	interceptor func(w http.ResponseWriter, r *http.Request) bool
}

func NewProxy() *Proxy {
	this := &Proxy{
		header: make(map[string]string),
	}
	return this
}
func (this *Proxy) CreateProxy() (func(http.ResponseWriter, *http.Request), error) {
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

func (this *Proxy) modifyResponse() func(*http.Response) error {
	return func(resp *http.Response) error {
		return nil
	}
}

func (this *Proxy) modifyRequest(req *http.Request) {
	req.Host = this.host
	for k, v := range this.header {
		req.Header.Set(k, v)
	}
}

func (this *Proxy) ListenProxy(address, path string, handler func(http.ResponseWriter, *http.Request)) {
	// handle all requests to your server using the proxy
	//http.HandleFunc(path, handler)
	http.Handle(path, middle.EnableCors(middle.HandleOptions(http.HandlerFunc(handler))))
	//http.Handle("/conv", middle.EnableCors(middle.HandleOptions(middle.AuthMiddleware(http.HandlerFunc(api.HandleConv)))))
	log.Fatal(http.ListenAndServe(address, nil))
}

func (this *Proxy) proxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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

func (this *Proxy) errorHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, req *http.Request, err error) {
		fmt.Printf("Got error while modifying response: %v \n", err)
		return
	}
}

func (this *Proxy) SetInterceptor(fun func(w http.ResponseWriter, r *http.Request) bool) {
	this.interceptor = fun
}
func (this *Proxy) SetHost(host string) {
	this.host = host
}
func (this *Proxy) GetHost() string {
	return this.host
}
func (this *Proxy) SetProtocol(protocol string) {
	this.protocol = protocol
}
func (this *Proxy) SetApiPath(path string) {
	this.apipath = path
}
func (this *Proxy) AddHeader(key, value string) {
	this.header[key] = value
}
