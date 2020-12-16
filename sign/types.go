package sign

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"strings"
)

// WalletTx 钱包可读数据
/*
* version : 2
* locktime : 0
* vin : [{"txid":"07d25a5793dd24cd6d1a810d8bb2958c271ed1971d7e1fb823217a1947170fed","output":0,"sequence":4294967295,"address":"38pfGw2jtkRvwJqXYqLtcFbPS7gHmkWfsN"}]
* vout : [{"address":"38pfGw2jtkRvwJqXYqLtcFbPS7gHmkWfsN","amount":0.084},{"address":"1QLGpxXUfJUVfVNDUJsuQ4dxBppgeuGX5R","amount":0.1}]
 */
type WalletTx struct {
	Version  int32  `json:"version"`
	Locktime uint32 `json:"locktime"`
	Vin      []vin  `json:"vin"`
	Vout     []vout `json:"vout"`
}

type vin struct {
	Txid     string `json:"txid"`
	Vout     uint32 `json:"output"`
	Sequence uint32 `json:"sequence"`
	Address  string `json:"address"`
}

type vout struct {
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}

// 服务端发送的消息
// "{hex:"",rawTxInput:[{scriptPubKey:"xxx",redeemScript:"xxx",amount:0}]}"
type CustomHexMsg struct {
	SignRawTransactionCmd
	DecodeTransaction func(cmd *btcjson.DecodeRawTransactionCmd, params *chaincfg.Params) (
		btcjson.TxRawDecodeResult, error) `json:"-,omitempty"`
	SignTransaction func()   `json:"-,omitempty"`
	walletTx        WalletTx `json:"-"` // covert from txRawDecodeResult
}

func (c *CustomHexMsg) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
}

func (c *CustomHexMsg) UnmarshalJSON(msg string) (err error) {
	msg = strings.TrimPrefix(msg, "0x")
	data, err := hex.DecodeString(msg)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &c)
	if err != nil {
		return
	}
	return
}

func (c *CustomHexMsg) getAddressesFromScriptPubKey(txId string, chainCfg *chaincfg.Params) (addresses []string, err error) {
	if c.Inputs == nil {
		err = errors.New("c.Inputs is nil")
		return
	}
	for item := range *c.Inputs {
		if (*c.Inputs)[item].Txid == txId {
			var reply btcjson.DecodeScriptResult
			reply, err = DecodeScript(&btcjson.DecodeScriptCmd{HexScript: (*c.Inputs)[item].ScriptPubKey}, chainCfg)
			if err != nil {
				return
			}
			addresses = reply.Addresses
		}
	}
	return
}

func (c *CustomHexMsg) MarshalToWalletTxJSON(chainCfg *chaincfg.Params) (tx string, err error) {
	if c.DecodeTransaction == nil {
		err = errors.New("decodeTransaction func not set")
		return
	}
	result, err := c.DecodeTransaction(&btcjson.DecodeRawTransactionCmd{HexTx: c.RawTx}, chainCfg)
	if err != nil {
		return
	}
	c.walletTx.Version = result.Version
	c.walletTx.Locktime = result.Locktime
	for item := range result.Vin {
		var in vin
		var addresses []string
		addresses, err = c.getAddressesFromScriptPubKey(result.Vin[item].Txid, chainCfg)

		if len(addresses) != 1 {
			err = errors.New("decode addresses len not equal 1")
			return
		}
		in.Txid = result.Vin[item].Txid
		in.Vout = result.Vin[item].Vout
		in.Sequence = result.Vin[item].Sequence
		in.Address = addresses[0]
		c.walletTx.Vin = append(c.walletTx.Vin, in)
	}
	for item := range result.Vout {
		var out vout
		if addresses := result.Vout[item].ScriptPubKey.Addresses; len(addresses) != 1 {
			// op_return + omni
			// 6a146f6d6e69
			continue
		} else {
			out.Address = addresses[0]
		}
		out.Amount = result.Vout[item].Value
		c.walletTx.Vout = append(c.walletTx.Vout, out)
	}

	data, err := json.Marshal(c.walletTx)
	if err != nil {
		return
	}
	tx = string(data)
	return
}

// DecodeScript handles decodescript commands.
func DecodeScript(c *btcjson.DecodeScriptCmd, chainCfg *chaincfg.Params) (reply btcjson.DecodeScriptResult, err error) {

	// Convert the hex script to bytes.
	hexStr := c.HexScript
	script, err := hex.DecodeString(hexStr)
	if err != nil {
		return
	}

	// The disassembled string will contain [error] inline if the script
	// doesn't fully parse, so ignore the error here.
	disbuf, _ := txscript.DisasmString(script)

	// Get information about the script.
	// Ignore the error here since an error means the script couldn't parse
	// and there is no additinal information about it anyways.
	scriptClass, addrs, reqSigs, _ := txscript.ExtractPkScriptAddrs(script, chainCfg)
	addresses := make([]string, len(addrs))
	for i, addr := range addrs {
		addresses[i] = addr.EncodeAddress()
	}

	// Convert the script itself to a pay-to-script-hash address.
	p2sh, err := btcutil.NewAddressScriptHash(script, chainCfg)
	if err != nil {
		return
	}

	// Generate and return the reply.
	reply = btcjson.DecodeScriptResult{
		Asm:       disbuf,
		ReqSigs:   int32(reqSigs),
		Type:      scriptClass.String(),
		Addresses: addresses,
	}

	if scriptClass != txscript.ScriptHashTy {
		reply.P2sh = p2sh.EncodeAddress()
	}
	return reply, nil
}
