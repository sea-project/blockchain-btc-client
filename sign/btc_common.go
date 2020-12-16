// description: keybox
//
// @author: xuwenchao
// @date: 2020/8/21 0021
package sign

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
)

func ParseNetworkToConf(network string) (*chaincfg.Params, error) {
	switch network {
	case "mainnet":
		return &chaincfg.MainNetParams, nil
	case "regtest":
		return &chaincfg.RegressionNetParams, nil
	case "testnet":
		return &chaincfg.RegressionNetParams, nil
	case "testnet3":
		return &chaincfg.TestNet3Params, nil
	case "simnet":
		return &chaincfg.SimNetParams, nil
	default:
		return nil, fmt.Errorf("network is error")
	}
}