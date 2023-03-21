/*
 * @Author: kingford
 * @Date: 2023-03-21 15:35:00
 * @LastEditTime: 2023-03-21 19:24:54
 */
package router

import (
	"go-gin-template/internal/api"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysApiRouter)
}

// registerSysApiRouter
func registerSysApiRouter(v1 *gin.RouterGroup) {
	api := api.SysApi{}
	r := v1.Group("/sys-api")
	{
		r.GET("", api.Find)
		r.GET("/:id", api.First)
		r.POST("/:id", api.Create)
		r.PUT("/:id", api.Update)
	}
}
