package utility

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

//Exec 执行一个外部程序
//参考 https://stackoverflow.com/questions/10781516/how-to-pipe-several-commands-in-go
// 例：`ps -eaf|grep "nginx: master"|grep -v "grep"|awk '{print $2}'`
func Exec(command string, bashEnv ...string) ([]string, error) {
	bash := `/bin/sh`
	if 0 < len(bashEnv) {
		bash = bashEnv[0]
	}
	out, err := exec.Command(bash, `-c`, command).Output()
	if err != nil {
		return []string{}, err
	}
	return strings.Split(strings.Trim(string(out), "\n"), "\n"), nil
}

func Getenv(varname string) string {
	return os.Getenv(varname)
}

func Getcwd() (dir string, err error) {
	return os.Getwd()
}

//返回当前的函数名,如果对次函数封装，则适当设置skip的值为2默认是1
func FuncName(skip ...int) string {
	sk := 1
	if 0 < len(skip) && skip[0] > 0 {
		sk = skip[0]
	}
	pc, _, _, b := runtime.Caller(sk)
	if !b {
		return ""
	}
	return runtime.FuncForPC(pc).Name()
}
