package leetcode

func rotate(matrix [][]int)  {
	//N:= len(matrix)
	//for i := 0; i < N/2; i++ {
	//	matrix[i], matrix[N-1-i] = matrix[N-1-i] , matrix[i]
	//}
	//for i := 0; i < N; i++ {
	//	for j := 0; j < i; j++ {
	//		matrix[i][j],matrix[j][i] = matrix[j][i],matrix[i][j]
	//	}
	//}
	var result [][]int
	for i := 0; i < len(matrix); i++ {
		var tmp []int
		for j := len(matrix[i])-1; j >= 0; j-- {
			tmp = append(tmp,matrix[j][i])
		}
		result = append(result,tmp)
	}
	for i := 0; i < len(matrix) ; i++ {
		matrix[i] = result[i]
	}
}