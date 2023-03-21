/*
 * @Author: kingford
 * @Date: 2023-03-21 14:37:04
 * @LastEditTime: 2023-03-21 14:37:18
 */
package middleware

import (
	"net/http"
	"strings"

	"github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestId 自动增加requestId
func RequestId(trafficKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}
		requestId := c.GetHeader(trafficKey)
		if requestId == "" {
			requestId = c.GetHeader(strings.ToLower(trafficKey))
		}
		if requestId == "" {
			requestId = uuid.New().String()
		}
		c.Request.Header.Set(trafficKey, requestId)
		c.Set(trafficKey, requestId)
		c.Set(pkg.LoggerKey,
			logger.NewHelper(logger.DefaultLogger).
				WithFields(map[string]interface{}{
					trafficKey: requestId,
				}))
		c.Next()
	}
}
