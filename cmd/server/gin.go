/*
 * @Author: kingford
 * @Date: 2023-03-21 23:27:30
 * @LastEditTime: 2023-03-23 20:34:36
 */
package server

import (
	"context"
	"fmt"
	"go-gin-template/common/config"
	"go-gin-template/common/global"
	"go-gin-template/internal/router"
	"go-gin-template/pkg"
	"go-gin-template/pkg/env"
	"go-gin-template/pkg/logger"
	"go-gin-template/pkg/shutdown"
	"go-gin-template/pkg/timeutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	GinCmd = &cobra.Command{
		Use:   "gin",
		Short: "Start a Gin HTTP server",
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetInt("port")
			mode, _ := cmd.Flags().GetString("mode")
			run(port, mode)
		},
	}
)

func init() {
	GinCmd.Flags().IntP("port", "p", 8080, "HTTP server port number")
	GinCmd.Flags().StringP("mode", "m", gin.DebugMode, "Gin mode (debug, release, test)")

}

func setup() {
	pkg.NewViper()
}

func run(port int, mode string) {
	if config.ApplicationConfig.Server.Mode == config.ModeProd.String() {
		gin.SetMode(gin.ReleaseMode)
	}

	s, err := HttpServer()

	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", "0.0.0.0", port),
		Handler: s.Mux,
	}

	go func() {
		// 服务连接
		if config.ApplicationConfig.Ssl.Enable {
			if err := srv.ListenAndServeTLS(config.ApplicationConfig.Ssl.Pem, config.ApplicationConfig.Ssl.KeyStr); err != nil && err != http.ErrServerClosed {
				s.Logger.Error("http server startup err", zap.Error(err))
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				s.Logger.Error("http server startup err", zap.Error(err))
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

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			fmt.Printf("%s Shutdown Server ... \r\n", pkg.GetCurrentTimeStr())

			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("server shutdown err", zap.Error(err))
			}
		},
	)

}

func tip() {
	usageStr := `欢迎使用 ` + pkg.Green(`go-gin-template `+global.Version) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s \n\n", usageStr)
}

func HttpServer() (*router.Server, error) {
	// 初始化 access logger
	accessLogger, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", config.ProjectName, env.Active().Value())),
		logger.WithTimeLayout(timeutil.CSTLayout),
		logger.WithFileP(config.ProjectAccessLogFile),
	)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = accessLogger.Sync()

	}()

	return router.NewHTTPServer(accessLogger)
}
