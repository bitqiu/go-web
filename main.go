package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/app/models/article"
	"go-web/app/models/tag"
	"go-web/app/models/user"
	"go-web/bootstrap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	db, err := gorm.Open(sqlite.Open("blog.db"))
	if err != nil {
		fmt.Println(err.Error())
	}

	db.AutoMigrate(
		&user.User{},
		&article.Article{},
		&tag.Tag{},
	)
}

func main() {
	// gin 实例
	router := gin.New()
	// 路由初始化
	bootstrap.SetupRoute(router)
	// 启动服务
	router.Run(":3000")
}
