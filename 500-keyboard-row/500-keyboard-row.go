
func findWords(words []string) []string {
    wordResult := []string{}
        
    for _, word := range words {
        result := checkWord(word)
        if result == true{
            wordResult = append(wordResult, word)
        }
    } 
    
    return wordResult
}

func checkWord(word string) bool {
    word = strings.ToLower(word)
    mapChar := []int{1, 2, 2, 1, 0, 1, 1, 1, 0, 1, 1, 1, 2, 2, 0, 0, 0, 0, 1, 0, 0, 2, 0, 2, 0, 2}
    
    firstChar := mapChar[word[0]-'a']
    for i := 1; i < len(word); i++{
        if mapChar[word[i]-'a'] != firstChar{
            return false
        }
    }
    
    return true
}