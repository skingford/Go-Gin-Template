/*
 * @Author: kingford
 * @Date: 2023-03-13 11:39:16
 * @LastEditTime: 2023-03-18 23:52:43
 */
package main

import (
	"go-gin-template/internal/middleware"
	"net/http"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths: []string{
			"stdout",
		},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"), // 使用自定义的时间格式
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	logger, _ := config.Build()
	defer logger.Sync()

	logger.Info("Hello world", zap.Time("timestamp", time.Now()))

	// 使用错误处理中间件
	r.Use(middleware.LoggingMiddleware(logger), middleware.ErrorMiddleware())

	// 定义一个抛出错误的路由
	r.GET("/test", func(c *gin.Context) {
		// 抛出错误并传入 code 参数
		c.Set("code", http.StatusBadRequest)
		panic("something went wrong")
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	r.POST("/hello", func(c *gin.Context) {
		var request struct {
			Name string `json:"name"`
		}
		c.BindJSON(&request)

		c.JSON(200, gin.H{
			"message": "Hello, " + request.Name + "!",
		})
	})

	// 启动服务
	r.Run(":8080")
}
