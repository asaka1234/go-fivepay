package go_fivepay

import (
	"fmt"
	"testing"
)

func TestWithdrawCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &FivePayInitParams{
		MERCHANT_ID,
		ACCESS_KEY,
		DEPOSIT_URL_BY_CN,
		DEPOSIT_URL_BY_EN,
		DEPOSIT_URL_BY_ID,
		DEPOSIT_URL_BY_VI,
		DEPOSIT_URL_BY_TH,
		DEPOSIT_BYF2F_URL_BY_CN,
		DEPOSIT_BYF2F_URL_BY_EN,
		DEPOSIT_BYF2F_URL_BY_ID,
		DEPOSIT_BYF2F_URL_BY_VI,
		DEPOSIT_BYF2F_URL_BY_TH,
		WITHDRAW_URL_BY_CN,
		WITHDRAW_URL_BY_EN,
		WITHDRAW_URL_BY_ID,
		WITHDRAW_URL_BY_VI,
		WITHDRAW_URL_BY_TH,
		NOTIFY_URL_BY_DEPOSIT,
		NOTIFY_URL_BY_WITHDRAW,
		RETURN_URL,
		RETURN_URL,
	})

	req := FivePayWithdrawBackReq{
		WithdrawalId:      "5bc6cf83ad58c7d6",
		MerchantOrderNo:   "9635a171a5e457e2c73514fb1eba1761f7c561322ee298d1",
		WithdrawalAmount:  "a45ccd9bede3ef9c0a849a1cf816b719",
		WithdrawalCharges: "376d0117ccdcf39cf951069861ea4d4f",
		Status:            "317337bb88110e20",
		Sign:              "127ce9a6b5742b06c483b5a1b9ea40f9",
	}

	//发请求
	err := cli.WithdrawCallback(req, WithdrawBackProcessor)
	if err != nil {
		fmt.Println("fail")
		return
	}

	fmt.Println("SUCCESS")
	return
}

func WithdrawBackProcessor(rsp FivePayWithdrawBackRsp) error {
	fmt.Printf("withdraw callback rsp: %v\n", rsp)
	return nil
}
