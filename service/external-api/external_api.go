package external_api

import (
	"fmt"
	"github.com/sea-project/blockchain-btc-client/request/httpclient"
	"github.com/sea-project/sea-pkg/util/decimal"
	"strconv"
)

// GetBalance 获取指定地址BTC余额
func GetBalance(address string, confirmations int) (string, error) {
	url := ""
	if confirmations != 0 {
		url = "https://blockchain.info/q/addressbalance/"+address+"/?"+strconv.Itoa(confirmations)
	} else {
		url = "https://blockchain.info/q/addressbalance/"+address
	}
	client := httpclient.NewClient(url)
	result, err := client.HttpGetRequest()
	if err != nil {
		return "", fmt.Errorf("GetBalance client.HttpGetRequest err:%v", err.Error())
	}
	// 转换成BTC，需要除以10^8
	resultDec, err := decimal.NewFromString(string(result))
	if err != nil {
		return "", fmt.Errorf("GetBalance decimal.NewFromString err:%v", err.Error())
	}
	balanceDec := resultDec.Div(decimal.New(100000000, 0))
	return balanceDec.String(), nil
}
