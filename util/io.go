package util

import "os"

func WriteFile(filename string, data []byte) error {
	rootPath := os.Getenv("HOME") + "/postgirl"
	CreateDir(rootPath)

	return os.WriteFile(rootPath+"/"+filename, data, os.ModePerm)
}

func CreateDir(dirName string) error {
	return os.Mkdir(dirName, os.ModePerm)
}
