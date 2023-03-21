/*
 * @Author: kingford
 * @Date: 2023-03-22 00:44:19
 * @LastEditTime: 2023-03-22 00:44:32
 */
package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		log.Printf("[%s] %s %s %v\n", end.Format("2006/01/02 - 15:04:05"), c.Request.Method, c.Request.URL.Path, latency)
	}
}
