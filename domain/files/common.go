package files

import (
	"strings"

	"github.com/KendoCross/kendoDDD/infrastructure"
	eh "github.com/looplab/eventhorizon"
)

const AgentAggregateType = eh.AggregateType("AggregateType_Files")
const FileRootPath = "/var/data/kendoDDD/"

//单例聚合根的特殊实例，ID为NIL
var SingleFilesAgg *filesAggregate

// 仓储（能封装成工作单元更好）
var captchaRepo = infrastructure.RepoFac.CaptchaRepo
var fileRepos = infrastructure.RepoFac.FilesRepo

func getcontentType(fileName string) (contentType string) {
	if strings.HasSuffix(fileName, ".bmp") {
		contentType = "image/bmp"
	} else if strings.HasSuffix(fileName, ".gif") {
		contentType = "image/gif"
	} else if strings.HasSuffix(fileName, ".jpeg") || strings.HasSuffix(fileName, ".jpg") || strings.HasSuffix(fileName, ".png") {
		contentType = "image/jpg"
	} else if strings.HasSuffix(fileName, ".html") {
		contentType = "text/html"
	} else if strings.HasSuffix(fileName, ".txt") {
		contentType = "text/plain"
	} else if strings.HasSuffix(fileName, ".vsd") {
		contentType = "application/vnd.visio"
	} else if strings.HasSuffix(fileName, ".pptx") || strings.HasSuffix(fileName, ".ppt") {
		contentType = "application/vnd.ms-powerpoint"
	} else if strings.HasSuffix(fileName, ".docx") || strings.HasSuffix(fileName, ".doc") {
		contentType = "application/msword"
	} else if strings.HasSuffix(fileName, ".xml") {
		contentType = "text/xml"
	}
	return
}
