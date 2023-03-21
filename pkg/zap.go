/*
 * @Author: kingford
 * @Date: 2023-03-11 01:12:51
 * @LastEditTime: 2023-03-21 16:52:01
 */
package pkg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZap() *zap.Logger {
	config := zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths: []string{
			"stdout",
		},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"), // 使用自定义的时间格式
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	logger, _ := config.Build()

	return logger
}
