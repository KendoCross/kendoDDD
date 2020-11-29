package interfaces

type IFileInfoRepo interface {
	//新增，
	AddObj(obj *FileInfo) (num int64, err error)
	//单条查询
	GetById(id string) (obj FileInfo, has bool, err error)

	Find(parm FindParmFiles) (objs []FileInfo, total int64, err error)
}

type FileInfo struct {
	Id          string `gorm:"primarykey"`
	FileName    string `json:"file_name" gorm:"comment('文件名') VARCHAR(255)"`
	FilePath    string `json:"file_path" gorm:"comment('文件目录') VARCHAR(256)"`
	ContentType string `json:"content_type" gorm:"comment('文件类型') VARCHAR(64)"`
	Size        int    `json:"size" gorm:"comment('文件大小') INT"`
	Status      int    `json:"status" gorm:"comment('文件保存状态 1：成功 2：失败') TINYINT"`
	ErrMsg      string `json:"err_msg" gorm:"comment('异常说明') VARCHAR(128)"`
	CreateAt    int64  `json:"create_at" gorm:"comment('创建时间') BIGINT"`
	UpdateAt    int64  `json:"update_at" gorm:"comment('更新时间') BIGINT"`
}

type FindParmFiles struct {
	Pages
}

func (a *FileInfo) TableName() string {
	return "t_sys_files"
}
