package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "app/handler"
  "app/interceptor"
)

func main() {
  e := echo.New()

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
  }))

  g := e.Group("/admin")
  g.Use(interceptor.BasicAuth())
  g.GET("/hello/:username", handler.MainPage())

  e.Start(":1323")
}
