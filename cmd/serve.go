package cmd

import (
	"go-web/pkg/console"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {

	// gin 实例
	router := gin.Default()

	// 声明路由
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	})

	// 运行服务器
	err := router.Run(":8080")
	if err != nil {
		console.Exit("无法启动服务 ERROR:" + err.Error())
	}
}
