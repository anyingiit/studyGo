package leetcode

func findMagicIndex(nums []int) int {
	//for i := 0; i < len(nums); i++ {
	//	if i == nums[i] {
	//		return i
	//	}
	//}
	//return -1
	for i := 0; i < len(nums); i++ {
		if i == nums[i] {
			return i
		}
		if nums[i] > i+1 {
			i = nums[i]
		}
	}
	return -1

}
