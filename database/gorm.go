package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"app/model"
)

func GormConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "docker_user:docker_user_pwd@tcp(db)/docker_db?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&model.User{})
	return db
}