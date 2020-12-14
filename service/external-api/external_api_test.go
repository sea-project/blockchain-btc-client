package external_api

import "testing"

func TestGetBalance(t *testing.T) {
	balance, err := GetBalance("1NVjJxqir4jyTvxWUG14g9ketBrzvqzs7z", 6)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	t.Logf("balance:%v", balance)
}
