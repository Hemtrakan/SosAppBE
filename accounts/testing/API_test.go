package testing

import (
	"accounts/httpclient"
	config "github.com/spf13/viper"
	"testing"
)

var (
	h            httpclient.HttpClient
	HttpResponse httpclient.HttpResponse
	Error        error
)

func TestGetAPI(t *testing.T) {
	account := config.GetString("url.account")
	URL := account + "user/14"
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
