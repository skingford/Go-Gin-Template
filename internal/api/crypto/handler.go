/*
 * @Author: kingford
 * @Date: 2023-03-23 16:25:38
 * @LastEditTime: 2023-03-23 16:50:26
 */
package crypto

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler interface {
	i()

	// Md5 加密
	// @Tags Helper
	// @Router /helper/md5/{str} [get]
	Md5(c *gin.Context)

	// Sign 签名
	// @Tags Helper
	// @Router /helper/sign [post]
	Sign(c *gin.Context)
}

type handler struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) i() {}
