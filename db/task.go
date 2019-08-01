package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Task struct {
	gorm.Model
	Title    string    `gorm:"type:varchar(255);not null"`
	Status   string    `gorm:"type:varchar(32)"`
	DueDate  time.Time `gorm:""`
	Priority int       `gorm:""`
	HideFlg  bool      `gorm:"type:boolean;not null;default false"`
}

type TaskRepository interface {
	Insert(task *Task) error
	Update(task *Task) error
	GetList() *[]Task
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

func (t *TaskRepositoryImpl) GetList() *[]Task {
	db := getDbConnection()
	defer db.Close()

	tasks := []Task{}
	db.Where("hide_flg = false").Order("id desc").Find(&tasks)

	return &tasks
}

func (p *TaskRepositoryImpl) DeleteTask(id int) {
	db := getDbConnection()
	defer db.Close()

	db.Delete(Task{}, "id = ?", id)
}
