/*
 * @Author: kingford
 * @Date: 2023-03-23 16:37:53
 * @LastEditTime: 2023-03-23 19:32:28
 */
package router

import (
	"go-gin-template/internal/api/crypto"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerCryptoApiRouter)
}

// registerSysApiRouter
func registerCryptoApiRouter(v1 *gin.RouterGroup, logger *zap.Logger) {
	handler := crypto.New(logger)

	r := v1.Group("/crypto")
	{
		r.GET("/md5/:str", handler.Md5)
		r.GET("/sign", handler.Sign)
	}
}
