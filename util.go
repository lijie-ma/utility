package utility

import (
	"fmt"
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

func Atoi(arg interface{}) int {
	num, err := strconv.Atoi(arg.(string))
	if nil != err {
		return 0
	}
	return num
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

func FloattoStr(f float64, precision int) string {
	return strconv.FormatFloat(f, 'g', precision, 64)
}