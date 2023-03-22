/*
 * @Author: kingford
 * @Date: 2023-03-22 16:40:27
 * @LastEditTime: 2023-03-22 16:40:32
 */
package runtime

import (
	"net/http"
)

type Runtime interface {

	// SetEngine 使用的路由
	SetEngine(engine http.Handler)
	GetEngine() http.Handler
}
