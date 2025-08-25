package middlewares

import (
	"wallet-api-go-bc/logging"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func RequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			logging.Logger.Info("Incoming request",
				zap.String("method", req.Method),
				zap.String("path", req.URL.Path),
				zap.String("remote_ip", c.RealIP()),
			)
			return next(c)
		}
	}
}