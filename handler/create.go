package handler

import (
	"net/http"
	"github.com/labstack/echo"
	"app/model"
	"app/database"
)

func CreateUser(c echo.Context) error {
	db := database.GormConnect()
	defer db.Close()

	u := &model.User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	u.Name = "test"
	db.Table("User").Create(&u)
	return c.JSON(http.StatusCreated, u)
}