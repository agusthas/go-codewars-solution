package src

func Incrementer(n []int) []int {
	for i := range n {
		n[i] = (i + 1 + n[i]) % 10
	}
	return n
}

// func Incrementer(n []int) []int {
// 	result := make([]int, len(n))
// 	for i := 0; i < len(n); i++ {
// 		result[i] = (n[i] + (i + 1)) % 10
// 	}
// 	// your code here
// 	return result
// }
