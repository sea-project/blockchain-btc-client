package api

import (
	"encoding/json"
	"fmt"
	"github.com/sea-project/blockchain-btc-client/model"
	"github.com/sea-project/blockchain-btc-client/pkg/constant"
	"github.com/sea-project/blockchain-btc-client/request"
	external_api "github.com/sea-project/blockchain-btc-client/service/external-api"
	"github.com/sea-project/blockchain-btc-client/sign"
	"github.com/sea-project/sea-pkg/util/decimal"
	"strconv"
)

// BTCClient 比特币客户端
type BTCClient struct {
	client 		*request.HttpClient
	clientType 	int
}

// NewBTCClient 创建比特币客户端实例
func NewBTCClient(url, userName, userPwd string, clientType int) *BTCClient {
	btcClient := new(BTCClient)
	httpClient := request.NewHTTPClient(url, userName, userPwd)
	btcClient.client = httpClient
	btcClient.clientType = clientType
	return btcClient
}

// GetBlockCount 获取最新区块高度
func (c *BTCClient) GetBlockCount() (uint64, error) {
	result, err := c.client.HttpRequest(constant.GetBlockCount, nil)
	if err != nil {
		return 0, err
	}
	strHeight := fmt.Sprintf("%v", result)
	height, err := strconv.ParseUint(strHeight, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("GetBlockCount strHeight Parse failed strHeight:%v", strHeight)
	}
	return height, nil
}

// GetBlockHash 获取指定高度的区块哈希值
func (c *BTCClient) GetBlockHash(height uint64) (string, error) {
	params := make([]interface{}, 0)
	params = append(params, height)
	result, err := c.client.HttpRequest(constant.GetBlockHash, params)
	if err != nil {
		return "", err
	}
	if blockHash, ok := result.(string); ok {
		return blockHash, nil
	}
	return "", fmt.Errorf("GetBlockHash blockHash Parse failed")
}

// GetBlock 获取指定哈希的区块信息
func (c *BTCClient) GetBlock(blockHash string) (*model.GetBlockResultV2, error) {
	params := make([]interface{}, 0)
	// 第二个参数传2表示解码所有交易信息
	format := 2
	params = append(params, blockHash)
	params = append(params, format)
	result, err := c.client.HttpRequest(constant.GetBlock, params)
	if err != nil {
		return nil, err
	}

	// 返回数据类型转换
	temp, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("GetBlock result json.Marshal failed err:%v", err.Error())
	}
	blockInfo := new(model.GetBlockResultV2)
	err = json.Unmarshal(temp, &blockInfo)
	if err != nil {
		return nil, fmt.Errorf("GetBlock result json.Unmarshal failed err:%v", err.Error())
	}

	return blockInfo, nil
}

// GetBlockByHeight 通过区块高度获取区块详细信息
func (c *BTCClient) GetBlockByHeight(height uint64) (*model.GetBlockResultV2, error) {
	// 根据区块高度获取区块哈希
	blockHash, err := c.GetBlockHash(height)
	if err != nil {
		return nil, err
	}
	// 根据区块高度获取区块信息
	return c.GetBlock(blockHash)
}

// GetRawTransaction 获取交易详情
func (c *BTCClient) GetRawTransaction(txid string, format bool) (interface{}, error) {
	params := make([]interface{}, 0)
	params = append(params, txid)
	params = append(params, format)
	result, err := c.client.HttpRequest(constant.GetRawTransaction, params)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ImportAddress 导入地址
func (c *BTCClient) ImportAddress(addressOrScript, account string, rescan bool) error {
	params := make([]interface{}, 0)
	params = append(params, addressOrScript)
	params = append(params, account)
	params = append(params, rescan)
	_, err := c.client.HttpRequest(constant.ImportAddress, params)
	if err != nil {
		return err
	}
	return nil
}

// ListUnspent 返回归属于本钱包的未消费交易输出数组
func (c *BTCClient) ListUnspent(address string, minimumConfirmations int) ([]*model.ListUnspentResult, error) {
	if c.clientType == TypeExternalAPI {
		// 该种方式是调用外部接口类型
		return external_api.ListUnspent(address, minimumConfirmations)
	} else if c.clientType == TypeAddListen {
		params := make([]interface{}, 0)
		params = append(params, minimumConfirmations)
		params = append(params, 9999999)
		addresses := make([]string, 0)
		addresses = append(addresses, address)
		params = append(params, addresses)
		result, err := c.client.HttpRequest(constant.ListUnspent, params)
		if err != nil {
			return nil, err
		}
		// 返回数据类型转换
		temp, err := json.Marshal(result)
		if err != nil {
			return nil, fmt.Errorf("ListUnspent result json.Marshal failed err:%v", err.Error())
		}
		listUnspentResult := make([]*model.ListUnspentResult, 0)
		err = json.Unmarshal(temp, &listUnspentResult)
		if err != nil {
			return nil, fmt.Errorf("ListUnspent result json.Unmarshal failed err:%v", err.Error())
		}
		return listUnspentResult, nil
	}
	return nil, fmt.Errorf("ListUnspent type not support")
}

// GetBalance 获取比特币余额
func (c *BTCClient) GetBalance(address string, confirmations int) (string, error) {
	if c.clientType == TypeExternalAPI {
		// 该种方式是调用外部接口类型
		return external_api.GetBalance(address, confirmations)
	} else if c.clientType == TypeAddListen {
		// 该种方式是向节点添加监听但不扫描之前交易类型
		// 获取某个地址所有UTXO
		allUTXOInfo, err := c.ListUnspent(address, 1)
		if err != nil {
			return "", fmt.Errorf("GetBalance c.ListUnspent err:%v", err)
		}
		balanceDec := decimal.NewFromFloat(0)
		for i := 0; i < len(allUTXOInfo); i++ {
			balanceDec = balanceDec.Add(decimal.NewFromFloat(allUTXOInfo[i].Amount))
		}
		return balanceDec.String(), nil
	}
	return "", fmt.Errorf("GetBalance type not support")
}

// SendRawTransaction 广播裸交易
func (c *BTCClient) SendRawTransaction(transaction string) (string, error) {
	params := make([]interface{}, 0)
	params = append(params, transaction)
	result, err := c.client.HttpRequest(constant.SendRawTransaction, params)
	if err != nil {
		return "", err
	}
	if txid, ok := result.(string); ok {
		return txid, nil
	}
	return "", fmt.Errorf("SendRawTransaction txid Parse failed")
}

// GetBlockFeeLatest 获取最新手续费信息
func (c *BTCClient) GetBlockFeeLatest() (*model.BlockFee, error) {
	if c.clientType == TypeExternalAPI || c.clientType == TypeAddListen {
		return external_api.GetBlockFeeLatest()
	}
	return nil, fmt.Errorf("GetBlockFeeLatest type not support")
}

// GetOMNIBalance 获取omni账户的余额
func (c *BTCClient) GetOMNIBalance(address string, propertyid int) (string, error) {
	params := make([]interface{}, 0)
	params = append(params, address)
	params = append(params, propertyid)
	result, err := c.client.HttpRequest(constant.OmniGetBalance, params)
	if err != nil {
		return "", err
	}

	// 返回数据类型转换
	temp, err := json.Marshal(result)
	if err != nil {
		return "", fmt.Errorf("GetOMNIBalance result json.Marshal failed err:%v", err.Error())
	}
	respOMNIGetBalance := new(model.RespOMNIGetBalance)
	err = json.Unmarshal(temp, &respOMNIGetBalance)
	if err != nil {
		return "", fmt.Errorf("GetOMNIBalance result json.Unmarshal failed err:%v", err.Error())
	}
	return respOMNIGetBalance.Balance, nil
}

// GetOMNITransaction 获取指定Omni交易的详细信息
func (c *BTCClient) GetOMNITransaction(txid string) (interface{}, error) {
	params := make([]interface{}, 0)
	params = append(params, txid)
	return c.client.HttpRequest(constant.OmniGetTransaction, params)
}

// OMNICreatePayloadSimpleSend 调用创建一个用于简单发送交易的载荷
func (c *BTCClient) OMNICreatePayloadSimpleSend(propertyid uint32, amount string) (string, error) {
	params := make([]interface{}, 0)
	params = append(params, propertyid)
	params = append(params, amount)
	result, err := c.client.HttpRequest(constant.OMNICreatePayloadSimpleSend, params)
	if err != nil {
		return "", err
	}
	if payload, ok := result.(string); ok {
		return "6a146f6d6e69" + payload, nil
	}
	return "", fmt.Errorf("OMNICreatePayloadSimpleSend payload Parse failed")
}

// SignRawTx 签名交易
func (c *BTCClient) SignRawTx(rawTx, privateKeyWif string, networkType string) (string, error) {
	return sign.SignRawTx(rawTx, privateKeyWif, networkType)
}