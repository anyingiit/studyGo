package leetcode

import "sort"

func 合并区间() {

}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i,j int) bool{
		return intervals[i][0] < intervals[j][0]
	})
	var result [][]int
	for i := 0; i < len(intervals);i++ {
		if len(result) == 0 || intervals[i][0] > result[len(result)-1][1] {
			result = append(result,intervals[i])
		} else if intervals[i][1] > result[len(result)-1][1]{
			result[len(result)-1][1] = intervals[i][1]
		}
	}
	//var tmp = intervals[0]
	//var result [][]int
	//for i:=0;i < len(intervals); i++ {
	//	if(tmp[1] >= intervals[i][0]){ //执行合并
	//		tmp[1] = intervals[i][1]
	//	} else { //执行提交
	//		result = append(result,tmp)
	//		tmp = intervals[i]
	//		if(i == len(intervals)-1){
	//			result = append(result,tmp)
	//		}
	//	}
	//}
	return result
}