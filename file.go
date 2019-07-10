package utility

import "os"

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
