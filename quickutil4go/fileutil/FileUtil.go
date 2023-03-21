package fileutil

import (
	"bufio"
	"io/ioutil"
	"os"
)

func GetCurrentPath() string {
	path, _ := os.Getwd()
	return path
}

func File2Byte(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func File2String(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return true, err
}

func File2Lines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	resultList := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		resultList = append(resultList, scanner.Text())
	}
	err2 := scanner.Err()
	return resultList, err2
}

func Byte2File(path string, bytes []byte) error {
	return ioutil.WriteFile(path, bytes, 0664)
}

func String2File(path string, content string) error {
	return ioutil.WriteFile(path, []byte(content), 0664)
}
