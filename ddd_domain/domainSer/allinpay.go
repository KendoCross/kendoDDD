package domainser

import (
	"fmt"
	"math/rand"
	"time"

	utility "../Utility"
)

var Allinpay *TLPay

type TLPay struct {
	TLQuickpayAdrr string
	TLUserDev      string
	MerchantIdDev  string
}

func init() {
	//allinpay = new(TLPay)
	Allinpay = &TLPay{utility.TLQuickpayAdrr, utility.TLUserDev, utility.MerchantIdDev}
}

func InitTLReqHeader(reqHead *TLReqHeader) *TLReqINFO {
	return &TLReqINFO{
		TRX_CODE:    reqHead.TRX_CODE,      // 交易代码
		VERSION:     "04",                  // 版本（04）
		DATA_TYPE:   "2",                   // 数据格式（2：xml格式）
		LEVEL:       reqHead.LEVEL,         // 处理级别（0-9  0优先级最低，默认为5）
		MERCHANT_ID: utility.MerchantIdDev, // 商户代码
		USER_NAME:   utility.TLUserDev,     // 用户名
		USER_PASS:   utility.TLPassWordDev, // 用户密码
		REQ_SN:      reqHead.REQ_SN,        // 交易流水号（必须全局唯一）
		SIGNED_MSG:  reqHead.SIGNED_MSG,    // 签名信息
	}
}

//通联快捷签约，触发短信验证码
func (tl *TLPay) SignGetCode(smsContent *QuickPaySignSmsContent) (resut ApiResp) {
	reqSN := fmt.Sprintf("XxxX_%4d%d", rand.Intn(10000), time.Now().Unix())
	// 请求头
	headerReq := &TLReqHeader{
		TRX_CODE:   "310001", // 交易代码
		LEVEL:      "6",      // 处理级别（0-9  0优先级最低，默认为5）
		REQ_SN:     reqSN,    // 交易流水号（必须全局唯一）
		SIGNED_MSG: "",       // 签名信息
	}
	smsContent.MERCHANT_ID = utility.MerchantIdDev
	payReq := &QuickPaySignSmsReq{
		INFO:  *InitTLReqHeader(headerReq),
		FAGRA: *smsContent,
	}

	// xml
	tlReqXML, err := utility.ToTLRequestXmlByte(payReq)
	if err != nil {
		resut.ErrMsg = fmt.Sprintln("转化为通联请求XML出错 : ", err)
		return
	}
	bodyByte, err := utility.PostAllinPayXml(tl.TLQuickpayAdrr, tlReqXML)
	if err != nil {
		resut.ErrMsg = fmt.Sprintln("通联服务请求异常 :  ", err)
		return
	}
	smsResp := &QuickPaySmsResp{}
	if err = utility.VerifyTLResponse(bodyByte, smsResp); err != nil {
		resut.ErrMsg = fmt.Sprintln("通联验签不通过：", err)
		return
	}

	//fmt.Printf("%#v\n", smsResp)

	if smsResp.INFO.RET_CODE != Succe_Code {
		resut.RetCode = smsResp.INFO.RET_CODE
		resut.ErrMsg = smsResp.INFO.ERR_MSG
		return
	}

	if smsResp.Content.RET_CODE != Succe_Code {
		resut.RetCode = smsResp.Content.RET_CODE
		resut.ErrMsg = smsResp.Content.ERR_MSG
		return
	}

	resut.IsSuccess = true
	resut.RetInfo = reqSN
	return
}

func (tl *TLPay) SignCfrm(smsCfrmContent *SmsCfrmContent) (resut ApiResp) {
	reqSN := fmt.Sprintf("XxxX_%4d%d", rand.Intn(10000), time.Now().Unix())
	// 请求头
	headerReq := &TLReqHeader{
		TRX_CODE:   "310002", // 交易代码
		LEVEL:      "6",      // 处理级别（0-9  0优先级最低，默认为5）
		REQ_SN:     reqSN,    // 交易流水号（必须全局唯一）
		SIGNED_MSG: "",       // 签名信息
	}
	smsCfrmContent.MERCHANT_ID = utility.MerchantIdDev

	payReq := &SmsCfrmReq{
		INFO:  *InitTLReqHeader(headerReq),
		FAGRC: *smsCfrmContent,
	}

	// xml
	tlReqXML, err := utility.ToTLRequestXmlByte(payReq)
	if err != nil {
		resut.ErrMsg = fmt.Sprintln("转化为通联请求XML出错 : ", err)
		return
	}
	bodyByte, err := utility.PostAllinPayXml(tl.TLQuickpayAdrr, tlReqXML)
	if err != nil {
		resut.ErrMsg = fmt.Sprintln("通联服务请求异常 :  ", err)
		return
	}

	smsCfrmResp := &SmsCfrmResp{}
	if err = utility.VerifyTLResponse(bodyByte, smsCfrmResp); err != nil {
		resut.ErrMsg = fmt.Sprintln("通联验签不通过：", err)
		return
	}
	if smsCfrmResp.INFO.RET_CODE != Succe_Code {
		resut.RetCode = smsCfrmResp.INFO.RET_CODE
		resut.ErrMsg = smsCfrmResp.INFO.ERR_MSG
		return
	}

	if smsCfrmResp.Content.RET_CODE != Succe_Code {
		resut.RetCode = smsCfrmResp.Content.RET_CODE
		resut.ErrMsg = smsCfrmResp.Content.ERR_MSG
		return
	}

	resut.IsSuccess = true
	resut.RetInfo = smsCfrmResp.Content.AGRMNO
	return
}

//通联快捷代扣
func (tl *TLPay) Collect(tradeReq QuickTradeReqFASTTRX) (result string, err error) {
	var reqSN = fmt.Sprintf("XxxX_%4d%d", rand.Intn(10000), time.Now().Unix())
	// 请求头
	headerReq := &TLReqHeader{
		TRX_CODE:   "310011", // 交易代码
		LEVEL:      "6",      // 处理级别（0-9  0优先级最低，默认为5）
		REQ_SN:     reqSN,    // 交易流水号（必须全局唯一）
		SIGNED_MSG: "",       // 签名信息
	}

	tradeReq.MERCHANT_ID = utility.MerchantIdDev
	payReq := &QuickTradeReq{
		INFO:    *InitTLReqHeader(headerReq),
		FASTTRX: tradeReq,
	}
	// xml
	tlReqXML, err := utility.ToTLRequestXmlByte(payReq)
	if err != nil {
		fmt.Println("转化为通联请求XML出错 : ", err)
		return "", err
	}
	bodyByte, err := utility.PostAllinPayXml(tl.TLQuickpayAdrr, tlReqXML)
	if err != nil {
		fmt.Println("通联服务请求异常 : ", err)
		return
	}
	payResp := &QuickTradeResp{}
	if err = utility.VerifyTLResponse(bodyByte, payResp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", payResp)
	return
}
