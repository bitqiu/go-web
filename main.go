package main

import "go-web/cmd"

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
	cmd.Execute()
}
