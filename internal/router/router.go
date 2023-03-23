/*
 * @Author: kingford
 * @Date: 2023-03-21 11:55:53
 * @LastEditTime: 2023-03-23 19:55:00
 */
package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	routerNoCheckRole = make([]func(v1 *gin.RouterGroup, logger *zap.Logger), 0)
	// routerCheckRole   = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware), 0)
)

// InitRouter 路由初始化
func InitRouter(r *resource) func() {
	return func() {
		// 注册业务路由
		v1 := r.mux.Group("/v1")

		InitSystemRouter(v1, r.logger)
	}
}

// 路由示例
func InitSystemRouter(v1 *gin.RouterGroup, logger *zap.Logger) {

	// 无需认证的路由
	NoCheckRoleRouter(v1, logger)

	// 需要认证的路由
	// CheckRoleRouter(v1, authMiddleware)

}

// 无需认证的路由示例
func NoCheckRoleRouter(v1 *gin.RouterGroup, logger *zap.Logger) {
	for _, f := range routerNoCheckRole {
		f(v1, logger)
	}
}

// 需要认证的路由示例
// func CheckRoleRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
// 	// 可根据业务需求来设置接口版本
// 	v1 := r.Group("/v1")
// 	for _, f := range routerCheckRole {
// 		f(v1, authMiddleware)
// 	}
// }
