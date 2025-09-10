package go_fivepay

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/asaka1234/go-fivepay/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"log"
)

// 集成接口
func (cli *Client) Withdraw(req FivePayWithdrawReq) (*FivePayWithdrawDecodeRsp, error) {

	var param map[string]interface{}
	mapstructure.Decode(req, &param)

	//发送请求的psp地址
	rawURL := ""
	if req.Token == "VND" {
		rawURL = cli.Params.WithdrawUrlByVi
	} else if req.Token == "IDR" {
		rawURL = cli.Params.WithdrawUrlById
	} else if req.Token == "THB" {
		rawURL = cli.Params.WithdrawUrlByTh
	} else if req.Token == "CNY" {
		rawURL = cli.Params.WithdrawUrlByCn
	} else {
		rawURL = cli.Params.WithdrawUrlByEn
	}

	//补充字段
	param["merchantId"] = cast.ToString(cli.Params.MerchantId)
	param["notifyUrl"] = cast.ToString(cli.Params.NotifyUrlByWithdraw)
	log.Printf("Withdraw param: %+v", param)

	// 1. 加密所有需要加密的参数
	paramEncrypt, err := utils.EncryptAllByWithdraw(param, cli.Params.AccessKey)
	if err != nil {
		log.Fatalf("Withdraw Error encrypting parameters: %v", err)
		return nil, err
	}
	fmt.Println("Withdraw Encrypted Params (before sign):", paramEncrypt)

	// 2. 生成签名
	signature := utils.WithdrawSign(paramEncrypt) // 签名是基于加密后的参数
	paramEncrypt["sign"] = signature
	fmt.Println("Withdraw Final Params (with sign):", paramEncrypt)

	//----------------------
	var result *FivePayWithdrawRsp

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(paramEncrypt).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp))
	cli.logger.Infof("PSPResty#fivepay#withdraw->%s", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode())
	}

	if resp.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp.Error(), resp.Body())
	}

	if result.Success == false || result.Data == nil {
		return nil, fmt.Errorf("result status code: %d", resp.StatusCode())
	}

	if result.Data[0].Status == false {
		return nil, fmt.Errorf("%s", result.Data[0].Message)
	}

	var intermediate map[string]interface{}
	if err = json.Unmarshal([]byte(result.Data[0].Data), &intermediate); err != nil {
		return nil, fmt.Errorf("result unmarshal err: %+v", err.Error())
	}

	var rsp FivePayWithdrawDecodeRsp
	if err = mapstructure.Decode(intermediate, &rsp); err != nil {
		return nil, fmt.Errorf("result encode err: %+v", err.Error())
	}

	return &rsp, nil
}
