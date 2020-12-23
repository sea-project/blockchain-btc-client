package model

import "encoding/json"

type BlockFee struct {
	FastestFee	int	`json:"fastestFee"`
	HalfHourFee	int	`json:"halfHourFee"`
	HourFee		int	`json:"hourFee"`
}

func (b *BlockFee) String() string {
	str, _ := json.Marshal(b)
	return string(str)
}
