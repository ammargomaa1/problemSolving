func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sortedS := []rune(s)
	sort.Slice(sortedS, func(i, j int) bool {return sortedS[i] < sortedS[j]})

	sortedT := []rune(t)
	sort.Slice(sortedT, func(i, j int) bool {return sortedT[i] < sortedT[j]})
	
	return string(sortedT) == string(sortedS)
}
