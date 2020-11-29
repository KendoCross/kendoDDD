package files

import (
	"io/ioutil"
	"time"

	"github.com/KendoCross/kendoDDD/infrastructure/helper"
	"github.com/KendoCross/kendoDDD/interfaces"
	"github.com/google/uuid"
)

func newfileEnByOV(fileInfo interfaces.FileInfo) *fileEntity {
	return &fileEntity{
		FileInfo: &fileInfo,
	}
}

type fileEntity struct {
	*interfaces.FileInfo
}

func (en *fileEntity) EntityID() uuid.UUID {
	uid, err := uuid.Parse(en.Id)
	if err != nil {
		uid = uuid.Nil
	}
	return uid
}

func (en *fileEntity) AddFile(body []byte) (err error) {
	filePath := FileRootPath + time.Now().Format("2006/01/02/") + en.Id + en.FileName
	helper.MakesureFileExist(filePath)
	err = ioutil.WriteFile(filePath, body, 0755)
	if err != nil {
		return
	}
	en.FilePath = filePath
	en.Status = 1
	_, err = fileRepos.AddObj(en.FileInfo)
	return
}
