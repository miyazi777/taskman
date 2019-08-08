package repository

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

type MemoRepository interface {
	AddProjectSpace(projectId uint) error
	AddMemo(projectId uint, title string) error
	ChangeMemo(projectId uint, oldTitle string, newTitle string) error
	DeleteAllMemo() error
}

type MemoRepositoryImpl struct{}

func (m *MemoRepositoryImpl) AddProjectSpace(projectId uint) error {
	workDirPath := getProjectDirPath(projectId)
	return os.Mkdir(workDirPath, 0755)
}

func (m *MemoRepositoryImpl) AddMemo(projectId uint, title string) error {
	path := GetMemoPath(projectId, title)
	_, err := os.Create(path)
	if err != nil {
		return err
	}
	return nil
}

func (m *MemoRepositoryImpl) ChangeMemo(projectId uint, oldTitle string, newTitle string) error {
	oldTitlePath := GetMemoPath(projectId, oldTitle)
	newTitlePath := GetMemoPath(projectId, newTitle)
	err := os.Rename(oldTitlePath, newTitlePath)
	if err != nil {
		return err
	}
	return nil
}

func (m *MemoRepositoryImpl) DeleteAllMemo() error {
	if err := os.RemoveAll(getWorkDirPath()); err != nil {
		return err
	}
	return nil
}

func GetMemoPath(projectId uint, title string) string {
	fileName := title + ".md"
	return filepath.Join(getProjectDirPath(projectId), fileName)
}

func getProjectDirPath(projectId uint) string {
	workPath := getWorkDirPath()
	return filepath.Join(workPath, fmt.Sprint(projectId))
}

// TODO workDir構造体を作成する
func getWorkDirPath() string {
	home, err := homedir.Dir()
	if err != nil {
		panic("ホームディレクトリの取得に失敗しました") // TODO エラーハンドリングをどうにかすること
	}
	return filepath.Join(home, ".taskman")
}
