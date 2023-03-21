/*
 * @Author: kingford
 * @Date: 2023-03-21 14:52:34
 * @LastEditTime: 2023-03-21 17:58:05
 */
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"

	"go-gin-template/core/config"
)

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     config.SslConfig.Domain,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	}
}
