func numberOfLines(widths []int, s string) []int {
    tempTotal := 0
    // totalEndLines := 0
    endOfCounter := 1
    
    size := len(s)
    for i := 0; i < size; i++{
        tempTotal = tempTotal + widths[s[i]-97]
        if tempTotal == 100 && i == size-1{
            tempTotal = tempTotal
        } else if tempTotal == 100{
            tempTotal = 0
            endOfCounter++
        }else if tempTotal > 100{
            tempTotal = widths[s[i]-97]
            endOfCounter++
            
        }
    }
        
    return []int{endOfCounter, tempTotal}
}