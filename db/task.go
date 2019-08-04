package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	gorm.Model
	Title    string    `gorm:"type:varchar(255);not null"`
	Label    string    `gorm:"type:varchar(64)"`
	Status   string    `gorm:"type:varchar(32)"`
	DueDate  time.Time `gorm:""`
	Priority int       `gorm:""`
	HideFlg  bool      `gorm:"type:boolean;not null;default false"`
}

func (t *Task) GetTitle() string {
	displayTitle := t.Title
	if t.HideFlg {
		displayTitle = fmt.Sprintf("(close) %s\n", t.Title)
	}
	return displayTitle
}

type TaskRepository interface {
	Insert(task *Task) error
	Update(task *Task) error
	GetList(allFlag bool) *[]Task
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
	if db.First(&task, id).RecordNotFound() {
		return nil
	}
	return &task
}

func (t *TaskRepositoryImpl) GetList(allFlag bool) *[]Task {
	db := getDbConnection()
	defer db.Close()

	tasks := []Task{}
	if !allFlag {
		db = db.Where("hide_flg = false")
	}
	db = db.Order("id desc")
	db.Find(&tasks)

	return &tasks
}

func (p *TaskRepositoryImpl) DeleteTask(id int) {
	db := getDbConnection()
	defer db.Close()

	db.Delete(Task{}, "id = ?", id)
}
