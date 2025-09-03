package go_fivepay

import (
	"fmt"
	"testing"
)

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &FivePayInitParams{MERCHANT_ID, ACCESS_KEY, DEPOSIT_URL_BY_EN, WITHDRAW_URL_BY_EN, NOTIFY_URL_BY_DEPOSIT, NOTIFY_URL_BY_WITHDRAW, RETURN_URL})

	//发请求
	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}

	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() FivePayWithdrawHandleReq {
	return FivePayWithdrawHandleReq{
		CurrencyCode:    "VND", // 暂时支持MYR，VND，THB，HKD，IDR
		OrderAmount:     "100000",
		Name:            "jane",
		Email:           "jane.y@logtec.com",
		MemberId:        "2335",
		MerchantOrderNo: "21441113", //商户订单号
	}
}
