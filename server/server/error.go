package server

import (
	"errors"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/samnyan/go-telegram-archive/server/api"
	"github.com/samnyan/go-telegram-archive/server/logger"
)

func ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	var result api.ApiResult
	code := http.StatusInternalServerError

	if ae, ok := err.(*api.AppError); ok {
		result.SetAppError(ae)
		logger.Error().Stack().Err(ae).Str("message", ae.Message).Int("code", ae.Code).Msg("应用错误")
	} else if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if he.Internal != nil {
			var tokenExtractErr *echojwt.TokenExtractionError
			if errors.As(he.Internal, &tokenExtractErr) {
				code = api.ErrTokenInvalid.Code
				result.SetAppError(api.ErrTokenInvalid)
			}
		} else {
			result.Code = he.Code
			if str, ok := he.Message.(string); ok {
				result.Message = str
			} else {
				result.Message = "Error"
			}
		}
	} else {
		code = api.ErrInternalServer.Code
		result.SetAppError(api.ErrInternalServer)
		logger.Error().Err(err).Msg("Internal Error")
	}

	if c.Request().Method == echo.HEAD {
		err := c.NoContent(code)
		if err != nil {
			logger.Error().Err(err).Msg("Response Write Error")
		}
	} else {
		err := c.JSON(code, result)
		if err != nil {
			c.Logger().Error(err)
		}
	}
}
