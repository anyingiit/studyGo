package leetcode

import (
	"math"
)

func reverse(x int) int {
	var result int
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}
	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}
	return result
	//var str string
	//var result []uint8
	//if x < 0 {
	//	str = strconv.Itoa(x + (-x)*2)
	//	result = append(result, uint8('-'))
	//} else {
	//	str = strconv.Itoa(x)
	//}
	//for i, _ := range str {
	//	result = append(result, str[len(str)-1-i])
	//}
	//resultInt, _ := strconv.Atoi(string(result))
	//if resultInt <= math.MaxInt32 && resultInt >= math.MinInt32 {
	//	return resultInt
	//}
	//return 0
}
