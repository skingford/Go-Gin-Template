/*
 * @Author: kingford
 * @Date: 2023-03-21 14:52:34
 * @LastEditTime: 2023-03-23 09:58:18
 */
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"

	"go-gin-template/common/config"
)

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     config.ApplicationConfig.Ssl.Domain,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	}
}
