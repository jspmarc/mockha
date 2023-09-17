package utils

import (
	"github.com/jspmarc/mockha/dto"
	"github.com/labstack/echo/v4"
	"time"
)

func ResponseHttp(ctx echo.Context, httpCode int, errorMsg *string, data interface{}) error {
	response := dto.Response{
		Data:         data,
		ErrorMessage: errorMsg,
		ServerTime:   time.Now().UnixMilli(),
	}

	return ctx.JSON(httpCode, response)
}
