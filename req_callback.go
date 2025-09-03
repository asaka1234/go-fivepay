package go_fivepay

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asaka1234/go-fivepay/utils"
	"github.com/mitchellh/mapstructure"
	"log"
)

func (cli *Client) PaymentCallback(req FivePayPaymentBackReq, processor func(FivePayPaymentBackReq) error) error {
	log.Printf("FivePay#back#req: %+v", req)
	//验证签名
	var params map[string]string
	mapstructure.Decode(req, &params)

	// 1. 解密所有需要解密的参数
	paramDecrypt, err := utils.DecryptAll(params, cli.Params.AccessKey)
	if err != nil {
		log.Fatalf("Error decrypting parameters: %v", err)
	}
	fmt.Println("Decrypted Params (before sign):", paramDecrypt)

	// 2. 生成签名
	signature := utils.DepositSign(paramDecrypt) // 签名是基于加密后的参数
	paramDecrypt["sign"] = signature
	fmt.Println("Final Params (with sign):", paramDecrypt)

	// Verify signature
	flag, err := utils.DepositBackVerify(params, cli.Params.AccessKey)
	if err != nil {
		log.Printf("Signature verification error: %v", err)
		return err
	}
	if !flag {
		//签名校验失败
		reqJson, _ := json.Marshal(req)
		log.Printf("H2Pay back verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}
