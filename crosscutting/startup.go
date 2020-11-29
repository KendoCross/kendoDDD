package crosscutting

import (
	"fmt"

	"github.com/KendoCross/kendoDDD/domain"
	"github.com/KendoCross/kendoDDD/infrastructure/logs"
	mysql "github.com/KendoCross/kendoDDD/infrastructure/repos_mysql"

	//redis "github.com/KendoCross/kendoDDD/infrastructure/repos_redis"

	"github.com/spf13/viper"
)

func StartUp() {

	viper.SetConfigFile("conf/config.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Error while loading config file [conf/config.yaml]: %s", err.Error()))
	}

	// 日志组件
	logs.InitLogs(viper.GetString("LOG_FILE"))

	//仓储层DB引擎
	mysql.InitDB()
	//redis.InitRedis()

	//最后启动领域层
	domain.StartUp()
}
