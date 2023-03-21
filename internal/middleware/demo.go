/*
 * @Author: kingford
 * @Date: 2023-03-21 14:46:31
 * @LastEditTime: 2023-03-21 17:57:57
 */
package middleware

import (
	"net/http"

	"go-gin-template/core/config"

	"github.com/gin-gonic/gin"
)

func DemoEvn() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if config.ApplicationConfig.Mode == "demo" {
			if method == "GET" ||
				method == "OPTIONS" ||
				c.Request.RequestURI == "/api/v1/login" ||
				c.Request.RequestURI == "/api/v1/logout" {
				c.Next()
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  "谢谢您的参与，但为了大家更好的体验，所以本次提交就算了吧！\U0001F600\U0001F600\U0001F600",
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
