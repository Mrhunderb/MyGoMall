package db

import (
	"fmt"
	"mygomall/service/user/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(user, pass, host, port, name string) (*gorm.DB, error) {
	// connect to database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	db.AutoMigrate(&model.User{})
	return db, nil
}
