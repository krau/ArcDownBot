package dao

import (
	"arcdownbot/model"

	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/ncruces/go-sqlite3/gormlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	var err error
	db, err = gorm.Open(gormlite.Open("arcdownbot.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.AutoMigrate(&model.Version{})
}
