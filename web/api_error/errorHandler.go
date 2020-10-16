package api_error

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func InvalidRequestError(err error) *echo.HTTPError {
	return &echo.HTTPError{Code: 400, Message: "invalid request error", Internal: err}
}

type ApiError struct {
	StatusCode       int      `json:"status_code"`
	Message          string   `json:"message"`
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := ""
	// https://godoc.org/github.com/labstack/echo#HTTPError
	if ee, ok := err.(*echo.HTTPError); ok {
		code = ee.Code
		message = ee.Message.(string)
	}
	body := ApiError{
		StatusCode:       code,
		Message:          message,
	}
	c.JSON(code, body)
}