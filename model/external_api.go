package model

// RequestFromGetBalance 获取BTC地址余额请求参数
type RequestFromGetAddrBalance struct {
	Confirmations	int		`json:"confirmations"`
}

type RespFromListUnspentExternalAPI struct {
	TxHash 			string		`json:"tx_hash"`
	TxOutPutN 		uint32		`json:"tx_output_n"`
	Script 			string		`json:"script"`
	Value 			int64		`json:"value"`
	Confirmations	int64		`json:"confirmations"`
	TxIndex			int			`json:"tx_index"`
}
