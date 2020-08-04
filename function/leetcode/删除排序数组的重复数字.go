package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var ralLen = 1 //既然长度不为零就至少为1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			ralLen++
			nums[ralLen-1] = nums[i+1]
		}
	}
	nums = nums[:ralLen]
	return len(nums)
}
