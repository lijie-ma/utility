package utility

import "math"

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

func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}