package Factory_Method

import "testing"

func TestPay(t *testing.T) {
	aPay := APayReq{}
	if aPay.Pay() != "APay支付成功" {
		t.Fatal("aPay error")
	}

	bPay := BPayReq{}
	if bPay.Pay() != "BPay支付成功" {
		t.Fatal("bPay error")
	}
}
