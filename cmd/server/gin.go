/*
 * @Author: kingford
 * @Date: 2023-03-21 23:27:30
 * @LastEditTime: 2023-03-23 10:38:31
 */
package server

import (
	"context"
	"fmt"
	"go-gin-template/common"
	"go-gin-template/common/global"
	"go-gin-template/internal/middleware"
	"go-gin-template/internal/router"
	"go-gin-template/pkg"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-gin-template/common/config"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	AppRouters = make([]func(), 0)
	GinCmd     = &cobra.Command{
		Use:   "gin",
		Short: "Start a Gin HTTP server",
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			port, _ := cmd.Flags().GetInt("port")
			mode, _ := cmd.Flags().GetString("mode")
			return run(port, mode)
		},
	}
)

func setup() {
	pkg.NewViper()
}

func run(port int, mode string) error {

	initRouter()

	for _, f := range AppRouters {
		f()
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", "0.0.0.0", port),
		Handler: common.Runtime.GetEngine(),
	}

	go func() {
		// 服务连接
		if config.ApplicationConfig.Ssl.Enable {
			if err := srv.ListenAndServeTLS(config.ApplicationConfig.Ssl.Pem, config.ApplicationConfig.Ssl.KeyStr); err != nil && err != http.ErrServerClosed {
				log.Fatal("listen: ", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal("listen: ", err)
			}
		}
	}()

	fmt.Println(pkg.Red(string(global.LogoContent)))
	tip()
	fmt.Println(pkg.Green("Server run at:"))
	fmt.Printf("-  Local:   http://localhost:%d/ \r\n", port)
	fmt.Printf("-  Network: http://%s:%d/ \r\n", pkg.GetLocaHonst(), port)
	fmt.Println(pkg.Green("Swagger run at:"))
	fmt.Printf("-  Local:   http://localhost:%d/swagger/admin/index.html \r\n", port)
	fmt.Printf("-  Network: http://%s:%d/swagger/admin/index.html \r\n", pkg.GetLocaHonst(), port)
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", pkg.GetCurrentTimeStr())
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Printf("%s Shutdown Server ... \r\n", pkg.GetCurrentTimeStr())

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

	return nil

	// 设置Gin的模式
	// gin.SetMode(mode)

	// // 启动HTTP服务器
	// r.Run(fmt.Sprintf(":%d", port))

}

func init() {
	GinCmd.Flags().IntP("port", "p", 8080, "HTTP server port number")
	GinCmd.Flags().StringP("mode", "m", gin.DebugMode, "Gin mode (debug, release, test)")

	AppRouters = append(AppRouters, router.InitRouter)
}

func tip() {
	usageStr := `欢迎使用 ` + pkg.Green(`go-gin-template `+global.Version) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s \n\n", usageStr)
}

func initRouter() {
	var r *gin.Engine
	h := common.Runtime.GetEngine()
	if h == nil {
		h = gin.New()
		common.Runtime.SetEngine(h)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	r.Use(middleware.Sentinel())

}
