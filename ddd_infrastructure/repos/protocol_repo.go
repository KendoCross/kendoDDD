package repos

import "github.com/astaxie/beego/orm"
import "../../ddd_interfaces"

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
