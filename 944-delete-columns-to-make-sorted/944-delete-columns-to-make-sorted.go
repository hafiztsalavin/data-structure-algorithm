func minDeletionSize(strs []string) int {
    lengthArr := len(strs)
    lengthWord := len(strs[0])
    
    counter := 0
    
    for i:=0; i < lengthWord; i++{
        for j := 0; j < lengthArr-1; j++{
            if strs[j][i] > strs[j+1][i]{
                counter++
                break
            }
        }
    }
    
    fmt.Println(counter)
    return counter
}