package leetcode

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xStr := strconv.Itoa(x)
	for i, _ := range xStr {
		if !(xStr[len(xStr)-1-i] == xStr[i]) {
			return false
		}
	}
	return true
}

//func isPalindrome(x int) bool {
//	if x < 0 {
//		return false
//	}
//	tmp := x
//	target := 0
//	for x != 0 {
//
//		target = target*10 + x%10
//		x /= 10
//	}
//	if tmp == target {
//		return true
//	}
//	return false
//}
