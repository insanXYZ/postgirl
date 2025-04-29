package util

import (
	"os"
	"postgirl/model"
)

var (
	rootPath      = os.Getenv("HOME") + "/postgirl/"
	cacheFileName = ".cache"
)

func WriteFile(filename string, data []byte) error {
	CreateDir(rootPath)

	return os.WriteFile(rootPath+filename, data, os.ModePerm)
}

func WriteCache(data []byte) error {
	return WriteFile(cacheFileName, data)
}

func CreateDir(dirName string) error {
	return os.Mkdir(dirName, os.ModePerm)
}

func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(rootPath + filename)
}

func ReadCache() (map[string]model.Request, error) {
	b, err := ReadFile(cacheFileName)
	if err != nil {
		return nil, err
	}

	b, err = Decode(string(b))
	if err != nil {
		return nil, err
	}

	m := make(map[string]model.Request)

	err = JsonUnmarshal(b, &m)
	return m, err

}

func ReadDir(dirName string) ([]os.DirEntry, error) {
	return os.ReadDir(dirName)
}
