package handler

import (
	"net/http"
	"github.com/labstack/echo"
	"app/model"
	"app/database"
	"fmt"
)

func ListUser(c echo.Context) error {
	db := database.GormConnect()
	defer db.Close()
	users := []model.User{}
	db.Find(&users)
	fmt.Println(users)
	return c.JSON(http.StatusOK, users)
}