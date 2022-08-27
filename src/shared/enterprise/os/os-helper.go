package os

import "os"

type OsHelper interface {
	FileExist(path string) bool
	DirExist(path string) bool
	GetFileAsByteArray(path string) ([]byte, error)
	GetFileAsString(path string) (string, error)
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
