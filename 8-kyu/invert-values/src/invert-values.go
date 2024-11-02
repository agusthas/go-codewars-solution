package src

func Invert(arr []int) []int {
	newArr := make([]int, len(arr))
	for i, v := range arr {
		newArr[i] = v * -1
	}
	return newArr
}
