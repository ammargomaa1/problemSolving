func groupAnagrams(strs []string) [][]string {
    outPut := [][]string{}

    chosenWords := map[string] bool{}

    for index,word := range strs {
        angrams := []string{}
        if _,found := chosenWords[word]; !found && len(angrams) == 0 {
            angrams = append(angrams, word)
        }else{
            continue
        }

        for i:= index+1; i<len(strs); i++ {
            if isAngram(word, strs[i]){
                 angrams = append(angrams, strs[i])
                 chosenWords[strs[i]] = true
            }
        }

        outPut = append(outPut, angrams)
    }

    return outPut
}

func isAngram(firstWord string, secondWord string) bool {
    if len(firstWord) != len(secondWord) {
        return false
    }

    letterCounter := map[rune]int{}

    for _,letter := range firstWord {
        letterCounter[letter]++
    } 

    for _,letter2 := range secondWord {
        letterCounter[letter2]--

        if letterCounter[letter2] < 0 {
            return false
        }
    } 

    return true
}