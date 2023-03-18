/*
 * @Author: kingford
 * @Date: 2023-03-11 01:12:51
 * @LastEditTime: 2023-03-11 01:33:19
 */
package pkg

import (
	"go.uber.org/zap"
)

func NewZap() {
	// 创建 logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// 记录日志
	logger.Info("example message")
}
