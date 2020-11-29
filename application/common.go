//
// 应用层，核心思路能够给不同的表现层提供通用的数据。
//
package application

import (
	"github.com/KendoCross/kendoDDD/domain/testcmd"
	eh "github.com/looplab/eventhorizon"
)

const (
	ConstGRPCHeaderReqUuid = "requuid"
	TestOnlyCmd            = string(testcmd.TestOnlyCmdType)

	/*******  文件存储领域 **********/
	//AddAgentCmd       = string(agent.AddAgentCmdType)   //注册添加代理人

)

//api2Cmd  路由到领域CMD的映射
var api2Cmd map[string]eh.CommandType
