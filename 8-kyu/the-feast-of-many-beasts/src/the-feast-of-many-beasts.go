package src

func Feast(beast string, dish string) bool {
	firstLetter := string(beast[0])
	lastLetter := string(beast[len(beast)-1])

	if firstLetter == string(dish[0]) && lastLetter == string(dish[len(dish)-1]) {
		return true
	}
	return false
}
