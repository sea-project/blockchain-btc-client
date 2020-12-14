package model

type RespOMNIGetBalance struct {
	Balance		string	`json:"balance"`
	Reserved	string	`json:"reserved"`
	Frozen		string	`json:"frozen"`
}
