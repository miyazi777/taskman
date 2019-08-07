package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	gorm.Model
	Title     string    `gorm:"type:varchar(255);not null"`
	Label     string    `gorm:"type:varchar(64)"`
	Status    string    `gorm:"type:varchar(32)"`
	DueDate   time.Time `gorm:""`
	Priority  int       `gorm:""`
	HideFlg   bool      `gorm:"type:boolean;not null;default false"`
	ProjectID uint
}

func (t *Task) GetTitle() string {
	displayTitle := t.Title
	if t.HideFlg {
		displayTitle = fmt.Sprintf("(close) %s\n", t.Title)
	}
	return displayTitle
}

func (t *Task) GetDueDate() string {
	return t.DueDate.Format("2006/01/02")
}

type TaskRepository interface {
	Insert(task *Task) error
	Update(task *Task) error
	GetList(allFlag bool, label string, sort string) *[]Task
	FindById(id int) *Task
	DeleteTask(id int)
}

type TaskRepositoryImpl struct{}

func (t *TaskRepositoryImpl) Insert(task *Task) error {
	db := getDbConnection()
	defer db.Close()

	db.Create(task)
	return nil
}

func (t *TaskRepositoryImpl) Update(task *Task) error {
	db := getDbConnection()
	defer db.Close()

	db.Save(&task)
	return nil
}

func (p *TaskRepositoryImpl) FindById(id int) *Task {
	db := getDbConnection()
	defer db.Close()

	task := Task{}
	if db.Scopes(joinCurrentProject).First(&task, id).RecordNotFound() {
		return nil
	}
	return &task
}

func (t *TaskRepositoryImpl) GetList(allFlag bool, label string, sort string) *[]Task {
	db := getDbConnection()
	defer db.Close()

	tasks := []Task{}
	if !allFlag {
		db = db.Where("hide_flg = false")
	}
	if label != "" {
		db = db.Where("label = ?", label)
	}
	if sort != "" {
		sortCond := fmt.Sprintf("%s asc", sort)
		db = db.Order(sortCond)
	}
	db = db.Order("id desc").Scopes(joinCurrentProject)
	db.Find(&tasks)

	return &tasks
}

func (p *TaskRepositoryImpl) DeleteTask(id int) {
	db := getDbConnection()
	defer db.Close()

	db.Delete(Task{}, "id = ?", id)
}

func joinCurrentProject(db *gorm.DB) *gorm.DB {
	return db.Joins("join projects on projects.id = tasks.project_id and projects.current_flg = 1")
}
