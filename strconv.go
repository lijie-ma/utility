package utility

import (
	"fmt"
	"reflect"
	"strconv"
)

func Atoi(arg interface{}) int {
	num, err := strconv.Atoi(arg.(string))
	if nil != err {
		return 0
	}
	return num
}

// IntVal alias of AtoInt64
func IntVal(arg interface{}) int64 {
	return AtoInt64(arg)
}

// 接收基本类型，并转变为字符串
func StrVal(arg interface{}) string {
	refType := reflect.TypeOf(arg)
	typeValue := refType.Kind()
	switch typeValue {
	case reflect.String:
		return arg.(string)
	case reflect.Float64, reflect.Float32:
		return fmt.Sprintf("%f", arg)
	}
	if typeValue >= reflect.Int && typeValue < reflect.Uint64 {
		return fmt.Sprintf("%d", arg)
	}
	return ""
}

func AtoFloat64(arg interface{}) float64 {
	num, err := strconv.ParseFloat(arg.(string), 32)
	if nil != err {
		return 0.0
	}
	return num
}

func AtoInt64(arg interface{}) int64 {
	num, err := strconv.ParseInt(arg.(string), 10, 64)
	if nil != err {
		return int64(0)
	}
	return num
}

// alias of AtoFloat64
func FloatVal(arg interface{}) float64 {
	return AtoFloat64(arg)
}

func FloattoStr(f float64, precision int) string {
	return NumberFormat(f, precision)
}

// 十进制转换为二进制
func DecBin(i int64) string {
	return fmt.Sprintf(`%b`, i)
}

// 二进制转换为十进制
func BinDec(s string) int64 {
	i := int64(0)
	length := len(s)
	for j := 0; j < length; j++ {
		if `1` == s[j:j+1] {
			i += 1 << uint(length-j-1)
		}
	}
	return i
}

func DecHex(i int64) string {
	return fmt.Sprintf(`%x`, i)
}

func HexDec(s string) int64 {
	i := int64(0)
	length := len(s)
	if 0 == length {
		return i
	}
	mm := map[string]int64{
		`0`: 0,
		`1`: 1,
		`2`: 2,
		`3`: 3,
		`4`: 4,
		`5`: 5,
		`6`: 6,
		`7`: 7,
		`8`: 8,
		`9`: 9,
		`a`: 10,
		`b`: 11,
		`c`: 12,
		`d`: 13,
		`e`: 14,
		`f`: 15,
		`A`: 10,
		`B`: 11,
		`C`: 12,
		`D`: 13,
		`E`: 14,
		`F`: 15,
	}
	filter := ``
	for j := 0; j < length; j++ {
		if _, ok := mm[s[j:j+1]]; ok {
			filter += s[j : j+1]
		}
	}
	length = len(filter)
	for j := 0; j < length; j++ {
		if `0` != filter[j:j+1] {
			i += mm[filter[j:j+1]] * 1 << uint((length-j-1)*4)
		}
	}
	return i
}
