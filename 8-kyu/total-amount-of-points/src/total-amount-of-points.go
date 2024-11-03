package src

import (
	"strings"
)

func Points(games []string) (result int) {
	result = 0
	for _, game := range games {
		parts := strings.Split(game, ":")
		x, y := parts[0], parts[1]
		switch {
		case x > y:
			result += 3
		case x == y:
			result += 1
		}
	}
	return result
}

// func Points(games []string) int {
// 	var result int
// 	for _, game := range games {
// 		// split game string into two parts by ":"
// 		parts := strings.Split(game, ":")
// 		if len(parts) != 2 {
// 			panic("Invalid game string")
// 		}

// 		x, err := strconv.Atoi(parts[0])
// 		if err != nil {
// 			panic(err)
// 		}

// 		y, err := strconv.Atoi(parts[1])
// 		if err != nil {
// 			panic(err)
// 		}

// 		if x > y {
// 			result += 3
// 		} else if x == y {
// 			result += 1
// 		} else {
// 			result += 0
// 		}
// 	}
// 	return result
// }
