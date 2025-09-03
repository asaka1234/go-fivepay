package go_fivepay

import (
	"fmt"
	"github.com/asaka1234/go-fivepay/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"log"
)

// 集成接口
func (cli *Client) Deposit(req FivePayPaymentHandleReq) (map[string]interface{}, error) {
	//rawURL := cli.Params.DepositUrlByEn

	var param map[string]interface{}
	mapstructure.Decode(req, &param)

	//补充字段
	param["merchantId"] = cast.ToInt(cli.Params.MerchantId)
	param["returnUrl"] = cast.ToString(cli.Params.ReturnUrl)
	param["notifyUrl"] = cast.ToString(cli.Params.NotifyUrlByDeposit)
	log.Printf("param: %+v", param)

	// 1. 加密所有需要加密的参数
	paramEncrypt, err := utils.EncryptAll(param, cli.Params.AccessKey)
	if err != nil {
		log.Fatalf("Error encrypting parameters: %v", err)
		return nil, err
	}
	fmt.Println("Encrypted Params (before sign):", paramEncrypt)

	// 2. 生成签名
	signature := utils.DepositSign(paramEncrypt) // 签名是基于加密后的参数
	paramEncrypt["sign"] = signature
	fmt.Println("Final Params (with sign):", paramEncrypt)

	//发送请求的psp地址
	if req.CurrencyCode == "VND" {
		paramEncrypt["url"] = cli.Params.DepositUrlByVi
	} else if req.CurrencyCode == "IDR" {
		paramEncrypt["url"] = cli.Params.DepositUrlById
	} else if req.CurrencyCode == "THB" {
		paramEncrypt["url"] = cli.Params.DepositUrlByTh
	} else if req.CurrencyCode == "CNY" {
		paramEncrypt["url"] = cli.Params.DepositUrlByCn
	} else {
		paramEncrypt["url"] = cli.Params.DepositUrlByEn
	}

	return paramEncrypt, nil
}
