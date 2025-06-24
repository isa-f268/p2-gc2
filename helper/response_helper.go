package helper

import (
	"main/model"

	"github.com/labstack/echo/v4"
)

func RespHandler(message string, data any) model.Response[any] {
	resp := model.Response[any]{
		Message: message,
		Data:    data,
	}

	return resp
}

func SuccessHandler(c echo.Context, responCode int, message string, data any) error {
	resp := RespHandler(message, data)
	return c.JSON(responCode, resp)
}

func ErrorHandler(c echo.Context, responCode int, message string) error {
	resp := RespHandler(message, nil)

	return echo.NewHTTPError(responCode, resp)

}

func IdError(c echo.Context) error {
	resp := RespHandler("invalid request parameters", nil)
	return echo.NewHTTPError(400, resp)
}

func JsonErr(c echo.Context) error {
	resp := RespHandler("invalid request parameters", nil)
	return echo.NewHTTPError(400, resp)
}

// error new
func InternalErr(c echo.Context) error {
	resp := RespHandler("internal server error", nil)

	return echo.NewHTTPError(500, resp)

}
func InvalidReq(c echo.Context) error {
	resp := RespHandler("invalid request parameters", nil)
	return echo.NewHTTPError(400, resp)
}
