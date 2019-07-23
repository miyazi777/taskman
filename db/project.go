package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Project struct {
	gorm.Model
	Title string `gorm:"type:varchar(255);not null"`
}

type Projectrepository interface {
	Insert(project *Project) error
}

type ProjectRepositoryImpl struct{}

func (p *ProjectRepositoryImpl) Insert(project *Project) error {
	db := GetDbConnection()
	defer db.Close()

	db.Create(project)
	return nil
}
