package go_fivepay

import (
	"errors"
	"fmt"
	"github.com/asaka1234/go-fivepay/utils"
	"github.com/mitchellh/mapstructure"
	"log"
)

func (cli *Client) PaymentCallback(req FivePayPaymentBackReq, processor func(FivePayPaymentBackReq, map[string]interface{}) error) error {
	log.Printf("FivePay#deposit#back#req: %+v", req)
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	// 1. 验证签名
	signature := utils.DepositSign(params)
	if signature != req.Sign {
		return errors.New("sign verify error")
	}

	params["notifyUrl"] = cli.Params.NotifyUrlByDeposit
	params["returnUrl"] = cli.Params.ReturnUrlByDeposit

	// 2. 解密所有需要解密的参数
	paramDecrypt, err := utils.DecryptAll(params, cli.Params.AccessKey)
	if err != nil {
		log.Fatalf("Error decrypting parameters: %v", err)
	}
	fmt.Println("FivePay deposit callback decrypted Params :", paramDecrypt)

	// 3. 处理业务逻辑
	return processor(req, paramDecrypt)
}

func (cli *Client) WithdrawCallBack(req FivePayWithdrawBackReq, processor func(FivePayWithdrawBackReq, map[string]interface{}) error) error {
	log.Printf("FivePay#withdraw#back#req: %+v", req)
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	// 1. 验证签名
	signature := utils.DepositSign(params)
	if signature != req.Sign {
		return errors.New("sign verify error")
	}

	params["notifyUrl"] = cli.Params.NotifyUrlByWithdraw
	params["returnUrl"] = cli.Params.NotifyUrlByWithdraw

	// 2. 解密所有需要解密的参数
	paramDecrypt, err := utils.DecryptAll(params, cli.Params.AccessKey)
	if err != nil {
		log.Fatalf("Error decrypting parameters: %v", err)
	}
	fmt.Println("FivePay deposit callback decrypted Params :", paramDecrypt)

	// 3. 处理业务逻辑
	return processor(req, paramDecrypt)
}
