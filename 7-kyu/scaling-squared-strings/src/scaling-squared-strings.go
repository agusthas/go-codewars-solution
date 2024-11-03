package src

import "strings"

func Scale(s string, k, n int) (result string) {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		var hscale string
		for _, char := range line {
			hscale += strings.Repeat(string(char), k)
		}
		result += strings.Repeat(hscale+"\n", n)
	}

	return strings.TrimRight(result, "\n")
}

// func Scale(s string, k, n int) (result string) {
// 	if len(s) == 0 {
// 		return s
// 	}

// 	lines := strings.Split(s, "\n")
// 	for _, line := range lines {
// 		// horizontal scaling
// 		chars := strings.Split(line, "")
// 		for i, char := range chars {
// 			chars[i] = strings.Repeat(char, k)
// 		}
// 		horizontal := strings.Join(chars, "")
// 		// vertical scaling
// 		for i := 0; i < n; i++ {
// 			result += horizontal + "\n"
// 		}
// 	}
// 	// your code
// 	return strings.TrimRight(result, "\n")
// }
