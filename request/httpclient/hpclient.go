package httpclient

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Client
type Client struct {
	url string
	client *http.Client
}

// NewClient 初始化客户端
func NewClient(url string) *Client {
	return &Client{
		url:    url,
		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 10,  // 每台主机保持的最大空闲连接
				ResponseHeaderTimeout: 10 * time.Second,
			},
			Timeout: 5 * time.Second,
		},
	}
}

// httpRequest http请求
func (c *Client) HttpRequest(params string) (resBody []byte, err error) {

	// 初始化请求
	body := strings.NewReader(params)
	req, err := http.NewRequest("POST", c.url, body)
	if err != nil {
		return nil, errors.Wrap(err, "Http NewRequest")
	}

	// 执行请求
	req.Header.Add("Content-Type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Client Do")
	}
	defer res.Body.Close()

	// 接收返回结果
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll")
	}
	return resBody, nil
}

// httpRequest http请求
func (c *Client) HttpGetRequest() (resBody []byte, err error) {

	// 初始化请求
	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Http NewRequest")
	}

	// 执行请求
	req.Header.Add("Content-Type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Client Do")
	}
	defer res.Body.Close()

	// 接收返回结果
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll")
	}
	return resBody, nil
}

