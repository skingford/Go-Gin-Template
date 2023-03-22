/*
 * @Author: kingford
 * @Date: 2023-03-21 15:00:24
 * @LastEditTime: 2023-03-22 17:02:48
 */
package system

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware), 0)
)

// 路由示例
func InitSystemRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {

	// 无需认证的路由
	SystemNoCheckRoleRouter(r)

	// 需要认证的路由
	SystemCheckRoleRouter(r, authMiddleware)
	return r
}

// 无需认证的路由示例
func SystemNoCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/v1")
	for _, f := range routerNoCheckRole {
		f(v1)
	}
}

// 需要认证的路由示例
func SystemCheckRoleRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/v1")
	for _, f := range routerCheckRole {
		f(v1, authMiddleware)
	}
}
