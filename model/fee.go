package model

type BlockFee struct {
	FastestFee	uint64	`json:"fastestFee"`
	HalfHourFee	uint64	`json:"halfHourFee"`
	HourFee		uint64	`json:"hourFee"`
}
