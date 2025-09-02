package go_fivepay

type FivePayInitParams struct {
	MerchantId string `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"` // merchantId
	AccessKey  string `json:"accessKey" mapstructure:"accessKey" config:"accessKey"  yaml:"accessKey"`     // accessKey
	//DepositUrlByCn string `json:"depositUrlByCn" mapstructure:"depositUrlByCn" config:"depositUrlByCn"  yaml:"depositUrlByCn"` // 请求三方入金地址-中文
	DepositUrlByEn string `json:"depositUrlByEn" mapstructure:"depositUrlByEn" config:"depositUrlByEn"  yaml:"depositUrlByEn"` // 请求三方入金地址-英文
	//DepositUrlById  string `json:"depositUrlById" mapstructure:"depositUrlById" config:"depositUrlById"  yaml:"depositUrlById"`     // 请求三方入金地址-印尼
	//DepositUrlByVi  string `json:"depositUrlByVi" mapstructure:"depositUrlByVi" config:"depositUrlByVi"  yaml:"depositUrlByVi"`     // 请求三方入金地址-越南
	//DepositUrlByTh  string `json:"depositUrlByTh" mapstructure:"depositUrlByTh" config:"depositUrlByTh"  yaml:"depositUrlByTh"`     // 请求三方入金地址-泰国
	//WithdrawUrlByCn string `json:"withdrawUrlByCn" mapstructure:"withdrawUrlByCn" config:"withdrawUrlByCn"  yaml:"withdrawUrlByCn"` // 请求三方出金地址-中文
	WithdrawUrlByEn string `json:"withdrawUrlByEn" mapstructure:"withdrawUrlByEn" config:"withdrawUrlByEn"  yaml:"withdrawUrlByEn"` // 请求三方出金地址-英文
	//WithdrawUrlById string `json:"withdrawUrlById" mapstructure:"withdrawUrlById" config:"withdrawUrlById"  yaml:"withdrawUrlById"` // 请求三方出金地址-印尼
	//WithdrawUrlByVi string `json:"withdrawUrlByVi" mapstructure:"withdrawUrlByVi" config:"withdrawUrlByVi"  yaml:"withdrawUrlByVi"` // 请求三方出金地址-越南
	//WithdrawUrlByTh string `json:"withdrawUrlByTh" mapstructure:"withdrawUrlByTh" config:"withdrawUrlByTh"  yaml:"withdrawUrlByTh"` // 请求三方出金地址-泰国
	NotifyUrl string `json:"notifyUrl" mapstructure:"notifyUrl" config:"notifyUrl"  yaml:"notifyUrl"` //回调通知地址
	ReturnUrl string `json:"returnUrl" mapstructure:"returnUrl" config:"returnUrl"  yaml:"returnUrl"` //付款页重定向到该URL
}

// pay
type FivePayPaymentHandleReq struct {
	MemberId        string `json:"memberId" mapstructure:"memberId"`               // CRM Member Id
	Email           string `json:"email" mapstructure:"email"`                     // CRM Member Email
	Name            string `json:"name" mapstructure:"name"`                       // CRM Member Name, this name must be the bank account holder name which used to make transaction. MANDATORY for THB
	OrderAmount     string `json:"orderAmount" mapstructure:"orderAmount"`         // The order payment amount, IDR and INR order amount is not allowed decimal places.
	MerchantOrderNo string `json:"merchantOrderNo" mapstructure:"merchantOrderNo"` //唯一订单号
	NotifyUrl       string `json:"notifyUrl" mapstructure:"notifyUrl"`             //回调通知地址
	ReturnUrl       string `json:"returnUrl" mapstructure:"returnUrl"`             //付款页重定向到该URL
	CurrencyCode    string `json:"currencyCode" mapstructure:"currencyCode"`       //币种
	//sdk
	//MerchantId  int `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"`     // merchantId
}

type DepositRspData struct {
	OrderId         int `json:"orderId" mapstructure:"orderId"`
	MerchantOrderNo int `json:"merchantOrderNo" mapstructure:"merchantOrderNo"`
	ChannelName     int `json:"channelName" mapstructure:"channelName"`
	Details         int `json:"details" mapstructure:"details"`
}

type FivePayPaymentHandleRsp struct {
	HTMLString string `json:"HTMLString" mapstructure:"HTMLString"`
}

// callback
type FivePayPaymentBackReq struct {
	OrderNo         string `json:"orderNo" mapstructure:"orderNo"`                 // 平台给商家的唯一ID
	CurrencyCode    string `json:"currencyCode" mapstructure:"currencyCode"`       // 币种
	MerchantId      int    `json:"merchantId" mapstructure:"merchantId"`           // 商户ID
	MemberId        string `json:"memberId" mapstructure:"memberId"`               // 会员ID
	ChannelName     string `json:"channelName" mapstructure:"channelName"`         // OTC only
	OrderAmount     string `json:"orderAmount" mapstructure:"orderAmount"`         // 订单金额
	MerchantOrderNo string `json:"merchantOrderNo" mapstructure:"merchantOrderNo"` // 商户订单ID
	Status          string `json:"status" mapstructure:"status"`                   // 订单状态 1 – New order 2 – Waiting for payment 3 – Member has paid 4 – The payment has been confirmed 6 – Expired 7 – Cancelled
	Sign            string `json:"sign" mapstructure:"sign"`                       // 签名
}

// withdraw
type FivePayWithdrawHandleReq struct {
	MemberId        string `json:"memberId" mapstructure:"memberId"`               // CRM Member Id
	Email           string `json:"email" mapstructure:"email"`                     // CRM Member Email
	Name            string `json:"name" mapstructure:"name"`                       // CRM Member Name, this name must be the bank account holder name which used to make transaction. MANDATORY for THB
	OrderAmount     string `json:"orderAmount" mapstructure:"orderAmount"`         // The order payment amount, IDR and INR order amount is not allowed decimal places.
	MerchantOrderNo string `json:"merchantOrderNo" mapstructure:"merchantOrderNo"` //唯一订单号
	NotifyUrl       string `json:"notifyUrl" mapstructure:"notifyUrl"`             //回调通知地址
	ReturnUrl       string `json:"returnUrl" mapstructure:"returnUrl"`             //付款页重定向到该URL
	CurrencyCode    string `json:"currencyCode" mapstructure:"currencyCode"`       //币种
	//sdk
	//MerchantId  int `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"`     // merchantId
}

type FivePayWithdrawHandleRsp struct {
	HTMLString string `json:"HTMLString" mapstructure:"HTMLString"`
}
