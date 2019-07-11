package utility

import (
	"errors"
	"io/ioutil"
	"os"
)

func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

func IsDir(path string) bool {
	f, err := os.Stat(path)
	if nil != err {
		return false
	}
	return f.IsDir()
}

// 读取文件的全部内容并返回
func FileGetContents(filename string) (string, error) {
	if !FileExists(filename) {
		return "", errors.New("file not found")
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
