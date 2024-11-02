package src

import (
	"fmt"
	"strings"
)

func MultiTable(number int) (result string) {
	for _, i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		result += fmt.Sprintf("%d * %d = %d\n", i, number, i*number)
	}
	result = strings.TrimRight(result, "\n")
	return
}
