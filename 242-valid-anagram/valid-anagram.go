func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for _, sLetter := range s {
		tHasLetter := false
		for tIndex, tLetter := range t {
			if sLetter == tLetter {
				tHasLetter = true
				t = removeLetterFromString(tIndex, t)
				break
			}
		}

		if !tHasLetter {
			return tHasLetter
		}

		s = s[1:]

	}

	return true
}

func removeLetterFromString(index int, s string) string {
	if index < 0 || index >= len(s) {
		return s // Index out of bounds, return original string
	}
	return s[:index] + s[index+1:]
}
