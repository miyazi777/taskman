package repository

import (
	homedir "github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

type MemoRepository interface {
	AddMemo(title string) error
}

type MemoRepositoryImpl struct{}

func (m *MemoRepositoryImpl) AddMemo(title string) error {
	path := GetMemoPath(title)
	_, err := os.Create(path)
	if err != nil {
		return err
	}
	return nil
}

func GetMemoPath(title string) string {
	fileName := title + ".md"
	return filepath.Join(getWorkDirPath(), fileName)
}

// TODO workDir構造体を作成する
func getWorkDirPath() string {
	home, err := homedir.Dir()
	if err != nil {
		panic("ホームディレクトリの取得に失敗しました") // TODO エラーハンドリングをどうにかすること
	}
	return filepath.Join(home, ".taskman")
}
