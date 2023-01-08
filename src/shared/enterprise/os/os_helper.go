package os

import (
	"os"
)

type OsHelper interface {
	FileExist(path string) bool
	DirExist(path string) bool
	CreateDir(path string) error
	GetFileAsByteArray(path string) ([]byte, error)
	GetFileAsString(path string) (string, error)
	CreateTemporaryTextFile(content string) (string, error)
	GetEnvironmentVariable(key string) string
}

type osHelper struct {
}

func NewOsHelper() (OsHelper, error) {
	return &osHelper{},
		nil
}

func (oh *osHelper) FileExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	if fileInfo.IsDir() {
		return false
	}

	return true
}

func (oh *osHelper) DirExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	if !fileInfo.IsDir() {
		return false
	}

	return true
}

func (oh *osHelper) CreateDir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

func (oh *osHelper) GetFileAsByteArray(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (oh *osHelper) GetFileAsString(path string) (string, error) {
	buf, err := oh.GetFileAsByteArray(path)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (oh *osHelper) CreateTemporaryTextFile(content string) (string, error) {
	file, err := os.CreateTemp(os.TempDir(), "")
	if err != nil {
		return "", err
	}

	defer func() {
		_ = file.Close()
	}()

	_, err = file.WriteString(content)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func (oh *osHelper) GetEnvironmentVariable(key string) string {
	return os.Getenv(key)
}
