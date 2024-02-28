func isPrefixOfWord(sentence string, searchWord string) int {
    lenSearchWord := len(searchWord)
    
    words := strings.Fields(sentence)
    for idx, word := range words{
        if len(word) >= lenSearchWord{
            if word[:lenSearchWord] == searchWord{
                return idx+1
            }

        } 
    }
    return -1
}