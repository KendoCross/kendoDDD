package application

// 应用层的struct ,主要是为了适配表现层的数据格式。正常情况下，领域模型是无法直接暴露给用户直接使用的。

type Page struct {
	Page int `form:"page"`
	Size int `form:"size"`
}
