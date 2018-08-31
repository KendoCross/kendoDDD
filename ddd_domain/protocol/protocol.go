package protocol

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"../../crossutting"
	infra "../../ddd_infrastructure"
	"../../ddd_interfaces"
	"../../dddcore"
	"../domainser"
	"github.com/google/uuid"
)

// User aggregate root
type Protocol struct {
	id      uuid.UUID
	version int
	changes []*dddcore.Event
	SignVm  *SignModel
}

func (p *Protocol) Sign() (string, error) {

	smsContent := &domainser.QuickPaySignSmsContent{
		BankCode:     p.SignVm.BankCode,
		ACCOUNT_TYPE: p.SignVm.AccountType,
		ACCOUNT_NO:   strings.Replace(p.SignVm.BankCardNo, " ", "", -1),
		ACCOUNT_NAME: p.SignVm.CardName,
		ACCOUNT_PROP: p.SignVm.AccountProp,
		ID_TYPE:      p.SignVm.IDCardType,
		ID:           p.SignVm.IDCardNo,
		TEL:          p.SignVm.ReservePhone,
		MERREM:       p.SignVm.Merrem,
		REMARK:       p.SignVm.Remark,
		AisleType:    p.SignVm.AisleType,
	}
	result := domainser.Allinpay.SignGetCode(smsContent)
	jsonByte, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	if result.IsSuccess {
		infra.MemoryCache.Put(result.RetInfo, *smsContent, 5*time.Minute)
	}
	return string(jsonByte), nil
}

func (p *Protocol) CfrmSign(reqSn string, verCode string) (string, error) {

	result := domainser.ApiResp{}

	if !infra.MemoryCache.IsExist(reqSn) {
		result.ErrMsg = "无此请求号的记录！"
		jsonByte, err := json.Marshal(result)
		if err != nil {
			return "", err
		}
		return string(jsonByte), nil
	}

	smsContent := infra.MemoryCache.Get(reqSn).(domainser.QuickPaySignSmsContent)
	smsCfrm := &domainser.SmsCfrmContent{
		VERCODE:  verCode,
		SRCREQSN: reqSn,
	}

	result = domainser.Allinpay.SignCfrm(smsCfrm)
	if result.IsSuccess {

		accountType, _ := strconv.Atoi(smsContent.ACCOUNT_TYPE)
		accountProp, _ := strconv.Atoi(smsContent.ACCOUNT_PROP)
		idTp, _ := strconv.Atoi(smsContent.ID_TYPE)

		protocolOrm := &ddd_interfaces.ProtocolOrm{
			ProtocolNo:    uuid.New().String(),
			AisleType:     smsContent.AisleType,
			AisleProtocol: result.RetInfo,
			BankCode:      smsContent.BankCode,
			AccountType:   accountType,
			AccountProp:   accountProp,
			BankCardNo:    smsContent.ACCOUNT_NO,
			ReservePhone:  smsContent.TEL,
			IDCardType:    idTp,
			CardName:      smsContent.ACCOUNT_NAME,
			IDCardNo:      smsContent.ID,
			Merrem:        smsContent.MERREM,
			Remark:        smsContent.REMARK,
			Status:        1,
		}

		protocolRepo := crossutting.Repo.GetProtocolRepo()
		_, err := protocolRepo.Add(protocolOrm)

		if err != nil {
			result.ErrMsg = "保存数据库失败，严重异常！"
		}

		result.RetInfo = protocolOrm.ProtocolNo
	}

	jsonByte, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(jsonByte), nil

	// eventEnvelop, err := dddcore.NewEvent(p.id, fmt.Sprintf("%T", p), p.version, protocolOrm)
	// if err != nil {
	// 	return "-1", err
	// }
	// p.changes = append(p.changes, eventEnvelop)

}

func (p *Protocol) GetInfoByNo(protocolNo string) (*ddd_interfaces.ProtocolOrm, error) {
	protocolRepo := crossutting.Repo.GetProtocolRepo()
	return protocolRepo.Get(protocolNo)
}

func (p *Protocol) Changes() []*dddcore.Event {
	return p.changes
}

// New creates
func New() *Protocol {
	return &Protocol{}
}
