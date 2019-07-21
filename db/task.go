package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Task struct {
	gorm.Model
	Title     string     `gorm:"type:varchar(255);not null"`
	Status    string     `gorm:"type:varchar(32);not null"`
	Due_date  *time.Time `gorm:"not null"`
	Priority  int        `gorm:"not null"`
	Project   Project    `gorm:"foreignkey:ProjectId"`
	ProjectId uint
}
