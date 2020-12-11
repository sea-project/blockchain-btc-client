package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sea-project/blockchain-btc-client/model"
	"github.com/sea-project/go-logger"
	"io/ioutil"
	"net/http"
)

// httpClient manager http client
type HttpClient struct {
	url   		string	`json:"url"`
	userName	string	`json:"user_name"`
	userPwd		string	`json:"user_pwd"`
}

func NewHTTPClient(url, userName, userPwd string) *HttpClient {
	httpClient := new(HttpClient)
	httpClient.url = url
	httpClient.userName = userName
	httpClient.userPwd = userPwd
	return httpClient
}

func (c *HttpClient) HttpRequest(method string, params []interface{}) (interface{}, error) {
	rawParams := make([]json.RawMessage, 0, len(params))
	for _, param := range params {
		marshalParam, _ := json.Marshal(param)
		rawParams = append(rawParams, json.RawMessage(marshalParam))
	}
	request := &model.BtcApiRequest{
		Jsonrpc: "1.0",
		ID:      "curltest",
		Method:  method,
		Params:  rawParams,
	}
	marshalRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(marshalRequest)
	httpReq, err := http.NewRequest("POST", c.url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("http NewRequest error:%v", err.Error())
	}

	httpReq.Close = true
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.SetBasicAuth(c.userName, c.userPwd)

	client := httpClient()
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("http client do error:%v", err.Error())
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HttpRequest resp.StatusCode err:%v,%v", resp.StatusCode, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll error:%v", err.Error())
	}
	responseDetail := new(model.BtcApiResponse)
	err = json.Unmarshal(body, &responseDetail)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal btc response error:%v", err.Error())
	}
	if responseDetail.Error != nil {
		logger.Error("btc error type", "code", responseDetail.Error.Code, "msg", responseDetail.Error.Message)
		return nil, fmt.Errorf(responseDetail.Error.Message)
	}
	return responseDetail.Result, nil
}

// 长连接client使用方法
func httpClient() http.Client {
	tr := http.Transport{DisableKeepAlives: true}
	return http.Client{Transport: &tr}
}
