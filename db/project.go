package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Project struct {
	gorm.Model
	Title      string `gorm:"type:varchar(255);not null"`
	CurrentFlg bool   `gorm:"type:boolean;not null;default false"`
}

type ProjectRepository interface {
	Insert(project *Project) error
	GetList() *[]Project
	CheckExistsCurrentProject() bool
}

type ProjectRepositoryImpl struct{}

func (t *ProjectRepositoryImpl) Insert(project *Project) error {
	db := getDbConnection()
	defer db.Close()

	db.Create(project)
	return nil
}

func (t *ProjectRepositoryImpl) GetList() *[]Project {
	db := getDbConnection()
	defer db.Close()

	projects := []Project{}
	db = db.Order("id desc")
	db.Find(&projects)

	return &projects
}

func (t *ProjectRepositoryImpl) CheckExistsCurrentProject() bool {
	db := getDbConnection()
	defer db.Close()

	var count int
	projects := []Project{}
	db.Where("").Find(&projects).Count(&count)

	var result bool = false
	if count > 0 {
		result = true
	}
	return result
}
