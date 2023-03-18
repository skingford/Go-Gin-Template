package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Logger struct {
	log   *zap.Logger
	reqID string
}

type LogRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type RequestBody struct {
	body []byte
}

func (r *RequestBody) Read(p []byte) (n int, err error) {
	return bytes.NewReader(r.body).Read(p)
}

func (r *RequestBody) Close() error {
	return nil
}

func LoggingMiddleware(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		reqID := c.GetHeader("X-Request-ID")
		logger := &Logger{
			log:   log,
			reqID: reqID,
		}

		c.Set("logger", logger)

		var params interface{}
		if c.Request.Method == "GET" {
			params = c.Request.URL.Query()
		} else {
			bodyBytes, err := ioutil.ReadAll(c.Request.Body)
			if err == nil {
				params = string(bodyBytes)
				c.Request.Body = &RequestBody{body: bodyBytes}
			}
		}

		req := LogRequest{
			Method: c.Request.Method,
			Params: params,
		}
		reqJSON, _ := json.Marshal(req)

		logger.log.Info("Request",
			zap.String("req_id", logger.reqID),
			zap.String("ip", c.ClientIP()),
			zap.String("path", c.Request.URL.Path),
			zap.String("method", c.Request.Method),
			zap.String("params", string(reqJSON)),
		)

		c.Next()

		endTime := time.Now()
		latency := endTime.Sub(startTime)
		status := c.Writer.Status()

		if loggerValue, exists := c.Get("logger"); exists { // 使用断言语法
			if logger, ok := loggerValue.(*Logger); ok { // 进行类型断言
				logger.log.Info("Response",
					zap.String("req_id", logger.reqID),
					zap.String("ip", c.ClientIP()),
					zap.String("path", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
					zap.String("params", string(reqJSON)),
					zap.Int("status", status),
					zap.Duration("latency", latency),
				)
				return
			}
		}

		// 如果未找到 logger，则创建一个新的 logger
		newLogger := &Logger{
			log: log,
		}

		newLogger.log.Info("Response",
			zap.String("req_id", newLogger.reqID),
			zap.String("ip", c.ClientIP()),
			zap.String("path", c.Request.URL.Path),
			zap.String("method", c.Request.Method),
			zap.String("params", string(reqJSON)),
			zap.Int("status", status),
			zap.Duration("latency", latency),
		)

	}
}
