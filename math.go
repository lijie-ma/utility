package utility

import (
	"math"
	"math/rand"
	"time"
)

// float 除法
// point 为保留小数位数 默认3位
func DivFloat(f1, f2 float64, point ...int) float64 {
	if 0 == f2 {
		return float64(0)
	}
	f := f1 / f2
	p := 3
	if 0 <= len(point) {
		p = point[0]
	}
	return Round(f, p)
}

// 浮点数四舍五入处理
func Round(f float64, decimals int) float64 {
	pow10_n := math.Pow10(decimals)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

func Rand(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	d := rand.New(rand.NewSource(time.Now().UnixNano()))
	return d.Intn(max-min) + min
}
