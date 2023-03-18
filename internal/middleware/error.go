/*
 * @Author: kingford
 * @Date: 2023-03-11 01:36:47
 * @LastEditTime: 2023-03-18 23:49:36
 */
package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义错误响应格式
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 定义错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 获取 code 参数，默认为 500
				code, ok := c.Get("code")
				if !ok {
					code = http.StatusInternalServerError
				}
				// 将错误转换为结构体并返回 JSON 格式响应
				c.JSON(code.(int), ErrorResponse{
					Code:    code.(int),
					Message: fmt.Sprintf("%s", err),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
