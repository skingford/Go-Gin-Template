/*
 * @Author: kingford
 * @Date: 2023-03-23 17:56:04
 * @LastEditTime: 2023-03-23 20:25:11
 */
package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var AppRouters = make([]func(), 0)

type resource struct {
	logger *zap.Logger
	mux    *gin.Engine
}

type Server struct {
	Logger *zap.Logger
	Mux    http.Handler
}

func NewResource(mux *gin.Engine, logger *zap.Logger) *resource {
	return &resource{mux: mux, logger: logger}
}

func NewServer(mux http.Handler, logger *zap.Logger) *Server {
	return &Server{Mux: mux, Logger: logger}
}

func NewHTTPServer(logger *zap.Logger) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	// db

	// cache

	mux := gin.New()
	r := NewResource(mux, logger)
	s := NewServer(mux, logger)

	// 注册路由
	AppRouters = append(AppRouters, InitRouter(r))

	for _, f := range AppRouters {
		f()
	}

	return s, nil
}
