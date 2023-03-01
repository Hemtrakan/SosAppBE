package httpclient

import (
	"net/http"
	"time"
)

type HttpClient struct {
	MaxConnections  int
	Timeout         time.Duration
	CertSkipVerify  bool
	CertServerName  string
	CertPEMFileName string
	Charset         string
	ProxyURL        string
}

type HttpResponse struct {
	HttpStatusCode int
	HttpStatusMsg  string
	ResponseMsg    string
	HttpHeader     http.Header
	IsRedirect     bool
	RedirectUrl    string
}
