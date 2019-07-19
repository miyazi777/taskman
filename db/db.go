package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func InitDb() {
	var dbPath = "/tmp/test.db" // TODO 名前はあとで考える
	os.Remove(dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	var query string
	query = `
		CREATE TABLE projects (
		  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(255) NOT NULL,
			add_date DATETIME NOT NULL,
			update_date DATETIME NOT NULL
		)
	`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
		return
	}

	query = `
		CREATE TABLE tasks (
		  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			project_id INTEGER NOT NULL,
			title VARCHAR(255) NOT NULL,
			status VARCHAR(32),
			due_date DATE,
			priority INTEGER,
			add_date DATETIME NOT NULL,
			update_date DATETIME NOT NULL,
      foreign key (project_id) references projects(id)
		)
	`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
		return
	}
}
