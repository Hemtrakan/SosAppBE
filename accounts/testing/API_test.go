package testing

import (
	"accounts/httpclient"
	"testing"
)

var (
	h            httpclient.HttpClient
	HttpResponse httpclient.HttpResponse
	Error        error
)

func TestGetAPI(t *testing.T) {

	URL := "http://34.124.232.197:80/SosApp/accounts/user/14"
	HttpResponse, Error = h.Get(URL, nil)
	if Error != nil {
		t.Error("err : ", Error)
	}

	t.Log(HttpResponse.ResponseMsg)
	t.Log(HttpResponse.HttpStatusCode)
	t.Log(HttpResponse.HttpHeader)
}

func TestPostAPI(t *testing.T) {

}

func TestPutAPI(t *testing.T) {

}

func TestDeleteAPI(t *testing.T) {

}
