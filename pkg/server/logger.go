package server

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func getLoggerFunc() echo.MiddlewareFunc {
	env := os.Getenv("ENV")
	if env == "PRODUCTION" {
		return middleware.Logger()
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:        true,
		LogLatency:    true,
		LogMethod:     true,
		LogRemoteIP:   true,
		LogStatus:     true,
		LogUserAgent:  true,
		LogValuesFunc: makeLoggerFunc(log.Logger),
	})
}

func makeLoggerFunc(logger zerolog.Logger) func(echo.Context, middleware.RequestLoggerValues) error {
	return func(_ echo.Context, v middleware.RequestLoggerValues) error {
		logger.Info().
			Str("URI", v.URI).
			Str("latency", v.Latency.String()).
			Str("method", v.Method).
			Str("remote", v.RemoteIP).
			Int("status", v.Status).
			Str("user-agent", v.UserAgent).
			Msg("request")

		return nil
	}
}
