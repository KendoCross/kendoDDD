package logs

import (
	"fmt"

	"github.com/astaxie/beego/logs"
)

type LogTypes int32

const (
	FILE LogTypes = 1 << iota
	CONSO
)

//ConsoleLogs 控制台日志
var ConsoleLogs *logs.BeeLogger
var consoles *logs.BeeLogger

//FileLogs 文件日志
var FileLogs *logs.BeeLogger
var files *logs.BeeLogger

//初始化_日志工具
func InitLogs(logPath string) {

	logsConf := fmt.Sprintf(`{"filename":"%s","maxdays ":7,"level ": 6,"perm":"0777","separate":["error", "warning", "notice", "info", "debug"]}`,
		logPath)
	ConsoleLogs = logs.NewLogger(100)
	ConsoleLogs.SetLogger(logs.AdapterConsole)
	ConsoleLogs.EnableFuncCallDepth(true)

	FileLogs = logs.NewLogger(1000)
	FileLogs.EnableFuncCallDepth(true)
	FileLogs.Async()
	FileLogs.SetLogger(logs.AdapterMultiFile, logsConf)

	consoles = logs.NewLogger(1000)
	consoles.SetLogger(logs.AdapterConsole)
	consoles.EnableFuncCallDepth(true)
	consoles.SetLogFuncCallDepth(3)

	files = logs.NewLogger(1000)
	files.EnableFuncCallDepth(true)
	files.Async()
	files.SetLogFuncCallDepth(3)
	files.SetLogger(logs.AdapterMultiFile, logsConf)

	ConsoleLogs.Info(logsConf)
}

func Info(format string, v ...interface{}) {
	files.Info(format, v...)
	consoles.Info(format, v...)
}

func Error(format string, v ...interface{}) {
	files.Error(format, v...)
	consoles.Error(format, v...)
}

func Warning(format string, v ...interface{}) {
	files.Warning(format, v...)
	consoles.Warning(format, v...)
}

// func PrtRed(prtStr string) {
// 	files.Error(prtStr)
// 	fmt.Println(colors.Red(prtStr))
// }
// func PrtGreen(prtStr string) {
// 	fmt.Println(colors.Green(prtStr))
// }
