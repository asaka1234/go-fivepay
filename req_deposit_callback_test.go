package go_fivepay

import (
	"fmt"
	"testing"
)

func TestDepositCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &FivePayInitParams{
		MERCHANT_ID,
		ACCESS_KEY,
		DEPOSIT_URL_BY_EN,
		WITHDRAW_URL_BY_EN,
		NOTIFY_URL_BY_DEPOSIT,
		NOTIFY_URL_BY_WITHDRAW,
		RETURN_URL,
		RETURN_URL,
	})

	req := FivePayPaymentBackReq{
		OrderNo:         "11bccf360c95638f",
		MerchantId:      2,
		MemberId:        "c5a707fc7552c8408437b310a4b5d8a2",
		ChannelName:     "8c6b875886813bb750dbafc47cdd73953eec0cfc5928b725",
		OrderAmount:     "6ee8dcd629652024",
		MerchantOrderNo: "4f3d213650c20767e8dba869ddd44aa6c1ebcb30d3bcc4dd",
		CurrencyCode:    "4b798adf04415171",
		Status:          "55c429d4262a42c9",
		Sign:            "03a5dc7e4f5751f672d8dc3b3893f553",
	}

	//发请求
	err := cli.PaymentCallback(req, DepositBackProcessor)
	if err != nil {
		fmt.Println("fail")
		return
	}

	fmt.Println("SUCCESS")
	return
}

func DepositBackProcessor(rsp FivePayPaymentBackRsp) error {
	fmt.Printf("deposit callback rsp: %v\n", rsp)
	return nil
}
