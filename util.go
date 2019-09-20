package utility

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func Ip2Long(ip string) int64 {
	bits := strings.Split(ip, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

func Long2Ip(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func Md5(encode string) string {
	return fmt.Sprintf(`%x`, md5.Sum([]byte(encode)))
}

//Base64Encode 摘自go 源码
func Base64Encode(encode string) string {
	return base64.StdEncoding.EncodeToString([]byte(encode))
}

//Base64Decode
//解码失败返回 空字符串
func Base64Decode(encode string) string {
	b, e := base64.StdEncoding.DecodeString(encode)
	if nil != e {
		return ``
	}
	return string(b)
}

func IsNumeric(i interface{}) bool {
	t := reflect.TypeOf(i)
	if int(t.Kind()) > int(reflect.Bool) && int(t.Kind()) < int(reflect.Float64) {
		return true
	}
	if t.Kind() != reflect.String {
		return false
	}
	b, err := regexp.Match(`^-?\d+\.?\d*$`, []byte(i.(string)))
	if nil != err {
		return false
	}
	return b
}
