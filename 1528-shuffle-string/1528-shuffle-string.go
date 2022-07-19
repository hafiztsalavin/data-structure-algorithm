func restoreString(s string, indices []int) string {
    size := len(indices)
    resultArr := make([]byte, size)

    
    for i:=0; i < size; i++{
        resultArr[indices[i]] = s[i]
    }
    
    
    return string(resultArr)    
}