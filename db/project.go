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
	Update(project *Project) error
	GetList() *[]Project
	CheckExistsCurrentProject() bool
	GetCurrentProject() *Project
}

type ProjectRepositoryImpl struct{}

func (t *ProjectRepositoryImpl) Insert(project *Project) error {
	db := getDbConnection()
	defer db.Close()

	db.Create(project)
	return nil
}

func (t *ProjectRepositoryImpl) Update(project *Project) error {
	db := getDbConnection()
	defer db.Close()

	db.Save(project)
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

func (t *ProjectRepositoryImpl) FindById(id int) *Project {
	db := getDbConnection()
	defer db.Close()

	project := Project{}
	if db.First(&project, id).RecordNotFound() {
		return nil
	}
	return &project
}

func (t *ProjectRepositoryImpl) CheckExistsCurrentProject() bool {
	db := getDbConnection()
	defer db.Close()

	var count int
	projects := []Project{}
	db.Where("current_flg = 1").Find(&projects).Count(&count)

	var result bool = false
	if count > 0 {
		result = true
	}
	return result
}

func (t *ProjectRepositoryImpl) GetCurrentProject() *Project {
	db := getDbConnection()
	defer db.Close()

	var project = Project{}
	if db.Where("current_flg = 1").First(&project).RecordNotFound() {
		return nil
	}
	return &project
}
