package src

import "fmt"

func Digits(n uint64) int {
	return len(fmt.Sprintf("%d", n))
}

// import "strconv"

// func Digits(n uint64) int {
// 	// turn n into a string
// 	s := strconv.FormatUint(n, 10)

// 	// count the number of digits in the string
// 	return len(s)
// }
