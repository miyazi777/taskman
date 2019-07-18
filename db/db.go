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
	}

	_, err = db.Exec(`CREATE TABLE "world" ("id" INTEGER PRIMARY KEY AUTOINCREMENT, "country" VARCHAR(255), "capital" VARCHAR(255))`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`INSERT INTO "world" ("country", "capital") VALUES (?, ?) `,
		"日本",
		"test",
	)
	if err != nil {
		panic(err)
	}

	db.Close()
}
