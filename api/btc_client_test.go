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
	rawTx := "7b225261775478223a22303130303030303030316136613162333734623232633563353564313864396165656438356533613931333438646338326431303431383833643537646632666462346432366432313930323030303030303030666666666666303030333030303030303030303030303030303031363661313436663664366536393030303030303030383030303030303330303030303030303031333132643030323230323030303030303030303030303139373661393134316338653962626263663731653364363662343061636536633832303362373266393631353530373838616339613735353330323030303030303030313937366139313435383531306536316564323064656338393837333437326330303036383334303736643963663262383861633030303030303030222c22496e70757473223a5b7b2274786964223a2231396432323634646462326664663537336438383431313032646338386433343931336135656438656539613864643135353563326362323734623361316136222c22766f7574223a322c227363726970745075624b6579223a223736613931343538353130653631656432306465633839383733343732633030303638333430373664396366326238386163222c2272656465656d536372697074223a22227d5d2c22507269764b657973223a6e756c6c2c22466c616773223a6e756c6c7d"
	priKeyWif := "cSyz5HhtLGidariuHBVfqTotSFwu8fuXWVGqKWzNwqZvDzwtsc7i"
	netType := "regtest"
	result, err := client.SignRawTx(rawTx, priKeyWif, netType)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	t.Logf("result:%v", result)
}
