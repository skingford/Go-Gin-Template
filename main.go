/*
 * @Author: kingford
 * @Date: 2023-03-13 11:39:16
 * @LastEditTime: 2023-03-21 17:00:22
 */
package main

import (
	"go-gin-template/cmd"
	"go-gin-template/pkg"
)

func main() {

	logger := pkg.NewZap()

	defer logger.Sync()

	logger.Info("==== Starting======")

	cmd.Execute()

	// // 使用错误处理中间件
	// r.Use(middleware.LoggingMiddleware(logger), middleware.ErrorMiddleware())

	// // 定义一个抛出错误的路由
	// r.GET("/test", func(c *gin.Context) {
	// 	// 抛出错误并传入 code 参数
	// 	c.Set("code", http.StatusBadRequest)
	// 	panic("something went wrong")
	// })

	// r.GET("/hello", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello, world!",
	// 	})
	// })

	// r.POST("/hello", func(c *gin.Context) {
	// 	var request struct {
	// 		Name string `json:"name"`
	// 	}
	// 	c.BindJSON(&request)

	// 	c.JSON(200, gin.H{
	// 		"message": "Hello, " + request.Name + "!",
	// 	})
	// })

	// // 启动服务
	// r.Run(":8080")
}
