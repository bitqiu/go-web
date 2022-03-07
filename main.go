package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-web/bootstrap"
	"go-web/global"
	"go-web/pkg/config"
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

	sqliteDB := sqlite.Open("blog.db")

	var err error
	global.DB, err = gorm.Open(sqliteDB)

	if err != nil {
		fmt.Println(err.Error())
	}

	config.Viper = viper.New()
	config.Viper.SetConfigName("config")
	config.Viper.SetConfigType("toml")
	config.Viper.AddConfigPath("./config")
	err = config.Viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	// gin 实例
	router := gin.New()
	// 路由初始化
	bootstrap.SetupRoute(router)
	// 启动服务
	router.Run(":3000")
}
