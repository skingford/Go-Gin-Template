/*
 * @Author: kingford
 * @Date: 2023-03-21 23:27:30
 * @LastEditTime: 2023-03-22 01:16:48
 */
package cmd

import (
	"fmt"
	"go-gin-template/pkg"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var ginCmd = &cobra.Command{
	Use:   "gin",
	Short: "Start a Gin HTTP server",
	PreRun: func(cmd *cobra.Command, args []string) {
		setup()
	},
	Run: startGin,
}

func setup() {
	pkg.NewViper()
}

func startGin(cmd *cobra.Command, args []string) {
	port, _ := cmd.Flags().GetInt("port")
	mode, _ := cmd.Flags().GetString("mode")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	// 设置Gin的模式
	gin.SetMode(mode)

	// 启动HTTP服务器
	r.Run(fmt.Sprintf(":%d", port))
}

func init() {
	ginCmd.Flags().IntP("port", "p", 8080, "HTTP server port number")
	ginCmd.Flags().StringP("mode", "m", gin.DebugMode, "Gin mode (debug, release, test)")
}
