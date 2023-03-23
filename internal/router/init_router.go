/*
 * @Author: kingford
 * @Date: 2023-03-21 11:55:53
 * @LastEditTime: 2023-03-23 10:04:33
 */
package router

import (
	"fmt"
	"log"
	"os"

	"go-gin-template/common"
	"go-gin-template/internal/router/system"

	"github.com/gin-gonic/gin"
)

// InitRouter 路由初始化
func InitRouter() {
	var r *gin.Engine
	h := common.Runtime.GetEngine()
	if h == nil {
		log.Fatal("not found engine...")
		os.Exit(-1)
	}
	switch v := h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
		fmt.Printf("value %+v\n", v)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	// 注册业务路由
	system.SystemNoCheckRoleRouter(r)
}
