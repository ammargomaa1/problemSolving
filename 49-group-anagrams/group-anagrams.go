func groupAnagrams(strs []string) [][]string {
    outPut := [][]string{}

    wordsMapped := map[string][]string{}
    for _, word := range strs {
        sortedWord := []rune(word)

        sort.Slice(sortedWord,func(i,j int)bool{return sortedWord[i] < sortedWord[j]})
        
        sortedStr := string(sortedWord)

        wordsMapped[sortedStr] = append(wordsMapped[sortedStr], word)

    }

    for _, collectedWords := range wordsMapped {
        outPut = append(outPut, collectedWords)
    }

    return outPut
}