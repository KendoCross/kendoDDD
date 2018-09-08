package crossutting

import (
	"../ddd_infrastructure/repos"
	"../ddd_interfaces"
	"github.com/astaxie/beego/orm"
)

var (
	RepoFac *RepoFactory = &RepoFactory{}
)

type RepoFactory struct {
	protocolRepo ddd_interfaces.IProtocolRepo
}

func (r *RepoFactory) Init() {
	o := orm.NewOrm()
	o.Using("default")

	r.protocolRepo = repos.NewProtocolRepo(o)
}

//支付协议仓库
func (r *RepoFactory) GetProtocolRepo() ddd_interfaces.IProtocolRepo {
	return r.protocolRepo
}
