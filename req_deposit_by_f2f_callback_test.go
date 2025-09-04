package go_fivepay

import (
	"fmt"
	"testing"
)

func TestDepositByF2fCallback(t *testing.T) {
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

	req := FivePayDepositByF2fBackReq{
		OrderNo:         "fa92dfaae91c981f",
		MerchantId:      2,
		MemberId:        "c5a707fc7552c8408437b310a4b5d8a2",
		ChannelName:     "8c6b875886813bb750dbafc47cdd73953eec0cfc5928b725",
		OrderAmount:     "ed8878973d8ac89f",
		MerchantOrderNo: "9635a171a5e457e2c73514fb1eba17614a5f9df20b5b0072",
		CurrencyCode:    "4b798adf04415171",
		Status:          "55c429d4262a42c9",
		Sign:            "dd1bd407e4268a0f669a2497817bef4b",
	}

	//发请求
	err := cli.DepositByF2fCallBack(req, DepositByF2fBackProcessor)
	if err != nil {
		fmt.Println("fail")
		return
	}

	fmt.Println("SUCCESS")
	return
}

func DepositByF2fBackProcessor(rsp FivePayDepositByF2fBackRsp) error {
	fmt.Printf("deposit by f2f callback rsp: %v\n", rsp)
	return nil
}
