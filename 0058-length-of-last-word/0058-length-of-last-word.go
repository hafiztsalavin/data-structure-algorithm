func lengthOfLastWord(s string) int {
    total := 0
    for i := len(s)-1; i >= 0; i--{
        if string(s[i]) != " "{
            total+=1
        } else {
            if total > 0 {
                return total
            }
        }
    }
    return total
    
}