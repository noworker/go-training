package api_error

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func InvalidRequestError(err error) *echo.HTTPError {
	return &echo.HTTPError{Code: 400, Message: err.Error()}
}

func NotFoundError(err error) *echo.HTTPError {
	return &echo.HTTPError{Code: 404, Message: err.Error()}
}

func InternalError(err error) *echo.HTTPError {
	return &echo.HTTPError{Code: 500, Message: err.Error()}
}

type ApiError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := ""
	if ee, ok := err.(*echo.HTTPError); ok {
		code = ee.Code
		message = ee.Message.(string)
	}
	body := ApiError{
		StatusCode: code,
		Message:    message,
	}
	c.JSON(code, body)
}
