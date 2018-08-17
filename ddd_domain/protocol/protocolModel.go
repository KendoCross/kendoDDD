package protocol

type SignModel struct {
	//通道号
	AisleType string `json:"AisleType"`
	//银行code，各平台不一样
	BankCode string `json:"BankCode"`
	//账户类型
	AccountType string `json:"AccountType"`
	//账户属性
	AccountProp string `json:"AccountProp"`
	//银行卡号
	BankCardNo string `json:"BankCardNo"`
	//预留手机号
	ReservePhone string `json:"Tel"`
	//银行卡预留的证件类型
	IDCardType string `json:"IDType"`
	//银行卡开户姓名
	CardName string `json:"CardName"`
	//银行卡预留的证件号码
	IDCardNo string `json:"IDCardNo"`
	//提示标识
	Merrem string `json:"Merrem"`
	//备注
	Remark string `json:"Remark"`
}
