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

	//var data FivePayWithdrawReq
	//if err = json.Unmarshal([]byte(resp.Data), &data); err != nil {
	//	fmt.Println("json.Unmarshal error:", err)
	//	return
	//}
	//
	//fmt.Printf("success: %v msg: %s data: %+v\n", resp.Success, resp.Message, data)
}

func GenWithdrawRequestDemo() FivePayWithdrawReq {
	return FivePayWithdrawReq{
		MerchantOrderNo:    "19631262871196835841",
		CurrencyCode:       "VND",
		Wallet:             "Coin2Coin", // 钱包 OTC Buy – OTCBuy Wallet/Coin2Coin – Coin2Coin/Wallet Fiat 2 Fiat – Fiat2Fiat/Wallet Crypto Wallet – Crypto Wallet
		Token:              "BTC",
		WithdrawalAmount:   "605935",
		ByReceivableAmount: true,
		WalletAddress:      "sdRG58sgdee",
		//BeneficiaryName:          "jane",   //收款人姓名
		//BeneficiaryAccountNumber: "605935", //收款人账号
	}
}
