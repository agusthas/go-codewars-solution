package src

func FakeBin(x string) (result string) {
	for _, char := range x {
		if char < '5' {
			result += "0"
		} else {
			result += "1"
		}
	}
	return
}

// import "strconv"

// func FakeBin(x string) string {
// 	var result string
// 	// loop over characters in x
// 	for i := 0; i < len(x); i++ {
// 		// convert character to int
// 		char := x[i]
// 		num, err := strconv.Atoi(string(char))
// 		if err != nil {
// 			panic(err)
// 		}

// 		if num >= 5 {
// 			result += "1"
// 		} else {
// 			result += "0"
// 		}
// 	}
// 	return result
// }
