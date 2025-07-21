package server

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samnyan/go-telegram-archive/server/logger"
)

func RequestLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:          true,
		LogStatus:       true,
		LogLatency:      true,
		LogRemoteIP:     true,
		LogHost:         true,
		LogMethod:       true,
		LogError:        true,
		LogRequestID:    true,
		LogResponseSize: true,
		HandleError:     true, // 自动处理错误
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log := logger.Info()

			// 如果有错误，切换到错误级别
			if v.Error != nil {
				log = logger.Error().Err(v.Error)
			}

			// 记录请求信息
			log.Str("request_id", v.RequestID).
				Str("remote_ip", v.RemoteIP).
				Str("host", v.Host).
				Str("method", v.Method).
				Str("uri", v.URI).
				Int("status", v.Status).
				Int64("response_size", v.ResponseSize).
				Dur("latency", v.Latency).
				Time("time", time.Now()).
				Msg("HTTP Request")

			return nil
		},
	})
}
