package leetcode

import "sort"

//存在自行优化可能
func longestCommonPrefix(strs []string) string {
	// 将第一个字符串记录, 用剩下的字符串进行比较, 不断更新公共头的范围坐标
	// 将字符串长度进行排序, 使用最短的进行比较
	// 直接遍历整个数组, 发现一致就+1, 某个数组小就停
	if len(strs) == 0 {
		return ""
	}
	var resultIndex = 0
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs)-1; j++ {
			if strs[j][i] != strs[j+1][i] {
				return strs[0][:resultIndex]
			}
		}
		resultIndex++
	}
	return strs[0][:resultIndex]
	//for i := 0; i < len(strs)-1; i++ {
	//	for j := 0; j < len(strs[0]); j++ {
	//		if strs[i][j] != strs[i+1][j] {
	//			return strs[0][:resultIndex]
	//		}
	//	}
	//	resultIndex++
	//}
	//return strs[0][:resultIndex]
}
