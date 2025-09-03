package go_fivepay

import (
	"fmt"
	"github.com/asaka1234/go-fivepay/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"log"
)

// 集成接口
func (cli *Client) Deposit(req FivePayPaymentHandleReq) (map[string]string, error) {
	//rawURL := cli.Params.DepositUrlByEn

	var param map[string]string
	mapstructure.Decode(req, &param)

	//补充字段
	param["merchantId"] = cast.ToString(cli.Params.MerchantId)
	param["returnUrl"] = cast.ToString(cli.Params.ReturnUrl)
	param["notifyUrl"] = cast.ToString(cli.Params.NotifyUrl)
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

	return paramEncrypt, nil

	////----------------------
	//var result FivePayPaymentHandleRsp
	//
	//resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
	//	SetCloseConnection(true).
	//	R().
	//	SetHeader("Content-Type", "application/x-www-form-urlencoded").
	//	SetFormData(paramEncrypt).
	//	SetDebug(cli.debugMode).
	//	SetResult(&result).
	//	SetError(&result).
	//	Post(rawURL)
	//
	//restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp))
	//cli.logger.Infof("PSPResty#fivepay#deposit->%s", string(restLog))
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//if resp.StatusCode() != 200 {
	//	return nil, fmt.Errorf("status code: %d", resp.StatusCode())
	//}
	//
	//if resp.Error() != nil {
	//	//反序列化错误会在此捕捉
	//	return nil, fmt.Errorf("%v, body:%s", resp.Error(), resp.Body())
	//}
	//
	//// Log response
	//responseStr := string(resp.Body())
	//
	//// Build response struct
	//rsp := &FivePayPaymentHandleRsp{
	//	HTMLString: responseStr,
	//}
	//
	//return rsp, nil
}
