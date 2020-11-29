package application

import "mime/multipart"

type FileInfo struct {
	FileName    string `json:"file_name" xorm:"comment('文件名称') VARCHAR(255)"`
	FilePath    string `json:"file_path" xorm:"comment('文件目录') VARCHAR(256)"`
	ContentType string `json:"content_type" xorm:"comment('文件类型') VARCHAR(64)"`
	Size        int    `json:"size" xorm:"comment('文件大小') INT"`
}

type AddFileForm struct {
	UpFile *multipart.FileHeader `form:"up_file"`
	Remark string                `form:"remark"`
}
