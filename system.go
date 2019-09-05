package utility

import (
	"os"
	"os/exec"
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
	return strings.Split(strings.Trim(string(out),"\n"), "\n"), nil
}

func Getenv(varname string) string {
	return os.Getenv(varname)
}

func Getcwd() (dir string, err error) {
	return os.Getwd();
}