package handler

import (
	"net/http"
	"github.com/labstack/echo"
)

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		jsonMap := map[string]string{
			"username": username,
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}