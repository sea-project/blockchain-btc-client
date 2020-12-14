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
