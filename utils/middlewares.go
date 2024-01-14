package utils

import (
	"github.com/google/uuid"
	"github.com/jspmarc/mockha/entities"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func SetupEchoMiddlewares(e *echo.Echo) {
	e.Use(middlewareCustomContext())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middlewareRequestId())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		HandleError:  true,
		LogURI:       true,
		LogStatus:    true,
		LogMethod:    true,
		LogRequestID: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("requestId", v.RequestID).
				Str("Uri", v.URI).
				Str("method", v.Method).
				Int("status", v.Status).
				Msg("Finished request")

			return nil
		},
	}))
}

func middlewareCustomContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := entities.CustomContext{Context: c, Headers: &entities.Headers{}}
			return next(cc)
		}
	}
}

func middlewareRequestId() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			headerName := echo.HeaderXRequestID

			requestId := req.Header.Get(headerName)
			if requestId == "" {
				requestId = uuid.NewString()
			}

			req.Header.Set(headerName, requestId)
			res.Header().Set(headerName, requestId)
			c.(entities.CustomContext).Headers.RequestId = requestId

			return next(c)
		}
	}
}
