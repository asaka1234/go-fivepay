package go_fivepay

import (
	"fmt"
	"github.com/asaka1234/go-fivepay/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"log"
)

// 集成接口
// TODO withdraw不需要等待回调
func (cli *Client) Withdraw(req FivePayWithdrawHandleReq) (map[string]interface{}, error) {
	//rawURL := cli.Params.WithdrawUrlByEn

	var param map[string]interface{}
	mapstructure.Decode(req, &param)

	//补充字段
	param["merchantId"] = cast.ToString(cli.Params.MerchantId)
	param["returnUrl"] = cast.ToString(cli.Params.ReturnUrl)
	param["notifyUrl"] = cast.ToString(cli.Params.NotifyUrl)

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
	//----------------------
	//var result FivePayWithdrawHandleRsp
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
	//cli.logger.Infof("PSPResty#fivepay#withdraw->%s", string(restLog))
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
	//return &result, nil
}
