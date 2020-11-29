package repos_mysql

import (
	"fmt"

	"github.com/KendoCross/kendoDDD/infrastructure/logs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var readEngine *gorm.DB
var writeEngine *gorm.DB

//InitDB 初始化DB引擎
func InitDB() {
	var err error
	readEngine, err = gorm.Open(sqlite.Open("asset/mydb.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("DB error: %v", err))
	}
	writeEngine, err = gorm.Open(sqlite.Open("asset/mydb.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("DB error: %v", err))
	}
	// 迁移 schema
	//writeEngine.AutoMigrate(&interfaces.FileInfo{})
	logs.Info("DB, init end ......")
}
