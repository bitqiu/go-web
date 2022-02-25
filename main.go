package main

import (
	"github.com/gin-gonic/gin"
	"go-web/bootstrap"
)

//go:generate swag init --parseDependency --parseDepth=6

// @title go-web API
// @version 0.0.1
// @description go-web api 文档

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

func init() {
	// TODO 全局初始化 例 配置读取
}

func main() {
	// gin 实例
	router := gin.New()
	// 路由初始化
	bootstrap.SetupRoute(router)
	// 启动服务
	router.Run(":3000")
}
