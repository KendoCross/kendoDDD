package ddd_interfaces

import (
	"time"
)

type IProtocolRepo interface {
	Add(p *ProtocolOrm) (int64, error)
	Get(pNo string) (*ProtocolOrm, error)
}

type ProtocolOrm struct {
	ID int `orm:"auto;pk;column(ID);" `
	//协议号
	ProtocolNo string `orm:"column(ProtocolNo)"`
	//支付平台类型
	AisleType string `orm:"column(AisleType)"`
	//同一平台细分的子类
	AisleBranch int `orm:"column(AisleBranch)"`
	//支付平台生成的协议号
	AisleProtocol string `orm:"column(AisleProtocol)"`
	BankId        string `orm:"column(BankId)"`
	//银行code，各平台不一样
	BankCode string `orm:"column(BankCode)"`
	//账户类型
	AccountType int `orm:"column(AccountType)"`
	//账户属性
	AccountProp int `orm:"column(AccountProp)"`
	//银行卡号
	BankCardNo string `orm:"column(BankCardNo)"`
	//预留手机号
	ReservePhone string `orm:"column(ReservePhone)"`
	//银行卡预留的证件类型
	IDCardType int `orm:"column(IDCardType)"`
	//银行卡开户姓名
	CardName string `orm:"column(CardName)"`
	//银行卡预留的证件号码
	IDCardNo string `orm:"column(IDCardNo)"`
	//银行所在城市代号
	CityNo string `orm:"column(CityNo)"`
	//银行支行名称
	BranchNm string `orm:"column(BranchNm)"`
	//协议状态
	Status int `orm:"column(Status)"`
	//冻结
	Freeze bool `orm:"column(Freeze)"`
	//商户预留信息
	Merrem string `orm:"column(Merrem)"`
	//备注
	Remark string `orm:"column(Remark)"`
	//提示标识
	Notes string `orm:"column(Notes)"`
	//删除
	DelMark bool `orm:"column(DelMark)"`
	//新增时间
	AddTime time.Time `orm:"auto_now;type(datetime);column(AddTime)"`
	//更新时间
	UpdTime time.Time `orm:"auto_now;type(datetime);column(UpdTime)"`
}

func (p ProtocolOrm) TableName() string {
	return "Protocol"
}
