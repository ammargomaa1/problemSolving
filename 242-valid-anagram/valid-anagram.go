func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mappedLetterCounter := make([]int, 26)

	for _, letterS := range s {
		mappedLetterCounter[letterS - 'a']++
	}

	for _, letterT := range t {
		mappedLetterCounter[letterT - 'a']--

		if mappedLetterCounter[letterT - 'a'] < 0 {
			return false
		}
	}

	return true
}
