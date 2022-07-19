func restoreString(s string, indices []int) string {
    size := len(indices)
    resultArr := make([]string, size)
    result := ""
    
    for i:=0; i < size; i++{
        resultArr[indices[i]] = string(s[i])
    }
    
    for _, value := range resultArr{
        result = result + value
    }
    
    return result
    
}