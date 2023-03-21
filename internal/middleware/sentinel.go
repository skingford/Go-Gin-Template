/*
 * @Author: kingford
 * @Date: 2023-03-21 14:35:52
 * @LastEditTime: 2023-03-21 14:36:11
 */
package middleware

import (
	"github.com/alibaba/sentinel-golang/core/system"
	sentinel "github.com/alibaba/sentinel-golang/pkg/adapters/gin"
	"github.com/gin-gonic/gin"

	log "github.com/go-admin-team/go-admin-core/logger"
)

// Sentinel 限流
func Sentinel() gin.HandlerFunc {
	if _, err := system.LoadRules([]*system.Rule{
		{
			MetricType:   system.InboundQPS,
			TriggerCount: 200,
			Strategy:     system.BBR,
		},
	}); err != nil {
		log.Fatalf("Unexpected error: %+v", err)
	}
	return sentinel.SentinelMiddleware(
		sentinel.WithBlockFallback(func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(200, map[string]interface{}{
				"msg":  "too many request; the quota used up!",
				"code": 500,
			})
		}),
	)
}
