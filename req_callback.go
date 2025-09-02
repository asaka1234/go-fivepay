package go_fivepay

import (
	"log"
)

// https://developer.paysafe.com/en/neteller-api-1/#/#webhooks-events
// http://paysafegroup.github.io/neteller_rest_api_v1/#/introduction/technical-introduction/webhooks
func (cli *Client) PaymentCallback(req FivePayPaymentBackReq, processor func(FivePayPaymentBackReq) error) error {
	log.Printf("FivePay#back#req: %+v", req)

	//开始处理
	return processor(req)
}
