package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	homedir "github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

func InitDb() {
	db := getDbConnection()
	defer db.Close()

	db.AutoMigrate(&Task{})
}

func getDbConnection() *gorm.DB {
	err := makeWorkDir()
	if err != nil {
		panic("ワークディレクトリの作成に失敗しました") // TODO エラーハンドリングをどうにかする
	}

	db, err := gorm.Open("sqlite3", getDbPath())
	if err != nil {
		panic("データベースへの接続に失敗しました") // TODO エラーハンドリングをどうにかすること
	}
	return db
}

func makeWorkDir() error {
	workDirPath := getWorkDirPath()
	_, err := os.Stat(workDirPath)
	if err != nil {
		os.Mkdir(workDirPath, 0755)
	}

	return nil
}

func getDbPath() string {
	return filepath.Join(getWorkDirPath(), "tasks.db")
}

// TODO workDir構造体を作成する
func getWorkDirPath() string {
	home, err := homedir.Dir()
	if err != nil {
		panic("ホームディレクトリの取得に失敗しました") // TODO エラーハンドリングをどうにかすること
	}
	return filepath.Join(home, ".taskman")
}
