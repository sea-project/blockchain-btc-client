package model

// RequestFromGetBalance 获取BTC地址余额请求参数
type RequestFromGetAddrBalance struct {
	Confirmations	int		`json:"confirmations"`
}
