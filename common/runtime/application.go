/*
 * @Author: kingford
 * @Date: 2023-03-22 16:31:18
 * @LastEditTime: 2023-03-23 19:22:08
 */
package runtime

import (
	"net/http"
)

type Application struct {
	engine http.Handler
}

// NewConfig 默认值
func NewConfig() *Application {
	return &Application{}
}

// SetEngine 设置路由引擎
func (e *Application) SetEngine(engine http.Handler) {
	e.engine = engine
}

// GetEngine 获取路由引擎
func (e *Application) GetEngine() http.Handler {
	return e.engine
}
