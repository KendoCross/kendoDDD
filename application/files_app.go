package application

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/KendoCross/kendoDDD/domain/files"
	"github.com/KendoCross/kendoDDD/infrastructure"
	"github.com/KendoCross/kendoDDD/infrastructure/bus"
	"github.com/KendoCross/kendoDDD/infrastructure/ddd"
	"github.com/KendoCross/kendoDDD/infrastructure/logs"
	eh "github.com/looplab/eventhorizon"
)

func GetFileById(id string) (file FileInfo, has bool, err error) {
	obj, has, err := infrastructure.RepoFac.FilesRepo.GetById(id)
	if err != nil {
		logs.Error("FilesRepo GetById ERR:%v", err)
		return
	}
	if !has {
		return
	}

	file.ContentType = obj.ContentType
	file.FileName = obj.FileName
	file.FilePath = obj.FilePath
	file.Size = obj.Size
	return
}

func AddFile(parm AddFileForm) (fileId string, err error) {
	fileInfo := files.FileInfos{}
	f, err := parm.UpFile.Open()
	if err != nil {
		return
	}
	defer f.Close()
	fileInfo.FileBody, err = ioutil.ReadAll(f)
	if err != nil {
		return
	}
	fileInfo.ContentType = parm.UpFile.Header.Get("Content-Type")
	fileInfo.Size = int(parm.UpFile.Size)
	fileInfo.FileName = parm.UpFile.Filename

	fileId, err = files.SingleFilesAgg.AddNewFile(fileInfo)
	if err != nil {
		logs.Error("SingleFilesAgg AddFile ERR:", err.Error())
		return
	}
	return
}

func AddFileCommand(parm AddFileForm) (err error) {
	//looplab框架根据命令类型，创建命令实例
	cmd, err := eh.CreateCommand(files.AddFileCmdType)
	if err != nil {
		err = fmt.Errorf("could not create command: %w", err)
		return err
	}

	//命令进行校验
	if vldt, ok := cmd.(ddd.Validator); ok {
		err = vldt.Verify()
		if err != nil {
			return
		}
	}
	//通过命令总线，将命令发布出去，至于谁订阅了该命令则不关系
	if err = bus.HandleCommand(context.Background(), cmd); err != nil {
		return err
	}
	return
}
