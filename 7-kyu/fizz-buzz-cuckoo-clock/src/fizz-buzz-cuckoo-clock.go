package src

import (
	"strconv"
	"strings"
)

func FizzBuzzCuckooClock(time string) string {
	hour, _ := strconv.Atoi(time[0:2])
	minute, _ := strconv.Atoi(time[3:5])

	switch {
	case minute == 0:
		// (hour+11) -> shifts the hour by 11, helps handle conversion of midnight
		// %12 -> ensuring the result is in the range of 0-11
		// + "Cuckoo" -> adds an additional Cuckoo to the end of the string, properly account for the full 12 hour cycle
		return strings.Repeat("Cuckoo ", (hour+11)%12) + "Cuckoo"
	case minute == 30:
		return "Cuckoo"
	case minute%15 == 0:
		return "Fizz Buzz"
	case minute%3 == 0:
		return "Fizz"
	case minute%5 == 0:
		return "Buzz"
	default:
		return "tick"
	}
}

// func FizzBuzzCuckooClock(time string) string {
// 	timeParts := strings.Split(time, ":")
// 	hour, minute := timeParts[0], timeParts[1]
// 	hourInt, _ := strconv.Atoi(hour)
// 	minuteInt, _ := strconv.Atoi(minute)

// 	if minuteInt%3 != 0 && minuteInt%5 != 0 {
// 		return "tick"
// 	}

// 	if minuteInt%3 == 0 && minuteInt%5 != 0 {
// 		return "Fizz"
// 	}

// 	if minuteInt%5 == 0 && minuteInt%3 != 0 {
// 		return "Buzz"
// 	}

// 	// On the hour mark
// 	if minuteInt == 0 {
// 		hourMod := hourInt % 12
// 		if hourMod == 0 {
// 			hourMod = 12
// 		}
// 		return strings.TrimRight(strings.Repeat("Cuckoo ", hourMod), " ")
// 	}

// 	// Half hour
// 	if minuteInt%30 == 0 {
// 		return "Cuckoo"
// 	}

// 	return "Fizz Buzz"
// }
