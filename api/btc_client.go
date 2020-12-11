package api

import (
	"encoding/json"
	"fmt"
	"github.com/sea-project/blockchain-btc-client/model"
	"github.com/sea-project/blockchain-btc-client/pkg/constant"
	"github.com/sea-project/blockchain-btc-client/request"
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
