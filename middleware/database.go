package middleware

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var database = initDB()

func initDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "development.db")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func NewDatabase() *gorm.DB {
	return database
}
