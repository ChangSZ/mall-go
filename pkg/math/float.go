package math

import "math"

// RoundHalfEven 实现四舍五入到指定的小数位数
func RoundHalfEven(value float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	fv := value * shift
	fv = math.Round(fv)
	return fv / shift
}
