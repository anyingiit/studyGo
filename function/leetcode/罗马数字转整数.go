package leetcode

//func romanToInt(s string) int {
//	//输入确保其在1~3999范围内
//	//正序遍历, 元素根据get的结果进行解析, 如果发现前边的比后边的数字要小,就需要根据后面的数字减去前面的数字,否则直接相加
//	var result int = 0
//	for i,_:=range s{
//		if i != len(s)-1 && romanNumberToInt(s[i]) < romanNumberToInt(s[i+1]) {
//			result -= romanNumberToInt(s[i])
//		}else {
//			result += romanNumberToInt(s[i])
//		}
//	}
//	return result
//}
//func romanNumberToInt(s uint8) int{
//	//写出对应关系
//	romanNumbersMap := map[uint8]int{'I':1,'V':5,'X':10,'L':50,'C':100,'D':500,'M':1000}
//	return romanNumbersMap[s]
//}
//

func romanToInt(s string) int {
	//输入确保其在1~3999范围内
	//正序遍历, 元素根据get的结果进行解析, 如果发现前边的比后边的数字要小,就需要根据后面的数字减去前面的数字,否则直接相加
	romanNumbersMap := map[uint8]int{'I':1,'V':5,'X':10,'L':50,'C':100,'D':500,'M':1000}
	var result int = 0
	for i:=0;i<len(s)-1;i++{
		if romanNumbersMap[s[i]] < romanNumbersMap[s[i+1]]{
			result -= romanNumbersMap[s[i]]
		} else {
			result += romanNumbersMap[s[i]]
		}
	}
	result += romanNumbersMap[s[len(s)-1]]
	return result
}
//func getValue(s uint8) int {
//	switch s {
//	case 'I':
//		return 1
//	case 'V':
//		return 5
//	case 'X':
//		return 10
//	case 'L':
//		return 50
//	case 'C':
//		return 100
//	case 'D':
//		return 500
//	case 'M':
//		return 1000
//	default:
//		return 0
//	}
//}
//func romanToInt(s string) int {
//	//输入确保其在1~3999范围内
//	//正序遍历, 元素根据get的结果进行解析, 如果发现前边的比后边的数字要小,就需要根据后面的数字减去前面的数字,否则直接相加
//	var result int = 0
//	for i:=0;i<len(s)-1;i++{
//		if getValue(s[i]) < getValue(s[i+1]){
//			result -= getValue(s[i])
//		} else {
//			result += getValue(s[i])
//		}
//	}
//	result += getValue(s[len(s)-1])
//	return result
//}
