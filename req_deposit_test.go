package go_fivepay

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestDeposit(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &FivePayInitParams{MERCHANT_ID, ACCESS_KEY, DEPOSIT_URL_BY_EN, WITHDRAW_URL_BY_EN, NOTIFY_URL, RETURN_URL})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenDepositRequestDemo() FivePayPaymentHandleReq {
	return FivePayPaymentHandleReq{
		CurrencyCode:    "VND", // 暂时支持MYR，VND，THB，HKD，IDR
		OrderAmount:     "100000",
		Name:            "jane",
		Email:           "jane.y@logtec.com",
		MemberId:        "2335",
		MerchantOrderNo: "3324526277", //商户订单号
	}
}
