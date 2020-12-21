package main

import (
	"goletter-go-api/app/http/middlewares"
	"goletter-go-api/bootstrap"
	"goletter-go-api/config"
	c "goletter-go-api/pkg/config"
	"net/http"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	// 初始化 SQL
	bootstrap.SetupDB()

	// 初始化路由绑定
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
