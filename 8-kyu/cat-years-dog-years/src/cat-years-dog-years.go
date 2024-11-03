package src

func CalculateYears(years int) (result [3]int) {
	switch years {
	case 1:
		result = [3]int{1, 15, 15}
	case 2:
		result = [3]int{2, 24, 24}
	default:
		result = [3]int{years, 24 + 4*(years-2), 24 + 5*(years-2)}
	}
	return
}

// func CalculateYears(years int) (result [3]int) {
// 	result[0] = 0
// 	result[1] = 0
// 	result[2] = 0
// 	for i := 0; i < years; i++ {
// 		result[0]++
// 		if i == 0 {
// 			result[1] += 15
// 			result[2] += 15
// 		} else if i == 1 {
// 			result[1] += 9
// 			result[2] += 9
// 		} else {
// 			result[1] += 4
// 			result[2] += 5
// 		}
// 	}
// 	return result
// }
