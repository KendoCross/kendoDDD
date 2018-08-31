package repos

import (
	"../../ddd_interfaces"
	"github.com/astaxie/beego/orm"
)

type protocolRepo struct {
	dbOrm orm.Ormer
}

func NewProtocolRepo(beeOrm orm.Ormer) ddd_interfaces.IProtocolRepo {

	return &protocolRepo{
		dbOrm: beeOrm,
	}
}

func (r *protocolRepo) Add(p *ddd_interfaces.ProtocolOrm) (int64, error) {
	return r.dbOrm.Insert(p)
}

func (r *protocolRepo) Get(pNo string) (*ddd_interfaces.ProtocolOrm, error) {
	protocol := ddd_interfaces.ProtocolOrm{ProtocolNo: pNo}
	err := r.dbOrm.Read(&protocol, "ProtocolNo")
	if err != nil {
		return nil, err
	}
	return &protocol, nil
}
