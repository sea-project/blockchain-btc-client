package api

import "testing"

func TestBTCClient_GetBlockCount(t *testing.T) {
	client := NewBTCClient("http://8.210.178.221:8332/", "admin", "Q3Az6XHNdE", 0)
	num, err := client.GetBlockCount()
	if err != nil {
		t.Errorf("num:%v", num)
		return
	}
	t.Logf("num:%v", num)
}
