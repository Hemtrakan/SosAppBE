package httpclient

import (
	"bytes"
	"io"
	"net/http"
)

func (h HttpClient) send(method, reqURL, body string, httpHeaderMap map[string]string) (httpResp HttpResponse, Error error) {
	var httpReq *http.Request
	bodyBuffer := bytes.NewBufferString(body)
	httpReq, err := http.NewRequest(method, reqURL, bodyBuffer)
	if err != nil {
		Error = err
		return
	}

	if httpHeaderMap != nil {
		for k, v := range httpHeaderMap {
			httpReq.Header.Set(k, v)
		}
	}

	httpReq.Header.Set("Cache-Control", "no-cache")

	client := &http.Client{
		Timeout: h.Timeout,
	}

	var resp *http.Response
	resp, err = client.Do(httpReq)
	if err != nil {
		Error = err
		return
	}

	defer resp.Body.Close()

	var rawBody []byte
	rawBody, err = io.ReadAll(resp.Body)
	if err != nil {
		Error = err
		return
	}

	httpResp.HttpStatusCode = resp.StatusCode
	httpResp.HttpStatusMsg = resp.Status
	httpResp.ResponseMsg = string(rawBody)
	httpResp.HttpHeader = resp.Header
	return
}

func (h HttpClient) Get(url string, httpHeaderMap map[string]string) (httpResp HttpResponse, Error error) {
	return h.send("GET", url, "", httpHeaderMap)

}

func (h HttpClient) PostJson(url, body string, httpHeaderMap map[string]string) (httpResp HttpResponse, Error error) {
	if httpHeaderMap == nil {
		httpHeaderMap = make(map[string]string)
	}
	if httpHeaderMap != nil {
		httpHeaderMap["Content-Type"] = "application/json"
	}

	return h.send("POST", url, body, httpHeaderMap)
}

func (h HttpClient) PutJson(url, body string, httpHeaderMap map[string]string) (httpResp HttpResponse, Error error) {
	if httpHeaderMap == nil {
		httpHeaderMap = make(map[string]string)
	}
	if httpHeaderMap != nil {
		httpHeaderMap["Content-Type"] = "application/json"
	}

	return h.send("PUT", url, body, httpHeaderMap)
}

func (h HttpClient) DeleteJson(url, body string, httpHeaderMap map[string]string) (httpResp HttpResponse, Error error) {
	return h.send("DELETE", url, body, httpHeaderMap)
}
