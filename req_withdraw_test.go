package go_fivepay

import (
	"fmt"
	"testing"
)

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &FivePayInitParams{
		MERCHANT_ID,
		ACCESS_KEY,
		DEPOSIT_URL_BY_EN,
		WITHDRAW_URL_BY_EN,
		NOTIFY_URL_BY_WITHDRAW,
		NOTIFY_URL_BY_WITHDRAW,
		RETURN_URL,
		RETURN_URL,
	})

	//发请求
	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}

	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() FivePayWithdrawReq {
	return FivePayWithdrawReq{
		MerchantOrderNo:          "19631262871196835843",
		Wallet:                   "Fiat2Fiat", // 钱包 OTC Buy – OTCBuy Wallet/Coin2Coin – Coin2Coin/Wallet Fiat 2 Fiat – Fiat2Fiat/Wallet Crypto Wallet – Crypto Wallet
		Token:                    "VND",
		WithdrawalAmount:         "605935",
		ByReceivableAmount:       true,
		BeneficiaryName:          "jane",   //收款人姓名
		BeneficiaryAccountNumber: "605935", //收款人账号
		BeneficiaryBank:          "asga",
	}
}
