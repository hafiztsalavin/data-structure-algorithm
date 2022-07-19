func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    wordFirst := ""
    wordSecond := ""
    
    size1 := len(word1)
    size2 := len(word2)
    
    i := 0
    j := 0
    for i < size1 || j < size2{
        if i < size1{
            wordFirst = wordFirst + word1[i]
            i++
        }
        
        if j < size2{
            wordSecond = wordSecond + word2[j]
            j++
        }
    }
    fmt.Println(wordFirst, wordSecond)
    
    if wordFirst != wordSecond{
        return false
    }
    
    return true
}