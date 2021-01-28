package api

import "testing"

func TestBTCClient_GetBlockCount(t *testing.T) {
	client := NewBTCClient("http://8.210.178.221:8332/", "admin", "Q3Az6XHNdE", TypeAddListen)
	num, err := client.GetBlockCount()
	if err != nil {
		t.Errorf("num:%v", num)
		return
	}
	t.Logf("num:%v", num)
}

func TestBTCClient_GetBlockHeader(t *testing.T) {
	client := NewBTCClient("http://47.75.116.218:8332/", "admin", "123456", TypeAddListen)
	blockHash := "01d2e5d11420a3961617959c90e240d0ffa33e705c6c1dd1e55720bb0cfff229"
	result, err := client.GetBlockHeader(blockHash)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	t.Logf("result:%v", result)
}

func TestBTCClient_GetRawTransaction(t *testing.T) {
	client := NewBTCClient("http://8.210.195.246:8332/", "admin", "123456", TypeAddListen)
	txid := "3f2952a1afd495fa170b84b3432fa027e546ab36ab1bd504290a14a6ee31bfd0"
	result, err := client.GetRawTransaction(txid, true)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	t.Logf("result:%v", result)
}

func TestBTCClient_ImportAddress(t *testing.T) {
	client := NewBTCClient("http://47.75.116.218:8332/", "admin", "123456", TypeAddListen)
	address := "n4MiXKhTD69pzxq2Mr7onaPfNXv5KLm5Rb"
	err := client.ImportAddress(address, address, false)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	t.Logf("success")
}

func TestBTCClient_ListUnspent(t *testing.T) {
	client := NewBTCClient("http://47.75.116.218:8332/", "admin", "123456", TypeAddListen)
	//addresses := make([]string, 0)
	//addresses = append(addresses, "n4MiXKhTD69pzxq2Mr7onaPfNXv5KLm5Rb")
	address := "n4MiXKhTD69pzxq2Mr7onaPfNXv5KLm5Rb"
	result, err := client.ListUnspent(address, 6)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	for i := 0; i < len(result); i++ {
		t.Logf("result:%v", result[i])
	}
}

func TestBTCClient_GetBlockFeeLatest(t *testing.T) {
	client := NewBTCClient("http://47.75.116.218:8332/", "admin", "123456", TypeAddListen)
	result, err := client.GetBlockFeeLatest()
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	t.Logf("result:%v", result)
}

func TestBTCClient_GetOMNIBalance(t *testing.T) {
	client := NewBTCClient("http://8.210.178.221:8332/", "admin", "Q3Az6XHNdE", TypeAddListen)
	address := "3GyeFJmQynJWd8DeACm4cdEnZcckAtrfcN"
	num, err := client.GetOMNIBalance(address, 31)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	t.Logf("num:%v", num)
}

// 测试调用创建一个用于简单发送交易的载荷
func TestBTCClient_OMNICreatePayloadSimpleSend(t *testing.T) {
	client := NewBTCClient("http://8.210.178.221:8332/", "admin", "Q3Az6XHNdE", TypeAddListen)
	var propertyid uint32 = 31
	amount := "2"
	num, err := client.OMNICreatePayloadSimpleSend(propertyid, amount)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	t.Logf("num:%v", num)
}

func TestBTCClient_SignRawTx(t *testing.T) {
	client := new(BTCClient)
	//rawTx := "7b225261775478223a22303130303030303030316136613162333734623232633563353564313864396165656438356533613931333438646338326431303431383833643537646632666462346432366432313930323030303030303030666666666666303030333030303030303030303030303030303031363661313436663664366536393030303030303030383030303030303330303030303030303031333132643030323230323030303030303030303030303139373661393134316338653962626263663731653364363662343061636536633832303362373266393631353530373838616339613735353330323030303030303030313937366139313435383531306536316564323064656338393837333437326330303036383334303736643963663262383861633030303030303030222c22496e70757473223a5b7b2274786964223a2231396432323634646462326664663537336438383431313032646338386433343931336135656438656539613864643135353563326362323734623361316136222c22766f7574223a322c227363726970745075624b6579223a223736613931343538353130653631656432306465633839383733343732633030303638333430373664396366326238386163222c2272656465656d536372697074223a22227d5d2c22507269764b657973223a6e756c6c2c22466c616773223a6e756c6c7d"
	//priKeyWif := "cSyz5HhtLGidariuHBVfqTotSFwu8fuXWVGqKWzNwqZvDzwtsc7i"
	//rawTx := "7b225261775478223a2230313030303030303031643965353231633739303831326434636362656335636233326234393334343139343339313362656635306534383234613362396637313465313766303531313030303030303030303066666666666630303032313032373030303030303030303030303139373661393134336233316165653135373332383833356566393539393861316333633137626565393866643464633838616363306134663530353030303030303030313937366139313437316437626335363566336462373866363164396335326435313764353261663361363264346435383861633030303030303030222c22496e70757473223a5b7b2274786964223a2231313035376665313134663762396133323434383065663562653133333939343431333434393262623335636563636234633264383139306337323165356439222c22766f7574223a302c227363726970745075624b6579223a223736613931343731643762633536356633646237386636316439633532643531376435326166336136326434643538386163222c2272656465656d536372697074223a22227d5d2c22507269764b657973223a6e756c6c2c22466c616773223a6e756c6c7d"
	//priKeyWif := "KycDKyx24avpMvqEDXaS2mkV8zCJF18LvgpsFcKgdccAbzwZg6dn"

	rawTx := "7b225261775478223a223031303030303030303138316431663336393165623331306235343830396362356532623430353963366162333566373931323363356663376332663062393465386431363636366164303130303030303030306666666666663030303330303030303030303030303030303030313636613134366636643665363930303030303030303830303030303033303030303030303030313331326430303232303230303030303030303030303031376139313463353739333432633263346339323230323035653263646332383536313730343063393234613061383763653561373634383137303030303030313937366139313464353831306663376536333231663430366363623263316634346561303036666665633333656332383861633030303030303030222c22496e70757473223a5b7b2274786964223a2261643636363664316538393430623266376366636335323339316637333561626336353934303262356563623039343862353130623331653639663364313831222c22766f7574223a312c227363726970745075624b6579223a223736613931346435383130666337653633323166343036636362326331663434656130303666666563333365633238386163222c2272656465656d536372697074223a22227d5d2c22507269764b657973223a6e756c6c2c22466c616773223a6e756c6c7d"
	priKeyWif := "cTRHJC6Ng17z36cjca8tZWpPGGTVW5qDWdvi2B8nhiQ8KskkM4NW"
	netType := "regtest"
	result, err := client.SignRawTx(rawTx, priKeyWif, netType)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	t.Logf("result:%v", result)
}
