package repos_mysql

import (
	"errors"
	"time"

	"github.com/KendoCross/kendoDDD/interfaces"
	"gorm.io/gorm"
)

type fileRepo struct {
}

func NewfileRepo() *fileRepo {
	return new(fileRepo)
}

//新增，
func (r *fileRepo) AddObj(obj *interfaces.FileInfo) (num int64, err error) {
	now := time.Now().Unix()
	obj.CreateAt = now
	obj.UpdateAt = now
	tx := writeEngine.
		Create(obj)
	num, err = tx.RowsAffected, tx.Error
	return
}

//单条查询
func (r *fileRepo) GetById(id string) (obj interfaces.FileInfo, has bool, err error) {
	err = readEngine.
		Where("id=?", id).
		First(&obj).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
		return
	}
	has = true
	return
}

func (r *fileRepo) Find(parm interfaces.FindParmFiles) (objs []interfaces.FileInfo, total int64, err error) {
	objs = make([]interfaces.FileInfo, 0)
	err = readEngine.Order("created desc").
		Limit(parm.PageSize).
		Offset(parm.Page * parm.PageSize).
		Find(&objs).Error
	return
}
