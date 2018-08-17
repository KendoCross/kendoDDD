package ddd_infrastructure

import (
	"fmt"

	"github.com/astaxie/beego/cache"
)

var MemoryCache cache.Cache

func init() {
	var err error
	MemoryCache, err = cache.NewCache("memory", `{"interval":64}`)
	if err != nil {
		fmt.Println("严重异常！缓存初始化失败！")
		panic("严重异常！缓存初始化失败！")
	}
}
