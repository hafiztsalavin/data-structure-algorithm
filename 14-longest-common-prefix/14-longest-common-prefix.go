func longestCommonPrefix(strs []string) string {
    size := len(strs)
    
    if len(strs) == 0 {
        return ""
    } else if  len(strs) == 1{
        return strs[0]
    }
    
    compare := strs[0]
    for i:= 1; i < size; i++{
        lenCompare := len(compare)
        lenTemp := len(strs[i])
        
        length := minLength(lenCompare, lenTemp)
        
        j := 0
        for j < length{
            if compare[j] != strs[i][j]{
                break
            }
            j++
        }
        compare = compare[:j]
        
        if len(compare) == 0{
            return ""
        }
        
    }
    return compare
}

func minLength(a, b int) int{
    if a < b{
        return a
    } else {
        return b
    }
}