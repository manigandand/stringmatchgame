package engine

func checkIndex(lastCheckedChar []string, byterandomStr byte) bool {
	for _, element := range lastCheckedChar {
		if element == string(byterandomStr) {
			return false
		}
	}
	return true
}
