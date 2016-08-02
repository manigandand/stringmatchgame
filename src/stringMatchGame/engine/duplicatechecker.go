package engine

func checkIndex(lastCorrectChar []string, byterandomStr byte) bool {
	for _, element := range lastCorrectChar {
		if element == string(byterandomStr) {
			return false
		}
	}
	return true
}
