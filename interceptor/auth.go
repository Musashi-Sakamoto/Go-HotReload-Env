package interceptor

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
)

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
		if username == "musashi" && password == "password" {
			return true, nil
		}
		return false, nil
	})
}