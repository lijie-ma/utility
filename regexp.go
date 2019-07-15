package utility

import "regexp"

const RegChinese = `[\x{4e00}-\x{9fa5}]`

//执行一个全局正则表达式匹配
func RegMatchAll(hack, pattern string) [][]string {
	reg := regexp.MustCompile(pattern)
	return reg.FindAllStringSubmatch(hack, -1)
}

//RegReplace 执行一个正则表达式的搜索和替换
func RegReplace(src, repl, pattern string) string {
	s := regexp.MustCompile(pattern)
	return s.ReplaceAllString(src, repl)
}

//RegReplace 执行一个正则表达式的搜索和替换
func RegReplaceCallback(src string, repl func(v string) string, pattern string) string {
	s := regexp.MustCompile(pattern)
	return s.ReplaceAllStringFunc(src, repl)
}
