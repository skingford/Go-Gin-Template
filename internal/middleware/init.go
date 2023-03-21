/*
 * @Author: kingford
 * @Date: 2023-03-21 14:43:02
 * @LastEditTime: 2023-03-21 14:49:06
 */
package middleware

import (
	"github.com/gin-gonic/gin"
)

const (
	JwtTokenCheck   string = "JwtToken"
	RoleCheck       string = "AuthCheckRole"
	PermissionCheck string = "PermissionAction"
)

func InitMiddleware(r *gin.Engine) {
	r.Use(DemoEvn())
	// 数据库链接
	// r.Use(WithContextDb)
	// 日志处理
	// r.Use(LoggerToFile())
	// 自定义错误处理
	r.Use(CatchError)
	// NoCache is a middleware function that appends headers
	r.Use(NoCache)
	// 跨域处理
	r.Use(Options)
	// Secure is a middleware function that appends security
	r.Use(Secure)
	// 链路追踪
	//r.Use(middleware.Trace())
	//sdk.Runtime.SetMiddleware(JwtTokenCheck, (*jwt.GinJWTMiddleware).MiddlewareFunc)

}
