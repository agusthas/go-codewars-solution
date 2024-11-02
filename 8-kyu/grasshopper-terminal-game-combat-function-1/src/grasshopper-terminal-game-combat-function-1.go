package src

func combat(health, damage float64) float64 {
	// resulting health cannot be less than 0
	result := health - damage
	if result < 0 {
		result = 0
	}
	return result
}
