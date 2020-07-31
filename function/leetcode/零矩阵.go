package leetcode

func setZeroes(matrix [][]int) {
	var target [][]int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				target = append(target, []int{i, j})
			}
		}
	}
	for i := 0; i < len(target); i++ {
		for k := 0; k < len(matrix[target[i][0]]); k++ {
			matrix[target[i][0]][k] = 0
		}
		for l := 0; l < len(matrix); l++ {
			matrix[l][target[i][1]] = 0
		}
	}

}
