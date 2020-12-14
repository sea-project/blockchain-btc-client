package external_api

import (
	"encoding/json"
	"fmt"
	"github.com/sea-project/blockchain-btc-client/model"
	"github.com/sea-project/blockchain-btc-client/request/httpclient"
	"github.com/sea-project/sea-pkg/util/decimal"
	"strconv"
)

// GetBalance 获取指定地址BTC余额
func GetBalance(address string, confirmations int) (string, error) {
	url := ""
	if confirmations != 0 {
		url = "https://blockchain.info/q/addressbalance/"+address+"/?confirmations="+strconv.Itoa(confirmations)
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

// ListUnspent 获取所有UTXO信息
func ListUnspent(address string, confirmations int) ([]*model.ListUnspentResult, error) {
	url := ""
	if confirmations != 0 {
		url = "https://blockchain.info/unspent?active="+address+"/?confirmations="+strconv.Itoa(confirmations)
	} else {
		url = "https://blockchain.info/unspent?active="+address
	}
	client := httpclient.NewClient(url)
	result, err := client.HttpGetRequest()
	if err != nil {
		return nil, fmt.Errorf("ListUnspent client.HttpGetRequest err:%v", err.Error())
	}
	allUTXOInfoTemp := make([]*model.RespFromListUnspentExternalAPI, 0)
	err = json.Unmarshal(result, &allUTXOInfoTemp)
	if err != nil {
		return nil, fmt.Errorf("ListUnspent json.Unmarshal err:%v", err.Error())
	}
	allUTXOInfo := make([]*model.ListUnspentResult, 0)
	for i := 0; i < len(allUTXOInfoTemp); i++ {
		utxoInfo := new(model.ListUnspentResult)
		utxoInfo.TxID = allUTXOInfoTemp[i].TxHash
		utxoInfo.Vout = allUTXOInfoTemp[i].TxOutPutN
		//valueDec := decimal.New(allUTXOInfoTemp[i].Value, 0).Div(decimal.New(100000000, 0))
		strValue := decimal.New(allUTXOInfoTemp[i].Value, 0).Div(decimal.New(100000000, 0)).String()
		utxoInfo.Amount, err = strconv.ParseFloat(strValue, 64)
		if err != nil {
			return nil, fmt.Errorf("ListUnspent strconv.ParseFloat value:%v err:%v", strValue, err.Error())
		}
		utxoInfo.ScriptPubKey = allUTXOInfoTemp[i].Script
		utxoInfo.Confirmations = allUTXOInfoTemp[i].Confirmations

		allUTXOInfo = append(allUTXOInfo, utxoInfo)
	}
	return allUTXOInfo, nil
}
