package go_fivepay

import "strings"

type FivePayEventType string

const (
	MERCHANT_ID = "2"
	ACCESS_KEY  = "vCqPGHoPuaBwE1alf4PeNfdh"

	DEPOSIT_URL_BY_CN = "http://uat.cn-payment.my5pay.com/f2fOrder/createorder"
	DEPOSIT_URL_BY_EN = "http://uat.en-payment.my5pay.com/f2fOrder/createorder"
	DEPOSIT_URL_BY_ID = "http://uat.id-payment.my5pay.com/f2fOrder/createorder"
	DEPOSIT_URL_BY_VI = "http://uat.vi-payment.my5pay.com/f2fOrder/createorder"
	DEPOSIT_URL_BY_TH = "http://uat.th-payment.my5pay.com/f2fOrder/createorder"

	DEPOSIT_BYF2F_URL_BY_CN = ""
	DEPOSIT_BYF2F_URL_BY_EN = ""
	DEPOSIT_BYF2F_URL_BY_ID = ""
	DEPOSIT_BYF2F_URL_BY_VI = ""
	DEPOSIT_BYF2F_URL_BY_TH = ""

	WITHDRAW_URL_BY_CN = "http://uat.en-payment.my5pay.com/Withdrawal/SubmitWithdrawal"
	WITHDRAW_URL_BY_EN = "http://uat.en-payment.my5pay.com/Withdrawal/SubmitWithdrawal"
	WITHDRAW_URL_BY_ID = "http://uat.en-payment.my5pay.com/Withdrawal/SubmitWithdrawal"
	WITHDRAW_URL_BY_VI = "http://uat.en-payment.my5pay.com/Withdrawal/SubmitWithdrawal"
	WITHDRAW_URL_BY_TH = "http://uat.en-payment.my5pay.com/Withdrawal/SubmitWithdrawal"

	NOTIFY_URL_BY_DEPOSIT  = "https://api-test.logtec.dev/fapi/cpti/payment/psp/fivepay/deposit/back" //"https://api.cptmarkets.com/fapi/cpti/payment/psp/fivepay/deposit/back"
	NOTIFY_URL_BY_WITHDRAW = "https://api-test.logtec.dev/fapi/cpti/payment/psp/fivepay/withdraw/back"
	RETURN_URL             = "https://portal.cptmarkets.com/zh/depositReport"
)

// GetName returns the name of the event type (same as value in this case)
func (f FivePayEventType) GetName() string {
	return string(f)
}

// GetValue returns the value of the event type
func (f FivePayEventType) GetValue() string {
	return string(f)
}

// Eq checks if the event type equals the given value (case-insensitive)
func (f FivePayEventType) Eq(value string) bool {
	return strings.EqualFold(string(f), value)
}

// String implements the Stringer interface
func (f FivePayEventType) String() string {
	return string(f)
}
