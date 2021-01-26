package model

type RespOMNIGetBalance struct {
	Balance		string	`json:"balance"`
	Reserved	string	`json:"reserved"`
	Frozen		string	`json:"frozen"`
}

type RespOMNIGetTransaction struct {
	TxID 		string	`json:"txid"`
	Fee 		string	`json:"fee"`
	SendingAddress string	`json:"sendingaddress"`
	IsMine 		bool	`json:"ismine"`
	Version 	uint32	`json:"version"`
	TypeInt		uint32	`json:"type_int"`
	Type 		string	`json:"type"`
	PropertyID  uint64	`json:"propertyid"`
	Divisible	bool	`json:"divisible"`
	Ecosystem	string	`json:"ecosystem"`
	PropertyType	string	`json:"propertytype"`
	Category	string	`json:"category"`
	Subcategory	string	`json:"subcategory"`
	PropertyName	string	`json:"propertyname"`
	Data		string	`json:"data"`
	Url 		string	`json:"url"`
	Amount 		string	`json:"amount"`
	Valid 		bool	`json:"valid"`
	BlockHash	string	`json:"blockhash"`
	BlockTime 	uint64	`json:"blocktime"`
	Positioninblock	uint64	`json:"positioninblock"`
	Block 		uint64	`json:"block"`
	Confirmations	uint64	`json:"confirmations"`
}
