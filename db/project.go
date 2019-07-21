package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Project struct {
	gorm.Model
	Title string `gorm:"type:varchar(255);not null"`
}
