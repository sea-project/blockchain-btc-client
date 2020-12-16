package sign

import (
	"errors"
)

func SignRawTx(rawTx, privateKeyWif string, networkType string) (signedRawTx string, err error) {
	msg := new(CustomHexMsg)
	err = msg.UnmarshalJSON(rawTx)
	if err != nil {
		return
	}
	msg.PrivKeys = &[]string{privateKeyWif}
	if msg.Flags == nil {
		var flagALL = "ALL"
		msg.Flags = &flagALL
	}
	signCmd := &SignRawTransactionCmd{
		RawTx:    msg.RawTx,
		Inputs:   msg.Inputs,
		PrivKeys: msg.PrivKeys,
		Flags:    msg.Flags,
	}
	conf, err := ParseNetworkToConf(networkType)
	result, err := SignRawTransaction(signCmd, conf)
	if err != nil {
		return
	}
	if result.Errors != nil && len(result.Errors) > 0 {
		errs := ""
		for _, e := range result.Errors {
			errs = errs + e.Error
		}
		err = errors.New(errs)
		return
	}
	signedRawTx = result.Hex
	return
}
