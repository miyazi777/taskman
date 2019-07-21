package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func InitDb() {
	db, err := gorm.Open("sqlite3", "/tmp/test.db") // TODO dbの場所をどうにかすること
	if err != nil {
		panic("データベースへの接続に失敗しました") // TODO エラーハンドリングをどうにかすること
	}
	defer db.Close()

	db.AutoMigrate(&Project{})
	db.AutoMigrate(&Task{})
}
