package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Task struct {
	gorm.Model
	Title    string     `gorm:"type:varchar(255);not null"`
	Status   string     `gorm:"type:varchar(32)"`
	Due_date *time.Time `gorm:""`
	Priority int        `gorm:""`
}

type TaskRepository interface {
	Insert(task *Task) error
	GetList() *[]Task
}

type TaskRepositoryImpl struct{}

func (t *TaskRepositoryImpl) Insert(task *Task) error {
	db := getDbConnection()
	defer db.Close()

	db.Create(task)
	return nil
}

func (t *TaskRepositoryImpl) GetList() *[]Task {
	db := getDbConnection()
	defer db.Close()

	tasks := []Task{}
	db.Order("id desc").Find(&tasks)

	return &tasks
}
