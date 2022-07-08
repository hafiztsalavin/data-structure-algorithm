func lengthOfLastWord(s string) int {
    count := 0
    
    for word := range s {
        if s[word] != ' '{
            if word >0 && s[word - 1] == ' '{
                count = 0
            }
            
            count++
        }
    }
    return count
}