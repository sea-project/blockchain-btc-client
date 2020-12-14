package model

type BlockFee struct {
	FastestFee	int	`json:"fastestFee"`
	HalfHourFee	int	`json:"halfHourFee"`
	HourFee		int	`json:"hourFee"`
}
