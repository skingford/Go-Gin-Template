/*
 * @Author: kingford
 * @Date: 2023-03-21 23:27:30
 * @LastEditTime: 2023-03-22 02:24:46
 */
package cmd

import (
	"fmt"
	"go-gin-template/pkg"
	"os/exec"

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
	watch, _ := cmd.Flags().GetBool("watch")

	if watch {

		// 创建air命令
		airCmd := exec.Command("air")

		// 添加-w标志
		airCmd.Args = append(airCmd.Args, "-w")

		// 启动air命令
		if err := airCmd.Start(); err != nil {
			panic(fmt.Errorf("Failed to start air command: %v", err))
		}

		// 创建HTTP服务器
		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, world!",
			})
		})

		// 设置Gin的模式
		gin.SetMode(mode)

		// 等待air命令结束
		if err := airCmd.Wait(); err != nil {
			panic(fmt.Errorf("Air command exited with error: %v", err))
		}
	} else {
		// 创建HTTP服务器
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
}

func init() {
	// 添加-p标志
	ginCmd.Flags().IntP("port", "p", 8080, "HTTP server port number")

	// 添加-m标志
	ginCmd.Flags().StringP("mode", "m", gin.DebugMode, "Gin mode (debug, release, test)")

	// 添加-w标志
	ginCmd.Flags().BoolP("watch", "w", false, "Watch for file changes")

}
