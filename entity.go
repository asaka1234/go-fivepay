package go_fivepay

type FivePayInitParams struct {
	MerchantId          string `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"`                                     // merchantId
	AccessKey           string `json:"accessKey" mapstructure:"accessKey" config:"accessKey"  yaml:"accessKey"`                                         // accessKey
	DepositUrlByCn      string `json:"depositUrlByCn" mapstructure:"depositUrlByCn" config:"depositUrlByCn"  yaml:"depositUrlByCn"`                     // 请求三方入金地址-中文
	DepositUrlByEn      string `json:"depositUrlByEn" mapstructure:"depositUrlByEn" config:"depositUrlByEn"  yaml:"depositUrlByEn"`                     // 请求三方入金地址-英文
	DepositUrlById      string `json:"depositUrlById" mapstructure:"depositUrlById" config:"depositUrlById"  yaml:"depositUrlById"`                     // 请求三方入金地址-印尼
	DepositUrlByVi      string `json:"depositUrlByVi" mapstructure:"depositUrlByVi" config:"depositUrlByVi"  yaml:"depositUrlByVi"`                     // 请求三方入金地址-越南
	DepositUrlByTh      string `json:"depositUrlByTh" mapstructure:"depositUrlByTh" config:"depositUrlByTh"  yaml:"depositUrlByTh"`                     // 请求三方入金地址-泰国
	DepositByF2fUrlByCn string `json:"depositByF2fUrlByCn" mapstructure:"depositByF2fUrlByCn" config:"depositByF2fUrlByCn"  yaml:"depositByF2fUrlByCn"` // 请求三方入金地址-中文
	DepositByF2fUrlByEn string `json:"depositByF2fUrlByEn" mapstructure:"depositByF2fUrlByEn" config:"depositByF2fUrlByEn"  yaml:"depositByF2fUrlByEn"` // 请求三方入金地址-英文
	DepositByF2fUrlById string `json:"depositByF2fUrlById" mapstructure:"depositByF2fUrlById" config:"depositByF2fUrlById"  yaml:"depositByF2fUrlById"` // 请求三方入金地址-印尼
	DepositByF2fUrlByVi string `json:"depositByF2fUrlByVi" mapstructure:"depositByF2fUrlByVi" config:"depositByF2fUrlByVi"  yaml:"depositByF2fUrlByVi"` // 请求三方入金地址-越南
	DepositByF2fUrlByTh string `json:"depositByF2fUrlByTh" mapstructure:"depositByF2fUrlByTh" config:"depositByF2fUrlByTh"  yaml:"depositByF2fUrlByTh"` // 请求三方入金地址-泰国
	WithdrawUrlByCn     string `json:"withdrawUrlByCn" mapstructure:"withdrawUrlByCn" config:"withdrawUrlByCn"  yaml:"withdrawUrlByCn"`
	WithdrawUrlByEn     string `json:"withdrawUrlByEn" mapstructure:"withdrawUrlByEn" config:"withdrawUrlByEn"  yaml:"withdrawUrlByEn"`
	WithdrawUrlById     string `json:"withdrawUrlById" mapstructure:"withdrawUrlById" config:"withdrawUrlById"  yaml:"withdrawUrlById"`
	WithdrawUrlByVi     string `json:"withdrawUrlByVi" mapstructure:"withdrawUrlByVi" config:"withdrawUrlByVi"  yaml:"withdrawUrlByVi"`
	WithdrawUrlByTh     string `json:"withdrawUrlByTh" mapstructure:"withdrawUrlByTh" config:"withdrawUrlByTh"  yaml:"withdrawUrlByTh"`
	NotifyUrlByDeposit  string `json:"notifyUrlByDeposit" mapstructure:"notifyUrlByDeposit" config:"notifyUrlByDeposit"  yaml:"notifyUrlByDeposit"`     //入金回调通知地址
	NotifyUrlByWithdraw string `json:"notifyUrlByWithdraw" mapstructure:"notifyUrlByWithdraw" config:"notifyUrlByWithdraw"  yaml:"notifyUrlByWithdraw"` //出金回调通知地址
	ReturnUrlByDeposit  string `json:"returnUrlByDeposit" mapstructure:"returnUrlByDeposit" config:"returnUrlByDeposit"  yaml:"returnUrlByDeposit"`
	ReturnUrlByWithdraw string `json:"returnUrlByWithdraw" mapstructure:"returnUrlByWithdraw" config:"returnUrlByWithdraw"  yaml:"returnUrlByWithdraw"` //付款页重定向到该URL
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
	Sign            string `json:"sign" mapstructure:"sign"`
}

type FivePayPaymentBackRsp struct {
	OrderNo         string `json:"orderNo" mapstructure:"orderNo"`                 // 平台给商家的唯一ID
	MerchantOrderNo string `json:"merchantOrderNo" mapstructure:"merchantOrderNo"` // 商户订单号
	OrderAmount     string `json:"orderAmount" mapstructure:"orderAmount"`         // 订单金额
	Status          string `json:"status" mapstructure:"status"`                   // 订单状态 1 – New order 2 – Waiting for payment 3 – Member has paid 4 – The payment has been confirmed 6 – Expired 7 – Cancelled
}

// deposit - f2f
type FivePayDepositByF2fHandleReq struct {
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

type FivePayDepositByF2fHandleRsp struct {
	OrderNo         string `json:"orderNo" mapstructure:"orderNo"`                 //平台给商家的唯一ID
	CurrencyCode    string `json:"currencyCode" mapstructure:"currencyCode"`       //币种
	MerchantId      string `json:"merchantId" mapstructure:"merchantId"`           //商户号
	MemberId        string `json:"memberId" mapstructure:"memberId"`               //会员ID
	ChannelName     string `json:"channelName" mapstructure:"channelName"`         //OTC only
	OrderAmount     string `json:"orderAmount" mapstructure:"orderAmount"`         //订单金额
	MerchantOrderNo string `json:"merchantOrderNo" mapstructure:"merchantOrderNo"` //商家未分配给订单的唯一订单
	Status          string `json:"status" mapstructure:"status"`                   //1 – New order 2 – Waiting for payment 3 – Member has paid 4 – The payment has been confirmed 6 – Expired 7 – Cancelled
	Sign            string `json:"sign" mapstructure:"sign"`                       //签名
}

type FivePayDepositByF2fBackReq struct {
	OrderNo         string `json:"orderNo" mapstructure:"orderNo"`                 //平台给商家的唯一ID
	CurrencyCode    string `json:"currencyCode" mapstructure:"currencyCode"`       //币种
	MerchantId      int    `json:"merchantId" mapstructure:"merchantId"`           //商户号
	MemberId        string `json:"memberId" mapstructure:"memberId"`               //会员ID
	ChannelName     string `json:"channelName" mapstructure:"channelName"`         //OTC only
	OrderAmount     string `json:"orderAmount" mapstructure:"orderAmount"`         //订单金额
	MerchantOrderNo string `json:"merchantOrderNo" mapstructure:"merchantOrderNo"` //商家未分配给订单的唯一订单
	Status          string `json:"status" mapstructure:"status"`                   //1 – New order 2 – Waiting for payment 3 – Member has paid 4 – The payment has been confirmed 6 – Expired 7 – Cancelled
	Sign            string `json:"sign" mapstructure:"sign"`
}

type FivePayDepositByF2fBackRsp struct {
	OrderNo         string `json:"orderNo" mapstructure:"orderNo"`                 // 平台给商家的唯一ID
	MerchantOrderNo string `json:"merchantOrderNo" mapstructure:"merchantOrderNo"` // 商户订单号
	OrderAmount     string `json:"orderAmount" mapstructure:"orderAmount"`         // 订单金额
	Status          string `json:"status" mapstructure:"status"`                   // 订单状态 1 – New order 2 – Waiting for payment 3 – Member has paid 4 – The payment has been confirmed 6 – Expired 7 – Cancelled
}

// withdraw
type FivePayWithdrawReq struct {
	MerchantId      int    `json:"merchantId" mapstructure:"merchantId"`           //平台给商家的唯一ID
	MerchantOrderNo string `json:"merchantOrderNo" mapstructure:"merchantOrderNo"` //商家未分配给订单的唯一订单
	Wallet          string `json:"wallet" mapstructure:"wallet"`                   // 钱包 OTC Buy – OTCBuy Wallet/Coin2Coin – Coin2Coin/Wallet Fiat 2 Fiat – Fiat2Fiat/Wallet Crypto Wallet – Crypto Wallet

	Token              string `json:"token" mapstructure:"token"`                           //
	WithdrawalAmount   string `json:"withdrawalAmount" mapstructure:"withdrawalAmount"`     //出金金额
	ByReceivableAmount bool   `json:"byReceivableAmount" mapstructure:"byReceivableAmount"` //按应收金额 （如果金额为Withdrawalamount，则填写false，如果金额为ReceivableAmount，则填写True）
	//WalletAddress            string `json:"alletAddress" mapstructure:"walletAddress"`                        //钱包地址
	BeneficiaryName          string `json:"beneficiaryName" mapstructure:"beneficiaryName"`                   //收款人姓名
	BeneficiaryAccountNumber string `json:"beneficiaryAccountNumber" mapstructure:"beneficiaryAccountNumber"` //收款人账号

	BeneficiaryBank string `json:"beneficiaryBank" mapstructure:"beneficiaryBank"` //受益人银行
	//BeneficiaryBankAddress   string `json:"BeneficiaryBankAddress" mapstructure:"BeneficiaryBankAddress"`     //收款人银行地址
	//BeneficiaryBankCode      string `json:"BeneficiaryBankCode" mapstructure:"BeneficiaryBankCode"`           //收款人银行代码
	//BeneficiaryBankSwiftCode string `json:"BeneficiaryBankSwiftCode" mapstructure:"BeneficiaryBankSwiftCode"` //受益人银行Swift代码
	//BeneficiaryEmail         string `json:"BeneficiaryEmail" mapstructure:"BeneficiaryEmail"`                 //Beneficiary电子邮件
	//BeneficiaryPhoneNumber   string `json:"BeneficiaryPhoneNumber" mapstructure:"BeneficiaryPhoneNumber"`     //受益人电话号码
	//IFSC                     string `json:"IFSC" mapstructure:"IFSC"`                                         //收款银行 IFSC 为法币2法币钱包提款，强制性
	//PaymentMethod            string `json:"PaymentMethod" mapstructure:"PaymentMethod"`                       //
	//TaxNumber                string `json:"TaxNumber" mapstructure:"TaxNumber"`                               //仅当付款方式为 EVP 时，BRL 提款的税号
	NotifyUrl string `json:"notifyUrl" mapstructure:"notifyUrl"` //通知回调地址
	//Sign                     string `json:"Sign" mapstructure:"Sign"`                                         //签名
}

type FivePayWithdrawDecodeRsp struct {
	MerchantId               float64 `json:"merchantId" mapstructure:"MerchantId"`                             //平台给商家的唯一ID
	MerchantOrderNo          string  `json:"merchantOrderNo" mapstructure:"MerchantOrderNo"`                   //商家未分配给订单的唯一订单
	Wallet                   string  `json:"wallet" mapstructure:"Wallet"`                                     // 钱包 OTC Buy – OTCBuy Wallet/Coin2Coin – Coin2Coin/Wallet Fiat 2 Fiat – Fiat2Fiat/Wallet Crypto Wallet – Crypto Wallet
	Token                    string  `json:"token" mapstructure:"Token"`                                       //
	WithdrawalAmount         string  `json:"withdrawalAmount" mapstructure:"WithdrawalAmount"`                 //出金金额
	ByReceivableAmount       string  `json:"byReceivableAmount" mapstructure:"ByReceivableAmount"`             //按应收金额 （如果金额为Withdrawalamount，则填写false，如果金额为ReceivableAmount，则填写True）
	BeneficiaryName          string  `json:"beneficiaryName" mapstructure:"BeneficiaryName"`                   //收款人姓名
	BeneficiaryAccountNumber string  `json:"beneficiaryAccountNumber" mapstructure:"BeneficiaryAccountNumber"` //收款人账号
	BeneficiaryBank          string  `json:"beneficiaryBank" mapstructure:"BeneficiaryBank"`                   //受益人银行
	NotifyUrl                string  `json:"notifyUrl" mapstructure:"notifyUrl"`                               //通知回调地址
	Sign                     string  `json:"Sign" mapstructure:"Sign"`                                         //签名
}

type FivePayWithdrawSubDataRsp struct {
	Status  bool   `json:"status" mapstructure:"status"` //true false
	Message string `json:"message" mapstructure:"message"`
	Data    string `json:"data" mapstructure:"data"` //请求时传的参数
}

type FivePayWithdrawRsp struct {
	Success bool                        `json:"success" mapstructure:"success"` //true false
	Data    []FivePayWithdrawSubDataRsp `json:"data" mapstructure:"data"`       //请求时传的参数
}

type FivePayWithdrawBackReq struct {
	WithdrawalId      string `json:"withdrawalId" mapstructure:"withdrawalId"`           //提款ID 平台给商家的唯一ID
	MerchantOrderNo   string `json:"merchantOrderNo" mapstructure:"merchantOrderNo"`     //商家未分配给订单的唯一订单
	WithdrawalAmount  string `json:"withdrawalAmount" mapstructure:"withdrawalAmount"`   //提款金额
	WithdrawalCharges string `json:"withdrawalCharges" mapstructure:"withdrawalCharges"` //提款费用
	Status            string `json:"status" mapstructure:"status"`                       //订单状态1 – 待确认2 – 进行中3 – 已批准4 – 已拒绝
	//RejectedReason    string `json:"rejectedReason" mapstructure:"rejectedReason"`       //交易明细
	Sign string `json:"sign" mapstructure:"sign"` //签名
	//TransactionHash   string `json:"transactionHash" mapstructure:"transactionHash"`     //交易hash
}

type FivePayWithdrawBackRsp struct {
	MerchantOrderNo   string `json:"merchantOrderNo" mapstructure:"merchantOrderNo"`     // 商户订单号
	WithdrawalId      string `json:"withdrawalId" mapstructure:"withdrawalId"`           // 出金订单ID
	WithdrawalAmount  string `json:"withdrawalAmount" mapstructure:"withdrawalAmount"`   // 出金订单金额
	WithdrawalCharges string `json:"withdrawalCharges" mapstructure:"withdrawalCharges"` //
	Status            string `json:"status" mapstructure:"status"`                       // 订单状态1 – 待确认2 – 进行中3 – 已批准4 – 已拒绝
}
