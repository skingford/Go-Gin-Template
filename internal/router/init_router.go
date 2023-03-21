/*
 * @Author: kingford
 * @Date: 2023-03-21 11:55:53
 * @LastEditTime: 2023-03-21 15:32:06
 */
package router

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
)

// InitRouter 路由初始化，不要怀疑，这里用到了
func InitRouter() {
	var r *gin.Engine
	h := sdk.Runtime.GetEngine()
	if h == nil {
		log.Fatal("not found engine...")
		os.Exit(-1)
	}

	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	// the jwt middleware
	// authMiddleware, err := common.AuthInit()
	// if err != nil {
	// 	log.Fatalf("JWT Init Error, %s", err.Error())
	// }

	// 注册业务路由
	// TODO: 这里可存放业务路由，里边并无实际路由只有演示代码
	examplesNoCheckRoleRouter(r)
}
