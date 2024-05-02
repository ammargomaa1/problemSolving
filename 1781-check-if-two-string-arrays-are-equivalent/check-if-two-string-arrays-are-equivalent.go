func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    firstCompleteWord := ""
    secondCompleteWord := ""

    for _, word := range word1 {
        firstCompleteWord = firstCompleteWord + word
    } 

    for _, word := range word2 {
        secondCompleteWord = secondCompleteWord + word
    }    

    return firstCompleteWord == secondCompleteWord 
}