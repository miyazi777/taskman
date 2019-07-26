package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Project struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);not null"`
	Current bool   `gorm:"type:integer(1);not null;default:0"`
}

type Projectrepository interface {
	Insert(project *Project) error
	FindById(id int) *Project
	FindByCurrent() (*Project, error)
	CountByCurrent() int
	GetList() *[]Project
	DeleteProject(id int)
}

type ProjectRepositoryImpl struct{}

func (p *ProjectRepositoryImpl) Insert(project *Project) error {
	db := getDbConnection()
	defer db.Close()

	db.Create(project)
	return nil
}

func (p *ProjectRepositoryImpl) FindById(id int) *Project {
	db := getDbConnection()
	defer db.Close()

	project := Project{}
	if db.First(&project, id).RecordNotFound() {
		return nil
	}
	return &project
}

func (p *ProjectRepositoryImpl) FindByCurrent() (*Project, error) {
	db := getDbConnection()
	defer db.Close()

	count := p.CountByCurrent()
	if count == 0 {
		return nil, errors.New("Nothing current project.")
	}

	project := Project{}
	db.First(&project, "current = ?", true)
	return &project, nil
}

// TODO クエリの条件をscope化すること
func (p *ProjectRepositoryImpl) CountByCurrent() int {
	db := getDbConnection()
	defer db.Close()

	project := Project{}
	var count int = 0
	db.Where("current = ?", true).Find(&project).Count(&count)
	return count
}

func (p *ProjectRepositoryImpl) GetList() *[]Project {
	db := getDbConnection()
	defer db.Close()

	projects := []Project{}
	db.Order("id desc").Find(&projects)

	return &projects
}

func (p *ProjectRepositoryImpl) DeleteProject(id int) {
	db := getDbConnection()
	defer db.Close()

	db.Delete(Project{}, "id = ?", id)
}
