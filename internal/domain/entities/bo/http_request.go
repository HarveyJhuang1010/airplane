package bo

import "net/http"

type HttpRequest struct {
	Method    string
	URL       string
	ProxyURL  *string
	ProxyType *string
	Headers   map[string]string
	Params    map[string]string
	Cookies   map[string]string
	Body      []byte
}

type HttpResponse struct {
	Headers    http.Header
	Cookies    []*http.Cookie
	StatusCode int
	Body       []byte
}
