package domainser

import "encoding/xml"

// 通联请求头
type TLReqHeader struct {
	TRX_CODE   string // 交易代码
	LEVEL      string // 处理级别（0-9  0优先级最低，默认为5）
	REQ_SN     string // 交易流水号（必须全局唯一）
	SIGNED_MSG string // 签名信息
}

// 通联请求头部信息
type TLReqINFO struct {
	TRX_CODE    string `xml:"TRX_CODE,omitempty"`    // 交易代码
	VERSION     string `xml:"VERSION,omitempty"`     // 版本（04）
	DATA_TYPE   string `xml:"DATA_TYPE,omitempty"`   // 数据格式（2：xml格式）
	LEVEL       string `xml:"LEVEL,omitempty"`       // 处理级别（0-9  0优先级最低，默认为5）
	MERCHANT_ID string `xml:"MERCHANT_ID,omitempty"` // 商户代码
	USER_NAME   string `xml:"USER_NAME,omitempty"`   // 用户名
	USER_PASS   string `xml:"USER_PASS,omitempty"`   // 用户密码
	REQ_SN      string `xml:"REQ_SN,omitempty"`      // 交易流水号（必须全局唯一）
	SIGNED_MSG  string `xml:"SIGNED_MSG,omitempty"`  // 签名信息
}

// 通联响应头部信息
type TLRespINFO struct {
	TRX_CODE   string `xml:"TRX_CODE,omitempty"`   // 交易代码
	VERSION    string `xml:"VERSION,omitempty"`    // 版本（04）
	DATA_TYPE  string `xml:"DATA_TYPE,omitempty"`  // 数据格式（2：xml格式）
	REQ_SN     string `xml:"REQ_SN,omitempty"`     // 交易流水号(原请求报文的流水号，原样返回)
	RET_CODE   string `xml:"RET_CODE,omitempty"`   // 返回代码
	ERR_MSG    string `xml:"ERR_MSG,omitempty"`    // 错误信息
	SIGNED_MSG string `xml:"SIGNED_MSG,omitempty"` // 签名信息
}

// 通联响应大多数正文
type TLRespContent struct {
	RET_CODE string `xml:"RET_CODE,omitempty"` // 返回代码
	ERR_MSG  string `xml:"ERR_MSG,omitempty"`  // 错误文本
}

/******************************* 快捷支付签约短信 ****************************************/

type QuickPaySignSmsReq struct {
	XMLName xml.Name               `xml:"AIPG,omitempty"`
	INFO    TLReqINFO              `xml:"INFO,omitempty"`
	FAGRA   QuickPaySignSmsContent `xml:"FAGRA,omitempty"`
}

type QuickPaySignSmsContent struct {
	MERCHANT_ID  string `xml:"MERCHANT_ID,omitempty"`  // 商户代码
	BankCode     string `xml:"BANK_CODE,omitempty"`    // 银行code
	ACCOUNT_TYPE string `xml:"ACCOUNT_TYPE,omitempty"` //账户类型
	ACCOUNT_NO   string `xml:"ACCOUNT_NO,omitempty"`   // 账号（借记卡或信用卡）
	ACCOUNT_NAME string `xml:"ACCOUNT_NAME,omitempty"` // 账号名（借记卡或信用卡上的所有人姓名）
	ACCOUNT_PROP string `xml:"ACCOUNT_PROP,omitempty"` //账户属性
	ID_TYPE      string `xml:"ID_TYPE,omitempty"`      // 开户证件类型（0身份证，1户口簿，2护照，3军官证，4士兵证...）
	ID           string `xml:"ID,omitempty"`           // 证件号
	TEL          string `xml:"TEL,omitempty"`          // 预留手机号
	MERREM       string `xml:"MERREM,omitempty"`       // 商户预留
	REMARK       string `xml:"REMARK,omitempty"`       // 备注
	AisleType    string `xml:"-"`
}

type QuickPaySmsResp struct {
	XMLName xml.Name      `xml:"AIPG,omitempty"`
	INFO    TLRespINFO    `xml:"INFO,omitempty"`
	Content TLRespContent `xml:"FAGRARET,omitempty"`
}

/***********************************************************************/

/******************************* 快捷支付签约确认 ****************************************/

type SmsCfrmReq struct {
	XMLName xml.Name       `xml:"AIPG,omitempty"`
	INFO    TLReqINFO      `xml:"INFO,omitempty"`
	FAGRC   SmsCfrmContent `xml:"FAGRC,omitempty"`
}

type SmsCfrmContent struct {
	MERCHANT_ID string `xml:"MERCHANT_ID,omitempty"` // 商户代码
	VERCODE     string `xml:"VERCODE,omitempty"`     // 验证码
	SRCREQSN    string `xml:"SRCREQSN,omitempty"`    //原交易请求号
}

type SmsCfrmResp struct {
	XMLName xml.Name           `xml:"AIPG,omitempty"`
	INFO    TLRespINFO         `xml:"INFO,omitempty"`
	Content SmsCfrmRespContent `xml:"FAGRCRET,omitempty"`
}

type SmsCfrmRespContent struct {
	TLRespContent
	AGRMNO string `xml:"AGRMNO,omitempty"`
}

/***********************************************************************/

/******************************* 快捷代扣 ****************************************/
// 快捷支付交易请求
type QuickTradeReq struct {
	XMLName xml.Name             `xml:"AIPG,omitempty"`
	INFO    TLReqINFO            `xml:"INFO,omitempty"`
	FASTTRX QuickTradeReqFASTTRX `xml:"FASTTRX,omitempty"`
}

//通联快捷代扣请求
type QuickTradeReqFASTTRX struct {
	BUSINESS_CODE string `xml:"BUSINESS_CODE,omitempty"` // 业务代码
	MERCHANT_ID   string `xml:"MERCHANT_ID,omitempty"`   // 商户代码
	SUBMIT_TIME   string `xml:"SUBMIT_TIME,omitempty"`   // 提交时间（YYYYMMDDHHMMSS）
	AGRMNO        string `xml:"AGRMNO,omitempty"`        // 协议号（签约时返回的协议号）
	ACCOUNT_NO    string `xml:"ACCOUNT_NO,omitempty"`    // 账号（借记卡或信用卡）
	ACCOUNT_NAME  string `xml:"ACCOUNT_NAME,omitempty"`  // 账号名（借记卡或信用卡上的所有人姓名）
	AMOUNT        string `xml:"AMOUNT,omitempty"`        // 金额(整数，单位分)
	CURRENCY      string `xml:"CURRENCY,omitempty"`      // 货币类型(人民币：CNY, 港元：HKD，美元：USD。不填时，默认为人民币)
	ID_TYPE       string `xml:"ID_TYPE,omitempty"`       // 开户证件类型（0身份证，1户口簿，2护照，3军官证，4士兵证...）
	ID            string `xml:"ID,omitempty"`            // 证件号
	TEL           string `xml:"TEL,omitempty"`           // 手机号
	CVV2          string `xml:"CVV2,omitempty"`          // CVV2（信用卡时必填）
	VAILDDATE     string `xml:"VAILDDATE,omitempty"`     // 有效期（信用卡时必填，格式MMYY（信用卡上的两位月两位年））
	CUST_USERID   string `xml:"CUST_USERID,omitempty"`   // 自定义用户号（商户自定义的用户号，开发人员可当作备注字段使用）
	SUMMARY       string `xml:"SUMMARY,omitempty"`       // 交易附言（填入网银的交易备注）
	REMARK        string `xml:"REMARK,omitempty"`        // 备注（供商户填入参考信息）
}

// 协议支付 响应
type QuickTradeResp struct {
	XMLName    xml.Name             `xml:"AIPG,omitempty"`
	INFO       TLRespINFO           `xml:"INFO,omitempty"`
	FASTTRXRET QuickTradeFASTTRXRET `xml:"FASTTRXRET,omitempty"`
}

type QuickTradeFASTTRXRET struct {
	RET_CODE    string `xml:"RET_CODE,omitempty"`    // 返回代码
	SETTLE_DAY  string `xml:"SETTLE_DAY,omitempty"`  // 完成日期（YYYYMMDD）
	ERR_MSG     string `xml:"ERR_MSG,omitempty"`     // 错误文本
	ACCT_SUFFIX string `xml:"ACCT_SUFFIX,omitempty"` // 卡号后4位
}

/***********************************************************************/
